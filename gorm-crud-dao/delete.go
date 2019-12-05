package gorm_crud_dao

import "github.com/jinzhu/gorm"

func (model CRUDModel) Delete(d interface{}) (int64, error) {
	return model.Delete_(model.i.GetGormDB(), d)
}

func (model CRUDModel) Delete_(db *gorm.DB, d interface{}) (int64, error) {
	rdb := db.Delete(d)
	return rdb.RowsAffected, rdb.Error
}
