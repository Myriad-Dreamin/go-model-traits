package gorm_crud_dao

import "github.com/jinzhu/gorm"

func (model CRUDModel) Where1(template string) func(id interface{}) (interface{}, error) {
	return func(id interface{}) (object interface{}, err error) {
		object = model.i.ObjectFactory()
		rdb := model.i.GetGormDB().Where(template, id).Find(object)
		err = rdb.Error
		if rdb.RecordNotFound() {
			object = nil
			err = nil
		}
		return
	}
}

func (model CRUDModel) Where1_(template string) func(db *gorm.DB, id interface{}) (interface{}, error) {
	return func(db *gorm.DB, id interface{}) (object interface{}, err error) {
		object = model.i.ObjectFactory()
		rdb := db.Where(template, id).Find(object)
		err = rdb.Error
		if rdb.RecordNotFound() {
			object = nil
			err = nil
		}
		return
	}
}

