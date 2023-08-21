package datastore

type Page struct {
	Data       any `json:"data"`
	PageNo     int `json:"pageNo"`
	PageSize   int `json:"pageSize"`
	TotalPage  int `json:"totalPage"`
	TotalCount int `json:"totalCount"`
}

func NewPage(data any, no int, size int, totalCount int) Page {
	return Page{
		Data:       data,
		PageNo:     no,
		PageSize:   size,
		TotalCount: totalCount,
		TotalPage:  TotalPage(totalCount, size),
	}
}

func TotalPage(totalCount int, size int) int {
	carry := 0
	if totalCount%size != 0 {
		carry = 1
	}
	pageCal := totalCount / size
	return pageCal + carry

}
