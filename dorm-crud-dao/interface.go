//

package dorm_crud_dao

import (
	"github.com/Myriad-Dreamin/dorm"
)

// CRUDModelInterface Interface
type CRUDModelInterface interface {
	UpdateFields(d ORMObject, fields []string) (int64, error)
	UpdateFields_(db *dorm.DB, d ORMObject, fields []string) (int64, error)
	UpdateFields__(db dorm.SQLCommon, d ORMObject, fields []string) (int64, error)
}

