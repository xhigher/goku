package http

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"goku.net/framework/commons"
)

type ServerInfo struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	BuildTime int64  `json:"build_time"`
	StartTime int64  `json:"start_time"`
}

type HttpServer struct {
	port              int
	info              ServerInfo
	executorFactories map[string]ModuleExecutorFactory
	interceptors      []HttpInterceptor
}

var (
	basePathRegexp    = regexp.MustCompile(`^/(healthcheck|info)$`)
	servicePathRegexp = regexp.MustCompile(`^/v([\d]+)/([\w]+)/([\w]+)$`)
)

func NewServer(port int, info ServerInfo) HttpServer {
	return HttpServer{
		port:              port,
		info:              info,
		executorFactories: make(map[string]ModuleExecutorFactory),
	}
}

func (server HttpServer) AddInterceptor(interceptor HttpInterceptor) {
	server.interceptors = append(server.interceptors, interceptor)
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
func (server HttpServer) handler(writer http.ResponseWriter, request *http.Request) {

	url := request.RequestURI
	var lastIndex = strings.LastIndex(url, "?")
	if lastIndex > -1 {
		url = url[:lastIndex]
	}

	params := basePathRegexp.FindStringSubmatch(url)
	if len(params) == 2 {
		server.handleDefaultRoute(params[1], writer)
		return
	}

	params = servicePathRegexp.FindStringSubmatch(url)
	if len(params) != 4 {
		writer.WriteHeader(404)
		return
	}
	//path := params[0]
	version, _ := strconv.Atoi(params[1])
	module := params[2]
	action := params[3]

	executorFactory, ok := server.executorFactories[module]
	if !ok {
		writer.WriteHeader(404)
		return
	}

	executor := executorFactory.Create(version, action)
	if executor == nil {
		writer.WriteHeader(404)
		return
	}

	logicExecutor := &LogicExecutor{
		executor: executor,
	}
	if !logicExecutor.CheckMethod(request.Method) {
		writer.WriteHeader(405)
		return
	}
	if !logicExecutor.CheckParams(request) {
		return
	}

	for _, interceptor := range server.interceptors {
		commons.Logger().Info("interceptor", zap.Any("interceptor", interceptor))
		if interceptor.Intercept(executor) {
			writer.WriteHeader(405)
			return
		}
	}

	logicExecutor.Execute(writer)
}

func (server HttpServer) handleDefaultRoute(path string, writer http.ResponseWriter) {
	switch path {
	case "healthcheck":
		data := make(map[string]interface{})
		data["time"] = time.Now().Unix()
		writeResponseData(writer, ResponseData{Code: OK, Msg: "OK", Data: data})
	case "info":
		writeResponseData(writer, ResponseData{Code: OK, Msg: "OK", Data: server.info})
	default:
		writer.WriteHeader(404)
		return
	}
}
