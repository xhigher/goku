package apiclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	globalClients = map[string]*ApiClient{}
	mutext        sync.RWMutex
)

func NewApiClient(serviceName string) *ApiClient {
	mutext.RLock()
	if client, ok := globalClients[serviceName]; ok {
		mutext.RUnlock()
		return client
	}
	mutext.RUnlock()

	mutext.Lock()
	defer mutext.Unlock()

	if client, ok := globalClients[serviceName]; ok {
		return client
	}

	client := &ApiClient{
		serviceName: serviceName,
		endpoints:   make(map[string]*Endpoint),
		sdInstance: SdInstance{
			ServiceName: serviceName,
		},
		lb: RandLoadBalancer{},
		m:  sync.Mutex{},
	}
	client.sdInstance.ServiceUpdate(client.updateEndpoints)
	globalClients[serviceName] = client
	return client
}

type ApiClient struct {
	serviceName string
	endpoints   map[string]*Endpoint
	sdInstance  SdInstance
	lb          LoadBalancer
	m           sync.Mutex
}

func (client *ApiClient) Do(uri string, method string, req, resp interface{}) (err error) {
	key := fmt.Sprintf("%s-%s", uri, method)
	ep, err := client.SelectEndpoint(key, req)
	if err != nil {
		return err
	}
	return ep.Do(uri, method, req, resp)
}

func (client *ApiClient) SelectEndpoint(key string, req interface{}) (*Endpoint, error) {
	client.m.Lock()
	defer client.m.Unlock()
	epMap := make(map[string]*Endpoint)
	for addr, endpoint := range client.endpoints {
		cbkey := fmt.Sprintf("%s:%s", addr, key)
		cb, _, err := hystrix.GetCircuit(cbkey)
		log.Println(addr, err, " db.IsOpen ", cb.IsOpen(), " cb.AllowRequest ", cb.AllowRequest())
		if cb != nil && err == nil {
			if !cb.IsOpen() {
				epMap[addr] = endpoint
			} else {
				if cb.AllowRequest() {
					epMap[addr] = endpoint
				} else {
					log.Println(addr, " open and not allow request")
				}
			}
		} else {
			hystrix.ConfigureCommand(cbkey, hystrixConfig)
			epMap[addr] = endpoint
		}
	}

	return client.lb.Select(req, key, epMap), nil
}

func (client *ApiClient) updateEndpoints(e ServiceEvent) error {
	client.m.Lock()
	newHostMap := make(map[string]bool, len(e.hosts))
	for _, host := range e.hosts {
		newHostMap[host] = true
	}
	newEndpointMap := make(map[string]*Endpoint)
	for host, e := range client.endpoints {
		if newHostMap[host] {
			newEndpointMap[host] = e
			delete(newHostMap, host)
			delete(client.endpoints, host)
		}
	}
	oldEndpoints := client.endpoints
	client.endpoints = newEndpointMap
	client.m.Unlock()
	for addr := range newHostMap {
		client.m.Lock()
		client.endpoints[addr] = newEndpoint(addr)
		client.m.Unlock()
	}
	for _, endpoint := range oldEndpoints {
		endpoint.Close()
	}
	return nil
}

func newEndpoint(addr string) *Endpoint {
	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	timeout := time.Second * 5
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   timeout / 2,
			KeepAlive: 3600 * time.Second,
		}).Dial,
		MaxIdleConnsPerHost: 10 * 10000,
		MaxIdleConns:        100,
		IdleConnTimeout:     100,
		TLSHandshakeTimeout: timeout,
		TLSClientConfig:     tlsCfg,
	}

	c := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	return &Endpoint{
		Addr:       addr,
		httpClient: c,
		Proportion: 100,
		Tags:       nil,
		m:          sync.RWMutex{},
	}
}

type Endpoint struct {
	Addr       string
	httpClient *http.Client
	Proportion int
	Tags       map[string]bool
	m          sync.RWMutex
}

func (e *Endpoint) Close() {
	e.m.Lock()
	defer e.m.Unlock()
	e.httpClient.CloseIdleConnections()
}

func (e *Endpoint) Do(uri string, method string, req, resp interface{}) (err error) {
	key := fmt.Sprintf("%s-%s", uri, method)
	cbkey := fmt.Sprintf("%s:%s", e.Addr, key)
	err = hystrix.Do(cbkey, func() error {
		return e.DoRaw(uri, method, req, resp)
	}, nil)
	if err != nil {
		log.Println(e.Addr, err)
	}
	return err
}

func (e *Endpoint) DoRaw(uri, method string, req, resp interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://%s%s", e.Addr, uri)
	log.Println(url, string(body))
	httpReq, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	httpResp, err := e.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %d", httpResp.StatusCode)
	}
	respData, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respData, resp)
	if err != nil {
		return err
	}
	return nil
}

type LoadBalancer interface {
	Select(req interface{}, uri string, ep map[string]*Endpoint) *Endpoint
}

type RandLoadBalancer struct {
}

func (RandLoadBalancer) Select(req interface{}, uri string, ep map[string]*Endpoint) *Endpoint {
	eps := make([]*Endpoint, 0, len(ep))
	steps := make([]int, 0, len(ep))
	totalWeight := 0
	for _, endpoint := range ep {
		if endpoint.Proportion == 0 {
			totalWeight += 100
		} else {
			totalWeight += endpoint.Proportion
		}
		steps = append(steps, totalWeight)
		eps = append(eps, endpoint)
	}
	n := rand.Intn(totalWeight)
	idx := 0
	for i, step := range steps {
		if n < step {
			idx = i
			break
		}
	}
	return eps[idx]
}

/**

1. 服务发现 resolver
2. 负载均衡 balancer
3. 断路器 circuit breaker
4. 重试策略 retry strategy
*/
