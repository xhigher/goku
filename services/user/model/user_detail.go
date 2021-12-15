package model

import (
	proto "goku.net/protos/model"
)

type UserDetailModel struct {
	*UserModel
	*proto.UserDetailTable
}

func (model *UserDetailModel) GetDetail(mid int64) (data proto.UserDetailTable) {
	model.Go().First(&data, mid)
	return
}
