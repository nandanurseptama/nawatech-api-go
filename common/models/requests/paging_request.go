package requests

type PagingRequest struct {
	ItemPerPage uint     `json:"item_per_page"`
	Page        uint     `json:"page"`
	SortBy      []SortBy `json:"sort_by"`
}
