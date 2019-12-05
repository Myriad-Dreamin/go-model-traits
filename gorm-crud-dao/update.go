package gorm_crud_dao

import "github.com/jinzhu/gorm"

func (model CRUDModel) Update(d interface{}) (int64, error) {
	return model.Update_(model.i.GetGormDB(), d)
}

func (model CRUDModel) Update_(db *gorm.DB, d interface{}) (int64, error) {
	rdb := db.Save(d)
	return rdb.RowsAffected, rdb.Error
}
