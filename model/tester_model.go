package model

//go:generate go run ../cmd/main/main.go generate dao --table tester_models --save --take-by Id=uint64=0 --batch-take-by Id=uint64

type TesterModel struct {
	Model

	Field1 *string `gorm:"type:varchar(255);not null;default:''"` // Example field 1
	Field2 *uint64 `gorm:"type:uint;not null;default:0"`          // Example field 2
}

func (cls *TesterModel) TableName() string {
	return "tester_models"
}
