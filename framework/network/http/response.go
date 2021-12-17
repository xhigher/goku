package http

type ErrorCode int32

const (
	OK             = 0
	NOK            = 1
	ERR_REQUEST    = 500
	ERR_PARAMETER  = 501
	ERR_INTERNAL   = 502
	ERR_SESSION    = 503
	ERR_TOKEN      = 504
	ERR_PERMISSION = 505
	ERR_CAPTCHA    = 506
	ERR_FREQUENTLY = 507
)

type ResponseData struct {
	Code ErrorCode   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
