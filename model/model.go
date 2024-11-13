package model

import "gorm.io/plugin/soft_delete"

type Model struct {
	Id         uint64                 `gorm:"type:uint;primary_key;auto_increment;not null"` // ID
	CreateTime *int64                 `gorm:"type:uint;not null"`                            // created time
	UpdateTime *int64                 `gorm:"type:uint;not null"`                            // last updated time
	DeleteTime *soft_delete.DeletedAt `gorm:"type:uint;not null;default:0"`                  // deleted time
}
