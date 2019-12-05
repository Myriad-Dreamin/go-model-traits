package dorm_crud_dao

import "github.com/Myriad-Dreamin/dorm"

func (model CRUDModel) UpdateFields(d ORMObject, fields []string) (int64, error) {
	return model.i.GetDormModel().Anchor(d).Select(fields...).UpdateFields()
}

func (model CRUDModel) UpdateFields_(db *dorm.DB, d ORMObject, fields []string) (int64, error) {
	return model.i.GetDormModel().Clone().FixDB(db).Anchor(d).Select(fields...).UpdateFields()
}

func (model CRUDModel) UpdateFields__(db dorm.SQLCommon, d ORMObject, fields []string) (int64, error) {
	return model.i.GetDormModel().Clone().FixSqlDB(db).Anchor(d).Select(fields...).UpdateFields()
}
