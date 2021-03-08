package request

// ProductTypeRequest DTO
type ProductTypeRequest struct {
	Name   string                    `form:"name" json:"name" binding:"required"`
	Fields []ProductTypeFieldRequest `form:"fields" json:"fields" binding:"required"`
}

// ProductTypeFieldRequest DTO
type ProductTypeFieldRequest struct {
	Name      string `form:"name" json:"name" binding:"required"`
	FieldType string `form:"fieldType" json:"fieldType" binding:"required"`
}
