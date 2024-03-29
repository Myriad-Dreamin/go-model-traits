package gorm_crud_dao

import "github.com/jinzhu/gorm"

type Filter struct {
	Order string `json:"order" form:"order"`
	Page int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
	BeforeID uint `json:"before_id" form:"before_id"`
}

func FilterOption(db *gorm.DB, f *Filter) *gorm.DB {
	db = PageOption(db, f.Page, f.PageSize)
	if len(f.Order) != 0 {
		db = db.Order(f.Order, false)
	}
	if f.BeforeID != 0 {
		db = db.Where("id <= ?", f.BeforeID)
	}
	return db
}

func (model CRUDModel) Filter(f *Filter, objs interface{}) error {
	return model.Filter_(model.i.GetGormDB(), f, objs)
}

func (model CRUDModel) Filter_(db *gorm.DB, f *Filter, objs interface{}) error {
	return FilterOption(db, f).Find(objs).Error
}

