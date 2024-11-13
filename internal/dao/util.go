package dao

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

const SelectForUpdate = "UPDATE"

func GormWhereEqualWithNotNil(session *gorm.DB, field string, v any) *gorm.DB {
	if reflect.ValueOf(v).IsNil() {
		return session
	}

	return session.Where(fmt.Sprintf("%s = ?", field), v)
}

func GormWhereInWithNotEmpty[T any](session *gorm.DB, field string, s []T) *gorm.DB {
	if 0 == len(s) {
		return session
	}

	return session.Where(fmt.Sprintf("%s in (?)", field), s)
}

func Lock(session *gorm.DB, lock bool) *gorm.DB {
	if !lock {
		return session
	}

	return session.Clauses(clause.Locking{Strength: SelectForUpdate})
}
