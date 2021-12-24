package apiclient

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	consul "github.com/hashicorp/consul/api"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	log.SetFlags(log.Flags() | log.LstdFlags | log.Lshortfile)
	var err error
	InitSdConsul(SdConsulConfig{
		Address:    "localhost:8500",
		Datacenter: "",
	})
	client := NewApiClient("service")

	for {

		req := Msg{Msg: "asdasdas"}
		resp := Msg{}
		err = client.Do("/hello", http.MethodPost, req, &resp)
		if err != nil {
			log.Println(err)
		}
	}

}

func StartServer() {
	log.SetFlags(log.Flags() | log.LstdFlags | log.Lshortfile)
	g := gin.New()
	g.POST("/hello", func(context *gin.Context) {
		resp := json.RawMessage([]byte(`{
	"msg": "hello"
}`))
		context.JSON(http.StatusOK, resp)
	})
	g.Any("/healthcheck", helathCheck)
	go g.Run(":5050")
	go g.Run(":5051")

	consulClient, err := consul.NewClient(&consul.Config{
		Address: "localhost:8500",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Println("del")
		DelService(consulClient.Agent(), 5050)
		DelService(consulClient.Agent(), 5051)
		DelService(consulClient.Agent(), 5052)
	}()
	AgentRegister(consulClient.Agent(), 5050)
	AgentRegister(consulClient.Agent(), 5051)
	AgentRegister(consulClient.Agent(), 5052)
	go g.Run(":5052")

	time.Sleep(time.Second * 10)
}

type Msg struct {
	Msg string `json:"msg"`
}

func TestRun(t *testing.T) {
	StartServer()
}

func TestDel(t *testing.T) {
	client, err := consul.NewClient(&consul.Config{
		Address: "localhost:8500",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		DelService(client.Agent(), 5050)
		DelService(client.Agent(), 5051)
		DelService(client.Agent(), 5052)
	}()

}

func DelService(agent *consul.Agent, port int) {
	portStr := fmt.Sprintf("%d", port)
	err := agent.ServiceDeregister(portStr)
	if err != nil {
		log.Fatal(err)
	}
}

func AgentRegister(agent *consul.Agent, port int) {
	portStr := fmt.Sprintf("%d", port)
	check := fmt.Sprintf("http://192.168.58.40:%d/healthcheck", port)
	log.Println(check)
	//check := fmt.Sprintf("/healthcheck")
	err := agent.ServiceRegister(&consul.AgentServiceRegistration{
		Kind:              "",
		ID:                portStr,
		Name:              "service",
		Tags:              nil,
		Port:              port,
		Address:           "192.168.58.40",
		SocketPath:        "",
		TaggedAddresses:   nil,
		EnableTagOverride: false,
		Meta:              nil,
		Weights:           nil,
		Checks: []*consul.AgentServiceCheck{
			{
				CheckID:                        portStr + ":check",
				Name:                           "",
				Args:                           nil,
				DockerContainerID:              "",
				Shell:                          "",
				Interval:                       "5s",
				Timeout:                        "5s",
				TTL:                            "",
				HTTP:                           check,
				Header:                         nil,
				Method:                         "GET",
				Body:                           "",
				TCP:                            "",
				Status:                         "",
				Notes:                          "",
				TLSServerName:                  "",
				TLSSkipVerify:                  false,
				GRPC:                           "",
				GRPCUseTLS:                     false,
				H2PING:                         "",
				H2PingUseTLS:                   false,
				AliasNode:                      "",
				AliasService:                   "",
				SuccessBeforePassing:           0,
				FailuresBeforeWarning:          0,
				FailuresBeforeCritical:         0,
				DeregisterCriticalServiceAfter: "",
			},
		},
		Proxy:     nil,
		Connect:   nil,
		Namespace: "",
		Partition: "",
	})
	if err != nil {
		log.Println(err)
	}

}

func helathCheck(c *gin.Context) {
	log.Println("health check")
	c.JSON(http.StatusOK, "ok")
}
