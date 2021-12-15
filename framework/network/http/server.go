package http

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type HttpServer struct {
	port              int
	executorFactories map[string]ModuleExecutorFactory
}

var (
	pathRegexp = regexp.MustCompile(`^/v([\d]+)/([\w]+)/([\w]+)$`)
)

func NewServer(port int) HttpServer {
	return HttpServer{
		port:              port,
		executorFactories: make(map[string]ModuleExecutorFactory),
	}
}

func (server HttpServer) AddModule(executorFactory ModuleExecutorFactory) {
	server.executorFactories[executorFactory.ModuleName()] = executorFactory
}

func (server HttpServer) Start() {
	addr := fmt.Sprintf("0.0.0.0:%d", server.port)
	http.HandleFunc("/", server.handler)
	http.ListenAndServe(addr, nil)

}

// handler
func (server HttpServer) handler(write http.ResponseWriter, request *http.Request) {

	url := request.RequestURI
	var lastIndex = strings.LastIndex(url, "?")
	if lastIndex > -1 {
		url = url[:lastIndex]
	}

	params := pathRegexp.FindStringSubmatch(url)
	if len(params) != 4 {
		write.WriteHeader(404)
		return
	}
	//path := params[0]
	version, _ := strconv.Atoi(params[1])
	module := params[2]
	action := params[3]

	executorFactory, ok := server.executorFactories[module]
	if !ok {
		write.WriteHeader(404)
		return
	}

	executor := executorFactory.Create(version, action)
	if executor == nil {
		write.WriteHeader(404)
		return
	}

	logicExecutor := &LogicExecutor{
		executor: executor,
	}
	if !logicExecutor.CheckMethod(request.Method) {
		write.WriteHeader(405)
		return
	}
	if !logicExecutor.CheckParams(request) {
		return
	}

	logicExecutor.Execute(write)
}
