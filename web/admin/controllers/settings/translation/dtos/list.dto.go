package dtos

type TranslationListDto struct {
	Search   string `query:"search"`
	Language string `query:"language"`
	Page     int    `query:"page,default:1"`
	Limit    int    `query:"limit,default:20"`
}
