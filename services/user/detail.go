package user

import (
	"log"

	"goku.net/framework/network/http"
	"goku.net/utils"
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

func (executor *UserDetail) Execute() http.ResponseData {
	result, _ := utils.ToJSONString(executor.BodyData)
	log.Println("UserDetail", result)
	return executor.ResultOKData(executor.BodyData)
}
