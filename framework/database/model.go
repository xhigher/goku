package database

import (
	"gorm.io/gorm"
)

type Model interface {
	DatabaseName() string
	TableName() string
}

type BaseModel struct {
	db *gorm.DB
}

func (model *BaseModel) DatabaseName() string {
	return ""
}

func (model *BaseModel) Go() *gorm.DB {
	db := GetDB(model.DatabaseName())
	if db != nil {
		model.db = db
		return model.db.Model(model)
	}
	return nil
}
