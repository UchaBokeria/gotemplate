package dtos

type ProductFilterDto struct {
	Page      int     `query:"page" default:"1"`
	Limit     int     `query:"limit" default:"12"`
	Search    string  `query:"search" default:""`
	Category  string  `query:"category" default:""`
	MinPrice  float64 `query:"min_price" default:"0"`
	MaxPrice  float64 `query:"max_price" default:"0"`
	SortBy    string  `query:"sort_by" default:"created_at"`
	SortOrder string  `query:"sort_order" default:"desc"`
}

type ProductSortDto struct {
	SortBy    string `query:"sort_by"`
	SortOrder string `query:"sort_order"`
}

type ProductSearchDto struct {
	Search   string `query:"search"`
	Category string `query:"category"`
}
