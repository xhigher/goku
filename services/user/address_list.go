package user

import (
	"goku.net/framework/network/http"
)

type UserAddressListParam struct {
	*http.BaseBodyData
	Mid int64 `json:"mid"`
}

type UserAddressList struct {
	*http.BaseExecutor
}

func (executor *UserAddressList) SupportMethods() []string {
	return []string{"POST"}
}

func (executor *UserAddressList) Execute() http.ResponseData {
	switch executor.Version {
	case 1:
		return executor.executeV1()
	}
	return executor.ResultErrorRequest()
}

func (executor *UserAddressList) executeV1() http.ResponseData {
	param := executor.BodyData.(*UserAddressListParam)
	param.Mid = 123456
	return executor.ResultOKData(param)
}
