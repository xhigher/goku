package model

import (
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
