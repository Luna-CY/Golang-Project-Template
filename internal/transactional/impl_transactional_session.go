package transactional

import "gorm.io/gorm"

func (cls *Transactional) Session() *gorm.DB {
	return cls.db
}
