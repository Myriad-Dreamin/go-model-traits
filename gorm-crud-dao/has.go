package gorm_crud_dao

import "github.com/jinzhu/gorm"

func (model CRUDModel)  Has(id uint) (has bool, err error) {
	rdb := model.i.GetGormDB().First(model.replacement, id)
	err = rdb.Error
	if err == nil {
		has = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}
