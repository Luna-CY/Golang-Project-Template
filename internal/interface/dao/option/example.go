package option

import (
	"github.com/Luna-CY/Golang-Project-Template/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ExampleOption Example 选项
// joinTables 已经连接的表名字典，避免重复连接
type ExampleOption func(session *gorm.DB, joinTables map[string]struct{}) *gorm.DB

// ExampleOptionWithLock 选项：加锁查询
func ExampleOptionWithLock() ExampleOption {
	return func(session *gorm.DB, joinTables map[string]struct{}) *gorm.DB {
		return session.Clauses(clause.Locking{Strength: SelectForUpdate})
	}
}

// ExampleOptionWithField4 选项：根据 field4 过滤
// 支持单值或多值
func ExampleOptionWithField4(field4 ...model.ExampleEnumFieldType) ExampleOption {
	return func(session *gorm.DB, joinTables map[string]struct{}) *gorm.DB {
		if 0 == len(field4) {
			return session
		}

		return session.Where("examples.field4 IN ?", field4)
	}
}

// ExampleOptionWithField4Not 选项：根据 field4 不等于给定值过滤
// 支持单值或多值
func ExampleOptionWithField4Not(field4 ...model.ExampleEnumFieldType) ExampleOption {
	return func(session *gorm.DB, joinTables map[string]struct{}) *gorm.DB {
		if 0 == len(field4) {
			return session
		}

		return session.Where("examples.field4 NOT IN ?", field4)
	}
}

// ExampleOptionWithOrderDefault 选项：默认排序
func ExampleOptionWithOrderDefault() ExampleOption {
	return func(session *gorm.DB, joinTables map[string]struct{}) *gorm.DB {
		return session.Order("examples.id DESC")
	}
}
