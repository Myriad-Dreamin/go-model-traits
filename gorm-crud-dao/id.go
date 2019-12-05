package gorm_crud_dao

import "github.com/jinzhu/gorm"

func (model CRUDModel) ID(id uint) (interface{}, error) {
	return model.ID_(model.i.GetGormDB(), id)
}

func (model CRUDModel) ID_(db *gorm.DB, id uint) (obj interface{}, err error) {
	obj = model.i.ObjectFactory()
	rdb := db.First(obj, id)
	err = rdb.Error
	if rdb.RecordNotFound() {
		obj = nil
		err = nil
	}
	return
}

