package dorm_crud_dao

import (
	"github.com/Myriad-Dreamin/dorm"
	modeltraits "github.com/Myriad-Dreamin/go-model-traits"
)

type ORMObject interface {
	dorm.ORMObject
}

type CRUDModelOperationInterface interface {
	modeltraits.DORMTraitsInterface
	modeltraits.BaseTraitsInterface
}

type CRUDModelInterface interface {
	UpdateFields(d ORMObject, fields []string) (int64, error)
}

type CRUDModel struct {
	i CRUDModelOperationInterface
}

func NewCRUDModel(t CRUDModelOperationInterface) CRUDModel {
	return CRUDModel{
		i: t,
	}
}
