package http

type HttpHandler struct {
	module string
}

type HttpRouter struct {
	version   int
	module    string
	action    string
	handler   HttpHandler
	paramType interface{}
}
