package requests

type SortDirection string

const (
	Asc  SortDirection = "ASC"
	Desc SortDirection = "DESC"
)

type SortBy struct {
	FieldName     string        `json:"field_name"`
	SortDirection SortDirection `json:"direction"`
}
