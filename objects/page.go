package objects

// The Page object contains the response page information.
type Page struct {
	// The page retrieved or, in the case of a request parameter, the specific page
	// requested.
	Page *int `json:"page,omitempty"`

	// The total pages available.
	TotalPages *int `json:"totalPages,omitempty"`

	// The number of objects on this page.
	PageSize *int `json:"pageSize,omitempty"`

	// The total number of objects available.
	Total *int `json:"total,omitempty"`
}
