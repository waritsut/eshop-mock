package struct_templates

type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
	TotalRows  int64  `json:"totalRows"`
	TotalPages int    `json:"totalPages"`
	PrevPage   int    `json:"prevPage"`
	NextPage   int    `json:"nextPage"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Page > 1 {
		p.PrevPage = p.Page - 1
	} else {
		p.PrevPage = p.Page
	}

	if p.Page == p.TotalPages {
		p.NextPage = p.Page
	} else {
		p.NextPage = p.Page + 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id DESC"
	}
	return p.Sort
}
