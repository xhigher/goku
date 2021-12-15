package user

import "goku.net/framework/network/http"

type UserFactory struct {
}

func (factory *UserFactory) Create(version int, action string) http.Executor {
	switch action {
	case "detail":
		return &UserDetail{&http.BaseExecutor{
			BodyData: &UserDetailParam{},
		}}

	default:
		return nil
	}
}
