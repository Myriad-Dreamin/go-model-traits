package mytraits

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/functional-go"
	modeltraits "github.com/Myriad-Dreamin/go-model-traits"
	dorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/dorm-crud-dao"
	gorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/jinzhu/gorm"
)

type ORMObject interface {
	dorm_crud_dao.ORMObject
}

type Filter = gorm_crud_dao.Filter


type ModelInterface interface {
	functional.BaseTraitsInterface
	modeltraits.GORMTraitsInterface
	modeltraits.DORMTraitsInterface
	dorm_crud_dao.CRUDModelInterface
	gorm_crud_dao.CRUDModelInterface
}

type Interface interface {
	ModelInterface
	modeltraits.DatabaseTraitsInterface
}

type Traits = Interface


func NewTraits(t dorm.ORMObject, gormDB *gorm.DB, dormDB *dorm.DB) *ModelTraits {
	model := NewModelTraits(t, gormDB, dormDB)
	return &model
}

func NewTraitsInterface(t dorm.ORMObject, gormDB *gorm.DB, dormDB *dorm.DB) Interface {
	return NewTraits(t, gormDB, dormDB)
}

func NewObjectModelTraits(t dorm.ORMObject, gormDB *gorm.DB, dormDB *dorm.DB) ObjectModelTraits {
	tr := ObjectModelTraits{}
	tr.ModelTraits = NewModelTraits(t, gormDB, dormDB)
	return tr
}

