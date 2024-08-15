package pagination

type PaginationRequest struct {
	Limit  int    `json:"limit,omitempty" query:"limit" example:"10"`
	Page   int    `json:"page,omitempty" query:"page" example:"1"`
	Search string `json:"search,omitempty" query:"search"`
	Sort   string `json:"sort,omitempty" query:"sort" example:"id DESC"`
	Total  int64  `json:"total,omitempty" example:"100"`
}

func (p *PaginationRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationRequest) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *PaginationRequest) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PaginationRequest) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}
