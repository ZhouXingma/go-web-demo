package model

// 分页参数信息
type PageParam struct {
	PageNumber int64 `json:"pageNumber" form:"pageNumber" binding:"min=1"`
	PageSize   int64 `json:"pageSize" form:"pageSize" binding:"min=1"`
}

// 获取游标值
// 返回：
// int64： 游标值
func (page PageParam) GetCursor() int64 {
	return (page.PageNumber - 1) * page.PageSize
}

// 分页结果信息
type PageResult[T any] struct {
	PageParam
	Datas []T   `json:"datas" form:"datas"`
	Total int64 `json:"total" form:"total"`
}

// 组装分页结果信息
// 参数：
// param：分页参数
// datas: 分页数据集合
// total： 数据总量
func (result *PageResult[T]) PageInfo(param *PageParam, datas []T, total int64) {
	result.PageNumber = (*param).PageNumber
	result.PageSize = (*param).PageSize
	result.Datas = datas
	result.Total = total
}
