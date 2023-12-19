package model

import (
	"time"
)

// 数据表的基本模型，几乎大多数的表都包含这些字段
type DalBasicModel struct {
	// 自增主键
	PkId uint64 `db:"pk_id"`
	// 业务id
	Id string `db:"id"`
	// 是否已经删除
	IsDeleted int16 `db:"is_deleted"`
	// 创建时间
	GmtCreated time.Time `db:"gmt_created"`
	// 修改时间
	GmtModified time.Time `db:"gmt_modified"`
	// 删除时间
	GmtDeleted time.Time `db:"gmt_deleted"`
}
