package modeltraits

import (
	"github.com/Myriad-Dreamin/functional-go"
	"github.com/jinzhu/gorm"
)


type GORMTraitsInterface interface {
	GetGormDB() *gorm.DB
}

type GormModelTraits struct {
	i functional.BaseTraits
	GormDB *gorm.DB
}

func (g GormModelTraits) GetGormDB() *gorm.DB {
	return g.GormDB
}

func NewGormModelTraits(t functional.BaseTraits, gormDB *gorm.DB) GormModelTraits {
	return GormModelTraits{
		i: t,
		GormDB: gormDB,
	}
}
