package gorm_crud_dao



func (model CRUDModel) ID(id uint) (obj interface{}, err error) {
	obj = model.i.ObjectFactory()
	rdb := model.i.GetGormDB().First(obj, id)
	err = rdb.Error
	if rdb.RecordNotFound() {
		obj = nil
		err = nil
	}
	return
}

