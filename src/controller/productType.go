package controller

import (
	response "goshop/src/Response"
	"goshop/src/product"
	"goshop/src/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProductType Create product type action
func CreateProductType(c *gin.Context) {
	var json request.ProductTypeRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fields []product.Field
	for _, field := range json.Fields {
		tp, error := product.NewFieldType(field.FieldType)
		if error != nil {
			c.JSON(400, response.BadRequestResponse{error.Error()})
			return
		}

		fields = append(fields, product.Field{Name: field.Name, FieldType: tp})
	}

	productType := product.Type{}.New(json.Name, fields)
	c.JSON(200, response.ProductType{}.CreateFromDomain(productType))
}

// GetProductTypes get available product types
func GetProductTypes(c *gin.Context) {
	c.JSON(200, response.ProductTypesList{product.AvailableFieldTypes()})
}
