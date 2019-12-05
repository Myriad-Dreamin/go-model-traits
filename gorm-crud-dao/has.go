package gorm_crud_dao

import "github.com/jinzhu/gorm"

func (model CRUDModel)  Has(id uint) (bool, error) {
	return model.Has_(model.i.GetGormDB(), id)
}

func (model CRUDModel)  Has_(db *gorm.DB, id uint) (has bool, err error) {
	rdb := db.First(model.replacement, id)
	err = rdb.Error
	if err == nil {
		has = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}
