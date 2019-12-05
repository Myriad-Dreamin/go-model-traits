package gorm_crud_dao

import (
	"github.com/jinzhu/gorm"
)

func (model CRUDModel) Create(d interface{}) (int64, error) {
	return model.Create_(model.i.GetGormDB(), d)
}

func (model CRUDModel) Create_(db *gorm.DB, d interface{}) (int64, error) {
	rdb := db.Create(d)
	return rdb.RowsAffected, rdb.Error
}

