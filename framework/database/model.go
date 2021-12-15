package database

import (
	"go.uber.org/zap"
	"goku.net/framework/commons"
	"gorm.io/gorm"
)

type BaseModel struct {
	DBName string
}

func (model *BaseModel) GormDB() *gorm.DB {
	commons.Logger().Info("BaseModel.GormDB()", zap.String("dbName", model.DBName))
	return GetDB(model.DBName)
}
