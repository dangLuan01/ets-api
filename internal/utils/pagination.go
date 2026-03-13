package utils

type Pagination struct {
	Page 			int32 `json:"page"`
	Limit 			int32 `json:"limit"`
	TotalRecords 	int32 `json:"total_records"`
	TotalPages 		int32 `json:"total_pages"`
	HasNext 		bool  `json:"has_next"`
	HasPrev 		bool  `json:"has_prev"`
}

func NewPagination(page int32, limit int32, totalRecords int64) *Pagination {

	totalPages := (totalRecords + int64(limit) - 1) / int64(limit)
	
	return &Pagination{
		Page:  page,
		Limit: limit,
		TotalRecords: int32(totalRecords),
		TotalPages: int32(totalPages),
		HasNext: page < int32(totalPages),
		HasPrev: page > 1,
	}
}

func NewPaginationResponse(page int32, limit int32, totalRecords int64, data any) map[string]any {
	return map[string]any {
		"response": data,
		"pagination": NewPagination(page, limit, totalRecords),
	}
}