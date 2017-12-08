package page

type Page struct {
	DefaultPageSize int64
	Total           int64
	Current         int64
	PageSize        int64
	TotalPage       int64
	Data            []interface{}
}

func NewPage(total, current, pageSize, totalPage int64, data []interface{}) (page Page, err error) {
	page = Page{
		DefaultPageSize: getDefaultPageSize(),
		Total:           total,
		Current:         current,
		PageSize:        pageSize,
		TotalPage:       totalPage,
		Data:            data,
	}
	return
}

func getDefaultPageSize() int64 {
	return 20
}

func (p Page) GetTotalPage() int64 {
	if p.Total == 0 {
		return 0
	}
	if p.TotalPage == 0 {
		if p.Total%p.PageSize == 0 {
			p.TotalPage = p.Total / p.PageSize
		} else {
			p.TotalPage = p.Total/p.PageSize + 1
		}
	}
	return p.TotalPage
}
