package dtos

type UpdateProductDto struct {
	ID          uint    `param:"id" validate:"required"`
	Name        string  `json:"name" form:"name" validate:"required,min=1,max=255"`
	CategoryIDs string  `json:"categoryIds" form:"categoryIds"`
	Price       float64 `json:"price" form:"price" validate:"required,min=0"`
	Stock       int     `json:"stock" form:"stock" validate:"required,min=0"`
	Status      string  `json:"status" form:"status" validate:"required,oneof=active inactive draft"`
	Description string  `json:"description" form:"description" validate:"max=5000"`
}
