package dtos

type UpdateCategoryDto struct {
	ID          uint   `param:"id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Label       string `json:"label" form:"label" validate:"max=255"`
	Description string `json:"description" form:"description" validate:"max=5000"`
	SortOrder   int    `json:"sortOrder" form:"sortOrder" validate:"min=0"`
	IsActive    bool   `json:"isActive" form:"isActive"`
}
