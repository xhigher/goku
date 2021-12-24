package apiclient

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"log"
)

type SdConsulConfig struct {
	Address    string
	Datacenter string
}

var (
	sdMgr *SdMgr
)

func InitSdConsul(sdConfig SdConsulConfig) {
	config := &consul.Config{
		Address:    "",
		Datacenter: "",
	}
	consulClient, err := consul.NewClient(config)
	if err != nil {
		return
	}
	sdMgr = &SdMgr{
		Type:   ConsulSd,
		client: consulClient,
	}
}

type SdType string

const (
	ConsulSd SdType = "consul"
)

type SdMgr struct {
	Type   SdType
	client *consul.Client
}

func (sd SdMgr) services(serviceName string, index uint64) (lastIndex uint64, hosts []string, err error) {
	services, meta, err := sd.client.Health().Service(serviceName, "", true, &consul.QueryOptions{
		WaitIndex: index,
	})
	if err != nil {
		return
	}
	for _, service := range services {
		addr := service.Node.Address
		if service.Service.Address != "" {
			addr = service.Service.Address
		}
		hosts = append(hosts, fmt.Sprintf("%s:%d", addr, service.Service.Port))
	}
	lastIndex = meta.LastIndex
	return
}

type ServiceEvent struct {
	serviceName string
	hosts       []string
}

//func (sd SdMgr) notify(serviceName string, index uint64) (<-chan ServiceEvent, error) {
//	var (
//		ch        = make(chan ServiceEvent)
//		lastIndex = index
//	)
//	go func() {
//		for {
//			services, meta, err := sd.client.Health().Service(serviceName, "", true, &consul.QueryOptions{
//				WaitIndex: lastIndex,
//			})
//			if err != nil {
//				continue
//			}
//			e := ServiceEvent{
//				serviceName: serviceName,
//				hosts:       nil,
//			}
//			for _, service := range services {
//				addr := service.Node.Address
//				if service.Service.Address != "" {
//					addr = service.Service.Address
//				}
//				e.hosts = append(e.hosts, fmt.Sprintf("%s:%d", addr, service.Service.Port))
//			}
//			select {
//			case ch <- e:
//				lastIndex = meta.LastIndex
//			}
//		}
//	}()
//	return ch, nil
//}

type SdInstance struct {
	ServiceName string
}

func (s SdInstance) ServiceUpdate(f func(e ServiceEvent) error) {
	index, hosts, err := sdMgr.services(s.ServiceName, 0)
	log.Println(hosts)
	if err != nil {
		// TODO log
	} else {
		f(ServiceEvent{
			serviceName: s.ServiceName,
			hosts:       hosts,
		})
	}
	go func() {
		for {
			lastIndex, hosts, err := sdMgr.services(s.ServiceName, index)
			if err != nil {
				continue
			}
			err = f(ServiceEvent{
				serviceName: s.ServiceName,
				hosts:       hosts,
			})
			if err != nil {
				continue
			}
			index = lastIndex
		}
	}()
}
