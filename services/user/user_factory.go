package user

import "goku.net/framework/network/http"

type UserFactory struct {
}

func (factory *UserFactory) Create(version int, action string) http.Executor {
	switch action {
	case "detail":
		return &UserDetail{&http.BaseExecutor{
			BodyData: &UserDetailParam{},
			Version:  version,
		}}
	case "address_list":
		return &UserAddressList{&http.BaseExecutor{
			BodyData: &UserAddressListParam{},
			Version:  version,
		}}
	case "address_update":
		return &UserAddressUpdate{&http.BaseExecutor{
			BodyData: &UserAddressUpdateParam{},
			Version:  version,
		}}
	default:
		return nil
	}
}
