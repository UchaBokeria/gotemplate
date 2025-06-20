package dto

type ByID struct {
	ID int `param:"id"`
}

type Visibility struct {
	ByID
	Status bool `json:"status" form:"status" param:"status" query:"status"`
}
