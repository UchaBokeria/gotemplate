package dtos

type UpdateCategoryDto struct {
	ID          uint   `json:"id" form:"id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Slug        string `json:"slug" form:"slug" validate:"required,min=1,max=255"`
	Description string `json:"description" form:"description" validate:"max=5000"`
	ParentID    *uint  `json:"parentId" form:"parentId"`
	SortOrder   int    `json:"sortOrder" form:"sortOrder" validate:"min=0"`
	IsActive    bool   `json:"isActive" form:"isActive"`
}
