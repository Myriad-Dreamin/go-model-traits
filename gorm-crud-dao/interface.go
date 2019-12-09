//

package gorm_crud_dao

import (
	"github.com/jinzhu/gorm"
)

// CRUDModelInterface Interface
type CRUDModelInterface interface {
	Create(d interface{}) (int64, error)
	Create_(db *gorm.DB, d interface{}) (int64, error)
	Delete(d interface{}) (int64, error)
	Delete_(db *gorm.DB, d interface{}) (int64, error)
	Filter(f *Filter, objs interface{}) error
	Filter_(db *gorm.DB, f *Filter, objs interface{}) error
	Has(id uint) (bool, error)
	Has_(db *gorm.DB, id uint) (has bool, err error)
	ID(id uint) (interface{}, error)
	ID_(db *gorm.DB, id uint) (obj interface{}, err error)
	Update(d interface{}) (int64, error)
	Update_(db *gorm.DB, d interface{}) (int64, error)
	Where1(template string) func(id interface{}) (interface{}, error)
	Where1_(template string) func(db *gorm.DB, id interface{}) (interface{}, error)
}
