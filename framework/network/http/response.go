package http

type ErrorCode int32

const (
	OK             = 0
	NOK            = 1
	ERR_INTERNAL   = 500
	ERR_PARAMETER  = 501
	ERR_SESSION    = 502
	ERR_TOKEN      = 503
	ERR_PERMISSION = 504
	ERR_CAPTCHA    = 505
	ERR_FREQUENTLY = 506
)

type ResponseData struct {
	Code ErrorCode
	Msg  string
	Data interface{}
}
