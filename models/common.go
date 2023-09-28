package models

type PaginationData struct {
	Limit  int    `json:"per_page"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Sort   string `json:"sort"`
	Total  int64  `json:"total"`
}

type PaginationResponse struct {
	Pagination PaginationData           `json:"pagination"`
	Data       []map[string]interface{} `json:"data"`
}

type IndexRequest struct {
	Search   string `json:"search"`
	Page     string `json:"page"`
	PerPage  string `json:"per_page"`
	Sort     string `json:"sort"`
	SortMode string `json:"mode"`
}
