package userapi

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/internal/web/bll/userbll"
	"go-web-demo/internal/web/pkg/loginfo"
	"go-web-demo/pkg/binding"
	"go-web-demo/pkg/idgenerater"
	"go-web-demo/pkg/model"
	"go-web-demo/pkg/response"
)

// 获取用户
func GetUser(c *gin.Context) {
	var findUserParam FindUserParam
	if err := c.ShouldBind(&findUserParam); err != nil {
		response.Fail(c, binding.GetBindingErrorMsg(err, &findUserParam))
		return
	}
	loginfo.Info("获取用户，入参：", findUserParam)
	findUserBo := findUserParam.Trans2FindUserBo()
	userPo, err := userbll.GetUser(findUserBo)
	if nil != err {
		loginfo.Log.Error("获取用户失败:", err.Error())
		response.Fail(c, "获取用户失败！")
		return
	}
	userInfo := UserInfo{}
	userInfo.fullBy(userPo)
	response.Success(c, userInfo)
}

// 添加用户
func AddUser(c *gin.Context) {
	var addUserParam AddUserParam
	if err := c.ShouldBind(&addUserParam); err != nil {
		response.Fail(c, binding.GetBindingErrorMsg(err, &addUserParam))
		return
	}
	loginfo.Info("添加用户，入参：", addUserParam)
	addUserBo := addUserParam.Trans2AddUserBo()
	(*addUserBo).Id = idgenerater.CreateUlid()
	_, e := userbll.AddUser(addUserBo)
	if nil != e {
		response.Fail(c, "添加用户失败！")
		return
	}
	response.Success(c, addUserBo.Id)
}

// 列表获取
func ListUser(c *gin.Context) {
	var findUserParam FindUserParam
	if err := c.ShouldBind(&findUserParam); err != nil {
		response.Fail(c, binding.GetBindingErrorMsg(err, &findUserParam))
		return
	}
	loginfo.Info("获取用户，入参：", findUserParam)
	getUserBo := findUserParam.Trans2FindUserBo()
	userPoList, err := userbll.ListUser(getUserBo)
	if nil != err {
		loginfo.Log.Error("获取用户失败:", err.Error())
		response.Fail(c, "获取用户失败！")
		return
	}
	count := len(userPoList)
	if count <= 0 {
		response.Success(c, userPoList)
		return
	}
	userInfoList := make([]UserInfo, 0)
	for _, userPoItem := range userPoList {
		var userInfo UserInfo
		userInfo.fullBy(&userPoItem)
		userInfoList = append(userInfoList, userInfo)
	}
	response.Success(c, userInfoList)
}

// 分页获取
func PageUser(c *gin.Context) {
	var pageUserparam PageUserParam
	if err := c.ShouldBind(&pageUserparam); err != nil {
		response.Fail(c, binding.GetBindingErrorMsg(err, &pageUserparam))
		return
	}
	loginfo.Info("分页获取用户，入参：", pageUserparam)
	pageUserBo := pageUserparam.Trans2PageUserBo()
	pageResult, err := userbll.PageUser(pageUserBo)
	if nil != err {
		loginfo.Log.Error("分页获取用户失败:", err.Error())
		response.Fail(c, "分页获取用户失败！")
		return
	}
	if len(pageResult.Datas) <= 0 {
		response.Success(c, pageResult)
		return
	}
	userInfoList := make([]UserInfo, 0)
	for _, userPoItem := range pageResult.Datas {
		var userInfo UserInfo
		userInfo.fullBy(&userPoItem)
		userInfoList = append(userInfoList, userInfo)
	}
	pageResultInfo := model.PageResult[UserInfo]{}
	pageResultInfo.PageInfo(&pageUserparam.PageParam, userInfoList, pageResult.Total)
	response.Success(c, pageResultInfo)

}

// 移除用户
func RemoveUser(c *gin.Context) {
	var removeUserParam RemoveUserParam
	if err := c.ShouldBind(&removeUserParam); err != nil {
		response.Fail(c, binding.GetBindingErrorMsg(err, &removeUserParam))
		return
	}
	loginfo.Info("删除用户，入参：", removeUserParam)
	_, e := userbll.Remove(removeUserParam.Id)
	if nil != e {
		response.Fail(c, "删除用户失败！")
		return
	}
	response.Success(c, nil)
}

// 更新用户
func UpdateUser(c *gin.Context) {
	var updateUserParam UpdateUserParam
	if err := c.ShouldBind(&updateUserParam); err != nil {
		response.Fail(c, binding.GetBindingErrorMsg(err, &updateUserParam))
		return
	}
	loginfo.Info("更新用户，入参：", updateUserParam)
	updateUserBo := updateUserParam.Trans2UpdateUserBo()
	_, err := userbll.UpdateUser(updateUserBo)
	if nil != err {
		loginfo.Log.Error("更新用户失败:", err.Error())
		response.Fail(c, "更新用户失败！")
		return
	}
	response.Success(c, nil)
}
