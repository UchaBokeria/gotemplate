package dtos

type ProductListDto struct {
	Search        string   `query:"search"`
	CategoryMulti []string `query:"categoryMulti"`
	StatusMulti   []string `query:"statusMulti"`
	Page          int      `query:"page,default:1"`
	Limit         int      `query:"limit,default:20"`
}
