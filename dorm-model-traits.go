package modeltraits

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/functional-go"
)

type DORMTraitsInterface interface {
	GetDormModel() *dorm.Model
	GetDormDB() *dorm.DB
}

type DormModelTraits struct {
	i         functional.BaseTraits
	DormDB    *dorm.DB
	DormModel *dorm.Model
}

func (g DormModelTraits) GetDormDB() *dorm.DB {
	return g.DormDB
}

func (traits DormModelTraits) GetDormModel() *dorm.Model {
	return traits.DormModel
}

func NewDormModelTraits(t functional.BaseTraits, dormDB *dorm.DB) DormModelTraits {
	return DormModelTraits{
		i:         t,
		DormDB:    dormDB,
		DormModel: new(dorm.Model),
	}
}
