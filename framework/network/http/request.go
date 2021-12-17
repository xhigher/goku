package http

import "strings"

const (
	UrlEncode    = "application/x-www-form-urlencoded"
	FormData     = "multipart/form-data"
	JSON         = "application/json"
	ContentType  = "Content-Type"
	ContentType2 = "content-type"
)

type HttpMethod int32

const (
	HttpMethod_GET  HttpMethod = 0
	HttpMethod_POST HttpMethod = 1
)

func IsJSONData(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if JSON == contentType || strings.HasPrefix(contentType, JSON) {
		return true
	}
	return false
}

func IsFormData(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if FormData == contentType || strings.HasPrefix(contentType, FormData) {
		return true
	}
	return false
}

func IsUrlEncode(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if UrlEncode == contentType || strings.HasPrefix(contentType, UrlEncode) {
		return true
	}
	return false
}

type BaseBodyData struct {
	App     string `json:"gk_app"`
	Ver     string `json:"gk_ver"`
	Token   string `json:"gk_token"`
	OS      string `json:"gk_os"`
	Model   string `json:"gk_model"`
	Channel string `json:"gk_chan"`
	DevId   string `json:"gk_devid"`
}
