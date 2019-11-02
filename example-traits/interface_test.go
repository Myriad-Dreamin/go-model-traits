package mytraits

import (
	modeltraits "github.com/Myriad-Dreamin/go-model-traits"
	dorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/dorm-crud-dao"
	gorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

var c Interface = &ModelTraits{}

var _ modeltraits.BaseTraitsInterface = c
var _ modeltraits.GORMTraitsInterface = c
var _ modeltraits.DORMTraitsInterface = c
var _ gorm_crud_dao.CRUDModelInterface = c
var _ dorm_crud_dao.CRUDModelInterface = c

