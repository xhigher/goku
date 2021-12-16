package model

import (
	proto_model "goku.net/protos/model"
)

func GetUserDetail(mid int64) (data proto_model.UserDetailModel) {
	MyDB().First(&data, mid)
	return
}
