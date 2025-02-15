package resp

type PageModel struct {
	Total int64 `json:"total" form:"total" query:"total"`
	Start int64 `json:"start" form:"start" query:"start"`
	End   int64 `json:"end" form:"end" query:"end"`
	Index int   `json:"index" form:"index" query:"index"`
	Limit int   `json:"limit" form:"limit" query:"limit"`
	Desc  bool  `json:"desc" form:"desc" query:"desc"`
}

func NewPage(total, start, end int64, index, limit int, desc bool) *PageModel {
	p := &PageModel{
		Total: total,
		Start: start,
		End:   end,
		Index: index,
		Limit: limit,
		Desc:  desc,
	}
	return p
}

// NewPageFull 获取全部数据
func NewPageFull() *PageModel {
	p := &PageModel{
		Total: 0,
		Start: 0,
		End:   0,
		Index: 1,
		Limit: 0,
		Desc:  false,
	}
	return p
}

// NewPageCount 获取数据总量
func NewPageCount() *PageModel {
	p := &PageModel{
		Total: 0,
		Start: 0,
		End:   0,
		Index: 1,
		Limit: -1,
		Desc:  false,
	}
	return p
}
