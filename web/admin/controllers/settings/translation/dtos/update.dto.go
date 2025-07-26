package dtos

type UpdateTranslationDto struct {
	ID       uint   `param:"id" form:"id" validate:"required"`
	Key      string `json:"key" form:"key" validate:"required,min=1,max=255"`
	Value    string `json:"value" form:"value" validate:"required,min=1,max=5000"`
	Language string `json:"language" form:"language" validate:"required,min=2,max=10"`
}
