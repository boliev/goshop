package response

import "goshop/src/product"

// ProductType struct
type ProductType struct {
	Name   string         `json:"name"`
	Fields []ProductField `json:"fields"`
}

// ProductField struct
type ProductField struct {
	Name      string `json:"name"`
	FieldType string `json:"fieldType"`
}

// CreateFromDomain creates ProductType response from ProductType domain
func (ProductType) CreateFromDomain(domainProductType product.TypeInterface) *ProductType {
	var tps []ProductField

	for _, tp := range domainProductType.Fields() {
		tps = append(tps, ProductField{Name: tp.Name, FieldType: tp.FieldType.Name()})
	}

	return &ProductType{Name: domainProductType.Name(), Fields: tps}
}
