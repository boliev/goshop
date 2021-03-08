package response

import "goshop/src/product"

// ProductType response
type ProductType struct {
	Name   string         `json:"name"`
	Fields []ProductField `json:"fields"`
}

// ProductField response
type ProductField struct {
	Name      string `json:"name"`
	FieldType string `json:"fieldType"`
}

// ProductTypesList response
type ProductTypesList struct {
	Types []string `json:"types"`
}

// CreateFromDomain creates ProductType response from ProductType domain
func (ProductType) CreateFromDomain(domainProductType product.TypeInterface) *ProductType {
	var tps []ProductField

	for _, tp := range domainProductType.Fields() {
		tps = append(tps, ProductField{Name: tp.Name, FieldType: tp.FieldType.Name()})
	}

	return &ProductType{Name: domainProductType.Name(), Fields: tps}
}
