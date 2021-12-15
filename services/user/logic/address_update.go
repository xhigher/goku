package logic

import (
	"goku.net/framework/network/http"
)

type UserAddressUpdateParam struct {
	*http.BaseBodyData
	Mid int64 `json:"mid"`
}

type UserAddressUpdate struct {
	*http.BaseExecutor
}

func (executor *UserAddressUpdate) SupportMethods() []string {
	return []string{"POST"}
}

func (executor *UserAddressUpdate) Execute() http.ResponseData {
	switch executor.Version {
	case 1:
		return executor.executeV1()
	}
	return executor.ResultErrorRequest()
}

func (executor *UserAddressUpdate) executeV1() http.ResponseData {
	param := executor.BodyData.(*UserDetailParam)
	param.Mid = 123456
	return executor.ResultOKData(param)
}
