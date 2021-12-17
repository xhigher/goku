package http

type HttpInterceptor interface {
	Intercept(executor Executor) bool
}
