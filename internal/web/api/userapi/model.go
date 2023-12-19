package userapi

import (
	"go-web-demo/internal/web/bll/userbll"
	"go-web-demo/internal/web/dal/userdao"
	"go-web-demo/pkg/commfunc"
	"go-web-demo/pkg/model"
	"time"
)

// 添加用户的参数
type AddUserParam struct {
	Name     string      `json:"name" form:"name" binding:"required" msg:"name 不能为空"`
	Sex      *int16      `json:"sex" form:"sex"`
	Birthday *model.Time `json:"birthday" form:"birthday"`
}

func (addUserParam AddUserParam) Trans2AddUserBo() *userbll.AddUserBo {
	return &userbll.AddUserBo{
		Id:       "",
		Name:     addUserParam.Name,
		Birthday: addUserParam.Birthday,
		Sex:      addUserParam.Sex,
	}
}

// 更新用户参数
type UpdateUserParam struct {
	Id       string     `json:"id"`
	Name     *string    `json:"name"`
	Sex      *int16     `json:"sex"`
	Birthday *time.Time `json:"birthday"`
}

func (updateUserParam UpdateUserParam) Trans2UpdateUserBo() *userbll.UpdateUserBO {
	return &userbll.UpdateUserBO{
		Id:       updateUserParam.Id,
		Sex:      updateUserParam.Sex,
		Name:     updateUserParam.Name,
		Birthday: updateUserParam.Birthday,
	}
}

// 查询用户的参数
type FindUserParam struct {
	Id   *string `json:"id" form:"id"`
	Name *string `json:"name" form:"name"`
	Sex  *int16  `json:"sex" form:"sex"`
}

func (findUserParam FindUserParam) Trans2FindUserBo() *userbll.FindUserBo {
	return &userbll.FindUserBo{
		Id:   findUserParam.Id,
		Name: findUserParam.Name,
		Sex:  findUserParam.Sex,
	}
}

// 分页获取用户的参数
type PageUserParam struct {
	model.PageParam         // 分页信息
	Id              *string `json:"id" form:"id"`
	Name            *string `json:"name" form:"name"`
	Sex             *int16  `json:"sex" form:"sex"`
}

func (pageUserParam PageUserParam) Trans2PageUserBo() *userbll.PageUserBo {
	return &userbll.PageUserBo{
		PageParam: pageUserParam.PageParam,
		Id:        pageUserParam.Id,
		Name:      pageUserParam.Name,
		Sex:       pageUserParam.Sex,
	}
}

// 移除用户的参数
type RemoveUserParam struct {
	Id string `json:"id" form:"id" binding:"required" msg:"id 不能为空"`
}

// 给前端的用户信息，不会全部字段都给
type UserInfo struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Sex        *int16      `json:"sex"`
	Birthday   *model.Time `json:"birthday"`
	GmtCreated model.Time  `json:"gmtCreated"`
}

// 从用户信息的po中填充用户信息
func (userInfo *UserInfo) fullBy(po *userdao.UserPo) {
	userInfo.Id = (*po).Id
	userInfo.Name = (*po).Name
	userInfo.Sex = commfunc.BoolHandle((*po).Sex.Valid, &((*po).Sex.Int16), nil)
	userInfo.Birthday = commfunc.BoolHandle(po.Birthday.Valid, model.TransTime2Time((*po).Birthday.Time), nil)
	userInfo.GmtCreated = model.Time(po.GmtCreated)
}
