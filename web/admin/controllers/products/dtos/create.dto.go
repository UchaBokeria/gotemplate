package dtos

type CreateProductDto struct {
	Name        string  `json:"name" form:"name" validate:"required,min=1,max=255"`
	Category    string  `json:"category" form:"category" validate:"required,min=1,max=100"`
	Price       float64 `json:"price" form:"price" validate:"required,min=0"`
	Stock       int     `json:"stock" form:"stock" validate:"required,min=0"`
	Status      string  `json:"status" form:"status" validate:"required,oneof=active inactive draft"`
	Description string  `json:"description" form:"description" validate:"max=5000"`
}
