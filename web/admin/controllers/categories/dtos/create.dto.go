package dtos

type CreateCategoryDto struct {
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Label       string `json:"label" form:"label" validate:"max=255"`
	Description string `json:"description" form:"description" validate:"max=5000"`
	ParentID    string `json:"parentId" form:"parentId"`
	SortOrder   int    `json:"sortOrder" form:"sortOrder" validate:"min=0"`
	IsActive    bool   `json:"isActive" form:"isActive"`
}
