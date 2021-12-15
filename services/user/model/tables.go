package model

type UserDetailModel struct {
	*UserModel
}

func (model *UserDetailModel) TableName() string {
	return "user_detail"
}

func (model *UserDetailModel) GetDetail(mid int64) (data UserDetailModel) {
	model.Go().First(&data, mid)
	return
}
