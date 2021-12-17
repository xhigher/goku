package logic

import (
	"goku.net/framework/network/http"
	"goku.net/services/user/cache"
	"goku.net/services/user/model"
)

type UserDetailParam struct {
	*http.BaseBodyData
	Mid int64 `json:"mid"`
}

type UserDetail struct {
	*http.BaseExecutor
}

func (executor *UserDetail) SupportMethods() []string {
	return []string{"POST"}
}

func (executor *UserDetail) RequireSession() bool {
	return true
}

func (executor *UserDetail) Execute() http.ResponseData {
	switch executor.Version {
	case 1:
		return executor.executeV1()
	case 2:
		return executor.executeV2()
	}
	return executor.ResultErrorRequest()
}

func (executor *UserDetail) executeV1() http.ResponseData {
	param := executor.BodyData.(*UserDetailParam)

	data, err := cache.GetUserDetail(param.Mid)
	if err != nil {
		return executor.ResultErrorInternal()
	}

	data = model.GetUserDetail(param.Mid)

	return executor.ResultOKData(data)
}

func (executor *UserDetail) executeV2() http.ResponseData {
	param := executor.BodyData.(*UserDetailParam)
	param.Mid = 654321
	return executor.ResultOKData(param)
}
