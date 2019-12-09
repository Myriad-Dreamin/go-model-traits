package gorm_crud_dao

import modeltraits "github.com/Myriad-Dreamin/go-model-traits"


type CRUDModelOperationInterface interface {
	modeltraits.GORMTraitsInterface
	modeltraits.BaseTraitsInterface
}

type CRUDModel struct {
	i           CRUDModelOperationInterface
	replacement interface{}
}

func NewCRUDModel(t CRUDModelOperationInterface) CRUDModel {
	return CRUDModel{
		i:           t,
		replacement: t.ObjectFactory(),
	}
}
