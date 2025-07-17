package dtos

// OrderListDto for filtering orders based on frontend
type OrderListDto struct {
	Page        string   `query:"page" validate:"omitempty,numeric"`
	PageSize    string   `query:"pageSize" validate:"omitempty,numeric,max=100"`
	Search      string   `query:"search" validate:"max=255"`
	StatusMulti []string `query:"statusMulti"`
	DateMulti   []string `query:"dateMulti"`
	SortBy      string   `query:"sort_by" validate:"omitempty,oneof=order_number customer_name date total items"`
	SortOrder   string   `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
