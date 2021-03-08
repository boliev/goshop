package product

import "fmt"

// TypeInterface Interface
type TypeInterface interface {
	Fields() []Field
	Name() string
}

// Type struct
type Type struct {
	name   string
	fields []Field
}

// Fields getter for ProductType
func (e Type) Fields() []Field {
	return e.fields
}

// Name getter for ProductType
func (e Type) Name() string {
	return e.name
}

// New ProductType constructor
func (e Type) New(name string, fields []Field) Type {
	return Type{
		name:   name,
		fields: fields,
	}
}

// Field struct
type Field struct {
	Name      string
	FieldType FieldType
}

// FieldType methods to manage field types
type FieldType struct {
	name string
}

// newFieldType creates new FieldType
func newFieldType(fieldType string) FieldType {
	return FieldType{"string"}
}

// Name getter
func (f FieldType) Name() string {
	return f.name
}

// NewFieldType creates new ProductFieldType
func NewFieldType(name string) (FieldType, error) {
	for _, available := range AvailableFieldTypes() {
		if available == name {
			return newFieldType(name), nil
		}
	}

	return FieldType{}, fmt.Errorf("Unknown field type %s", name)
}

// AvailableFieldTypes returns Available types for product fields
func AvailableFieldTypes() []string {
	return []string{"string", "int", "text", "return", "float", "file", "gallery"}
}
