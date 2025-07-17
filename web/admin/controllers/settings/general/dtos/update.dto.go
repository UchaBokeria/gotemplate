package dtos

type UpdateGeneralSettingsDto struct {
	CompanyName string `json:"companyName" form:"companyName" validate:"required,min=1,max=255"`
	Address     string `json:"address" form:"address" validate:"max=1000"`
	Phone       string `json:"phone" form:"phone" validate:"max=50"`
	Email       string `json:"email" form:"email" validate:"email,max=255"`
	Facebook    string `json:"facebook" form:"facebook" validate:"url,max=500"`
	Instagram   string `json:"instagram" form:"instagram" validate:"url,max=500"`
	TikTok      string `json:"tiktok" form:"tiktok" validate:"url,max=500"`
	AboutUs     string `json:"aboutUs" form:"aboutUs" validate:"max=5000"`
}
