package dorm_crud_dao

func (model CRUDModel) UpdateFields(d ORMObject, fields []string) (int64, error) {
	return model.i.GetDormModel().Anchor(d).Select(fields...).UpdateFields()
}
