package model

import "goku.net/framework/database"

type UserModel struct {
	*database.BaseModel
}

func (model *UserModel) DatabaseName() string {
	return "goku_user"
}
