package model

import (
	"github.com/hashicorp/consul/api"
	"goku.net/framework/database"
	"gorm.io/gorm"
)

type UserModel struct {
	database.BaseModel
}

func MyDB() *gorm.DB {
	model := &UserModel{
		database.BaseModel{
			DBName: "goku_user",
		},
	}
	return model.GormDB()
}
func f() {
	_ = api.ACL{}

}
