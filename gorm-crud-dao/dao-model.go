package gorm_crud_dao

import modeltraits "github.com/Myriad-Dreamin/go-model-traits"

type CRUDModelInterface interface {
	Create(d interface{}) (int64, error)
	Has(id uint) (has bool, err error)
	Update(d interface{}) (int64, error)
	Delete(d interface{}) (int64, error)
	ID(id uint) (obj interface{}, err error)
	Filter(f *Filter) (objs interface{}, err error)
	Where1(template string) func(id interface{}) (interface{}, error)
}

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
		i: t,
		replacement: t.ObjectFactory(),
	}
}