package gorm_crud_dao

func (model CRUDModel) Create(d interface{}) (int64, error) {
	rdb := model.i.GetGormDB().Create(d)
	return rdb.RowsAffected, rdb.Error
}
