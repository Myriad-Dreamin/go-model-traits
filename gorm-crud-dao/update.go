package gorm_crud_dao

func (model CRUDModel) Update(d interface{}) (int64, error) {
	rdb := model.i.GetGormDB().Save(d)
	return rdb.RowsAffected, rdb.Error
}
