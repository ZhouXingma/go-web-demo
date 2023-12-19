package userbll

import (
	"go-web-demo/internal/web/dal/userdao"
	"go-web-demo/internal/web/pkg/content"
	"go-web-demo/pkg/model"
)

// 添加用户
// 参数：
// bo: 添加用户的参数
// 返回：
// int64：数据库影响行
// error：异常
func AddUser(bo *AddUserBo) (int64, error) {
	addUserPo := bo.Trans2AddUserPo()
	// 开启事务
	tx, err := content.StartTx()
	if nil != err {
		return 0, err
	}
	// 关闭事务
	defer content.EndTxByError(tx, &err)
	// 执行操作
	return userdao.Insert(tx, addUserPo)
}

// 获取用户
// 参数：
// bo: 查找用户的参数
// 返回：
// userdao.UserPo： 获取用户的查询条件
// error：异常
func GetUser(bo *FindUserBo) (*userdao.UserPo, error) {
	findUserPo := bo.Trans2FindUserPo()
	userPo, err := userdao.Get(findUserPo)
	if nil != err {
		return nil, err
	}
	return userPo, nil
}

// 列表获取用户
// 参数：
// bo: 查找用户的参数
// 返回：
// []userdao.UserPo： 用户集合
// error：异常
func ListUser(bo *FindUserBo) ([]userdao.UserPo, error) {
	findUserPo := bo.Trans2FindUserPo()
	userPoList, err := userdao.List(findUserPo)
	if nil != err {
		return nil, err
	}
	return userPoList, nil
}

// 删除用户
// 参数：
// id: 用户id
// 返回：
// int64: 数据库影响行
// error：异常
func Remove(id string) (int64, error) {
	tx, err := content.StartTx()
	if nil != err {
		return 0, err
	}
	defer content.EndTxByError(tx, &err)
	return userdao.Delete(tx, id)
}

// 分页获取用户
// 参数：
// bo: 分页查询参数
// 返回：
// model.PageResult[userdao.UserPo]: 分页信息，包含了用户信息
// error：异常
func PageUser(bo *PageUserBo) (*model.PageResult[userdao.UserPo], error) {
	findUserPo := bo.Trans2FindUserPo()
	// 获取数量
	total, err := userdao.Count(findUserPo)
	if err != nil {
		return nil, err
	}
	poList, err := userdao.PageList(findUserPo, &bo.PageParam)
	if err != nil {
		return nil, err
	}
	pageResult := model.PageResult[userdao.UserPo]{}
	pageResult.PageInfo(&bo.PageParam, poList, total)
	return &pageResult, nil
}

// 更新用户信息
// 参数：
// bo: 用户信息
// 返回：
// int64：数据库影响行
// error：异常
func UpdateUser(bo *UpdateUserBO) (int64, error) {
	updatePo := bo.TransUpdatePo()
	tx, err := content.StartTx()
	if nil != err {
		return 0, err
	}
	defer content.EndTxByError(tx, &err)
	row, err := userdao.UpdateUser(tx, updatePo)
	if err != nil {
		return 0, err
	}
	return row, nil
}
