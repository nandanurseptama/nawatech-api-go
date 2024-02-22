package responses

// Paging Response
type PagingResponse[T comparable] struct {
	// Current page index
	CurrentPage uint `json:"current_page"`

	// Next page index
	//
	// if CurrentPage is equal to TotalPage
	//
	// NextPage should be nil
	NextPage *uint `json:"next_page,omitempty"`

	// Next page path
	//
	// if CurrentPage is equal to TotalPage
	//
	// NextPagePath should be nil
	NextPagePath *string `json:"next_page_path,omitempty"`

	// Prev page index
	//
	// if CurrentPage is 1
	//
	// PrevPage should be nil
	PrevPage *uint `json:"prev_page,omitempty"`

	// Prev page path
	//
	// if CurrentPage is equal to 1
	//
	// PrevPagePath should be nil
	PrevPagePath *string `json:"prev_page_path,omitempty"`

	// Results is list of data with type T
	//
	// When there is no data in data source
	//
	// Just return empty array
	Results []T `json:"results"`

	// Total record in data source
	TotalData uint `json:"total_data"`

	// TotalItemPerPage is total item per page that client requested
	TotalItemPerPage uint `json:"total_item_per_page"`

	// TotalPage in data source
	TotalPage uint `json:"total_page"`
}
