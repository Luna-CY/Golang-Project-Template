package model

type ExampleEnumFieldType int

const (
	ExampleEnumFieldTypeA = ExampleEnumFieldType(1) // Example enum field A
	ExampleEnumFieldTypeB = ExampleEnumFieldType(2) // Example enum field B
	ExampleEnumFieldTypeC = ExampleEnumFieldType(3) // Example enum field C
)

// //go:generate go run ../cmd/main/main.go generate dao --save --take-by Field1=string=,Field4=model.ExampleEnumFieldType=model.ExampleEnumFieldType(0) --delete-by Field1=string= --batch-take-by Field1=string

type Example struct {
	Model

	Field1 *string               `gorm:"type:varchar(255);not null;default:''"` // Example field 1
	Field2 *uint64               `gorm:"type:uint;not null;default:0"`          // Example field 2
	Field3 *bool                 `gorm:"type:bool;not null;default:false"`      // Example field 3
	Field4 *ExampleEnumFieldType `gorm:"type:uint;not null;default:1"`          // Example field 4
}
