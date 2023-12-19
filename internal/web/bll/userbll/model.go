package userbll

import (
	"go-web-demo/internal/web/dal/userdao"
	"go-web-demo/pkg/model"
	"time"
)

// 添加用户参数
type AddUserBo struct {
	Id          string
	Name        string
	Sex         *int16
	Birthday    *model.Time
	GmtCreated  *model.Time
	GmtModified *model.Time
}

func (m *AddUserBo) Trans2AddUserPo() *userdao.AddUserPo {
	return &userdao.AddUserPo{
		Id:          m.Id,
		Name:        m.Name,
		Sex:         model.NewSqlNullInt16(m.Sex),
		Birthday:    model.NewSqlNullTimeOfTime(m.Birthday),
		GmtCreated:  model.NewSqlNullTimeOfTime(m.GmtCreated),
		GmtModified: model.NewSqlNullTimeOfTime(m.GmtModified),
	}
}

// 查询用户参数
type FindUserBo struct {
	Id   *string
	Name *string
	Sex  *int16
}

func (m *FindUserBo) Trans2FindUserPo() *userdao.FindUserPo {
	return &userdao.FindUserPo{
		Id:   model.NewSqlNullString(m.Id),
		Name: model.NewSqlNullString(m.Name),
		Sex:  model.NewSqlNullInt16(m.Sex),
	}
}

// 分页查询用户参数
type PageUserBo struct {
	model.PageParam
	Id   *string
	Name *string
	Sex  *int16
}

func (m *PageUserBo) Trans2FindUserPo() *userdao.FindUserPo {
	return &userdao.FindUserPo{
		Id:   model.NewSqlNullString(m.Id),
		Name: model.NewSqlNullString(m.Name),
		Sex:  model.NewSqlNullInt16(m.Sex),
	}
}

// 更新用户参数
type UpdateUserBO struct {
	Id       string
	Name     *string
	Sex      *int16
	Birthday *time.Time
}

func (updateUserBo *UpdateUserBO) TransUpdatePo() *userdao.UpdatePo {
	return &userdao.UpdatePo{
		Id:       updateUserBo.Id,
		Name:     model.NewSqlNullString(updateUserBo.Name),
		Sex:      model.NewSqlNullInt16(updateUserBo.Sex),
		Birthday: model.NewSqlNullTime(updateUserBo.Birthday),
	}
}
