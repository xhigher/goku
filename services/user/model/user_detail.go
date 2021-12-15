package model

import (
	proto "goku.net/protos/model"
)

func GetUserDetail(mid int64) (data proto.UserDetailModel) {
	MyDB().First(&data, mid)
	return
}
