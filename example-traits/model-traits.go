package mytraits

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/functional-go"
	modeltraits "github.com/Myriad-Dreamin/go-model-traits"
	dorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/dorm-crud-dao"
	gorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/jinzhu/gorm"
)

type GormCRUDModel = gorm_crud_dao.CRUDModel

type DormCRUDModel = dorm_crud_dao.CRUDModel

type ModelTraits struct {
	functional.BaseTraits
	modeltraits.GormModelTraits
	modeltraits.DormModelTraits
	GormCRUDModel
	DormCRUDModel
}

func NewModelTraits(t dorm.ORMObject, gormDB *gorm.DB, dormDB *dorm.DB) ModelTraits {
	tr := ModelTraits{}
	tr.BaseTraits = functional.NewBaseTraits(t)
	tr.GormModelTraits = modeltraits.NewGormModelTraits(tr.BaseTraits, gormDB)
	tr.DormModelTraits = modeltraits.NewDormModelTraits(tr.BaseTraits, dormDB)
	tr.GormCRUDModel = gorm_crud_dao.NewCRUDModel(tr)
	tr.DormCRUDModel = dorm_crud_dao.NewCRUDModel(tr)
	return tr
}

func (traits *ModelTraits) Migrate() error {
	err := traits.GormDB.AutoMigrate(traits.ObjectFactory()).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	model, err := traits.DormDB.Model(traits.ObjectFactory().(dorm.ORMObject))
	if err != nil {
		return err
	}
	*traits.DormModel = *model
	return nil
}

