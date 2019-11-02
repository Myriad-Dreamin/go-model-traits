package gorm_crud_dao

func (model CRUDModel) Delete(d interface{}) (int64, error) {
	rdb := model.i.GetGormDB().Delete(d)
	return rdb.RowsAffected, rdb.Error
}
