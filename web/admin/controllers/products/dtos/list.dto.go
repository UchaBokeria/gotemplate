package dtos

type ProductListDto struct {
	Search      string   `query:"search"`
	CategoryIDs string   `query:"categoryIds"`
	StatusMulti []string `query:"statusMulti"`
	Page        int      `query:"page,default:1"`
	Limit       int      `query:"limit,default:20"`
}
