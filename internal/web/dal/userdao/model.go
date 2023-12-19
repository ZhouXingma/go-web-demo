package userdao

import (
	"database/sql"
	"go-web-demo/internal/web/model"
)

// 用户完整信息的数据po
type UserPo struct {
	model.DalBasicModel
	Name     string        `db:"name"`
	Sex      sql.NullInt16 `db:"sex"`      // null意味着允许为空
	Birthday sql.NullTime  `db:"birthday"` // null意味着允许为空
}

// 添加用户的参数
type AddUserPo struct {
	Id          string
	Name        string
	Sex         sql.NullInt16 // null意味着允许为空
	Birthday    sql.NullTime  // null意味着允许为空
	GmtCreated  sql.NullTime  // null意味着允许为空
	GmtModified sql.NullTime  // null意味着允许为空
}

// 查找用户的参数
type FindUserPo struct {
	Id       sql.NullString // null意味着允许为空
	Name     sql.NullString // null意味着允许为空
	Sex      sql.NullInt16  // null意味着允许为空
	Birthday sql.NullTime   // null意味着允许为空
}

// 更新用户的参数
type UpdatePo struct {
	Id       string
	Name     sql.NullString // null意味着允许为空
	Sex      sql.NullInt16  // null意味着允许为空
	Birthday sql.NullTime   // null意味着允许为空
}
