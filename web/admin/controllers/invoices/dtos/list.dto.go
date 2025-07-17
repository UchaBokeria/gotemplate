package dtos

// InvoiceListDto for filtering invoices based on frontend
type InvoiceListDto struct {
	Page        string   `query:"page" validate:"omitempty,numeric"`
	PageSize    string   `query:"pageSize" validate:"omitempty,numeric,max=100"`
	Search      string   `query:"search" validate:"max=255"`
	OrderID     string   `query:"orderId"`
	StatusMulti []string `query:"statusMulti"`
	SortBy      string   `query:"sort_by" validate:"omitempty,oneof=invoice_number customer_name date due_date amount"`
	SortOrder   string   `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
