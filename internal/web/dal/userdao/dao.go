package userdao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-web-demo/internal/web/pkg/content"
	"go-web-demo/internal/web/pkg/loginfo"
	"go-web-demo/pkg/commfunc"
	model2 "go-web-demo/pkg/model"
	"strings"
	"time"
)

// 添加用户
// 参数：
// tx：事务操作对象
// po: 新增的用户信息
// 返回：
// int64: 数据库影响行
// error: 操作的异常
func Insert(tx *sqlx.Tx, po *AddUserPo) (int64, error) {
	// 默认时间
	nowTime := time.Now()
	// sql语句
	sqlStr := "insert into user(id, name, sex, birthday, is_deleted, gmt_created, gmt_modified) values (?,?,?,?,?,?,?)"
	// 执行sql
	result, err := tx.Exec(sqlStr, po.Id, po.Name, po.Sex, po.Birthday, 0, commfunc.BoolHandle(po.GmtCreated.Valid, po.GmtCreated.Time, nowTime), commfunc.BoolHandle(po.GmtModified.Valid, po.GmtModified.Time, nowTime))
	// 返回结果
	if nil != err {
		return 0, err
	}
	return result.RowsAffected()
}

// 获取用户信息
// 参数：
// po: 获取用户的参数
// 返回：
// *UserPo： 用户信息
// error: 异常信息
func Get(po *FindUserPo) (*UserPo, error) {
	// 1 构建sql
	sql := strings.Builder{}
	sql.WriteString(" select pk_id, id, name, sex, birthday, is_deleted, gmt_created, gmt_modified, gmt_deleted from user ")
	sql.WriteString(" where is_deleted = 0 ")
	if po.Id.Valid {
		sql.WriteString(fmt.Sprintf(" and id = '%s' ", po.Id.String))
	}
	if po.Name.Valid {
		sql.WriteString(fmt.Sprintf(" and name = '%s' ", po.Name.String))
	}
	if po.Sex.Valid {
		sql.WriteString(fmt.Sprintf(" and sex = '%d' ", po.Sex.Int16))
	}
	sql.WriteString(" order by gmt_modified ")
	sql.WriteString(" limit 1 ")
	loginfo.Log.Info("sql:", sql.String())
	// 2 执行
	rows, err := content.Db.Queryx(sql.String())
	// 3 返回结果
	if err != nil {
		return nil, err
	}
	// 是否有数据
	if rows.Next() {
		var userPo UserPo
		e := rows.StructScan(&userPo)
		if nil != e {
			return nil, e
		}
		return &userPo, nil
	}
	return nil, nil
}

// 批量获取用户信息
// 注意：每次最多200个
// 参数：
// po: 查询条件
// 返回：
// []UserPo： 用户信息集合
// error: 异常信息
func List(po *FindUserPo) ([]UserPo, error) {
	// 1 构建sql
	sql := strings.Builder{}
	sql.WriteString(" select pk_id, id, name, sex, birthday, is_deleted, gmt_created, gmt_modified, gmt_deleted from user ")
	sql.WriteString(" where is_deleted = 0 ")
	if po.Id.Valid {
		sql.WriteString(fmt.Sprintf(" and id = '%s' ", po.Id.String))
	}
	if po.Name.Valid {
		sql.WriteString(fmt.Sprintf(" and name = '%s' ", po.Name.String))
	}
	if po.Sex.Valid {
		sql.WriteString(fmt.Sprintf(" and sex = '%d' ", po.Sex.Int16))
	}
	sql.WriteString(" order by gmt_modified ")
	sql.WriteString(" limit 200 ")
	loginfo.Log.Info("sql:", sql.String())
	// 2 执行
	rows, err := content.Db.Queryx(sql.String())
	// 3 返回结果
	if err != nil {
		return nil, err
	}
	userPoList := make([]UserPo, 0)
	// 是否有数据
	for rows.Next() {
		var userPo UserPo
		e := rows.StructScan(&userPo)
		if nil != e {
			return nil, e
		}
		userPoList = append(userPoList, userPo)
	}
	return userPoList, nil
}

// 删除用户
// 注意：每次最多200个
// 参数：
// tx： 事务操作对象
// id: 要删除的用户id
// 返回：
// int64：数据库影响行
// error: 异常信息
func Delete(tx *sqlx.Tx, id string) (int64, error) {
	// 1 构建sql
	sql := "update user set is_deleted = 1, gmt_deleted= now() where is_deleted = 0 and  id = ?"
	// 2 执行
	result, err := tx.Exec(sql, id)
	// 3 返回结果
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// 查询符合某些条件的用户的数量
// 参数：
// po： 查询条件
// 返回：
// int64：数量
// error: 异常信息
func Count(po *FindUserPo) (int64, error) {
	// 1 构建sql
	sql := strings.Builder{}
	sql.WriteString(" select count(*) from user ")
	sql.WriteString(" where is_deleted = 0 ")
	if po.Id.Valid {
		sql.WriteString(fmt.Sprintf(" and id = '%s' ", po.Id.String))
	}
	if po.Name.Valid {
		sql.WriteString(fmt.Sprintf(" and name = '%s' ", po.Name.String))
	}
	if po.Sex.Valid {
		sql.WriteString(fmt.Sprintf(" and sex = '%d' ", po.Sex.Int16))
	}
	loginfo.Log.Info("sql:", sql.String())
	// 2 执行
	rows, err := content.Db.Query(sql.String())
	if err != nil {
		return 0, err
	}
	// 3 返回结果
	var count int64
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
		return count, nil
	}
	return 0, nil
}

// 分页获取用户信息
// 参数：
// po： 查询条件
// pageParam：分页信息
// 返回：
// []UserPo：当前分页下用户信息集合
// error: 异常信息
func PageList(po *FindUserPo, pageParam *model2.PageParam) ([]UserPo, error) {
	// 1 构建sql
	sql := strings.Builder{}
	sql.WriteString(" select pk_id, id, name, sex, birthday, is_deleted, gmt_created, gmt_modified, gmt_deleted from user ")
	sql.WriteString(" where is_deleted = 0 ")
	if po.Id.Valid {
		sql.WriteString(fmt.Sprintf(" and id = '%s' ", po.Id.String))
	}
	if po.Name.Valid {
		sql.WriteString(fmt.Sprintf(" and name = '%s' ", po.Name.String))
	}
	if po.Sex.Valid {
		sql.WriteString(fmt.Sprintf(" and sex = '%d' ", po.Sex.Int16))
	}
	sql.WriteString(" order by gmt_modified ")
	sql.WriteString(fmt.Sprintf(" limit %d, %d ", pageParam.GetCursor(), pageParam.PageSize))
	loginfo.Log.Info("sql:", sql.String())
	// 2 执行
	rows, err := content.Db.Queryx(sql.String())
	// 3 返回结果
	if err != nil {
		return nil, err
	}
	userPoList := make([]UserPo, 0)
	// 是否有数据
	for rows.Next() {
		var userPo UserPo
		e := rows.StructScan(&userPo)
		if nil != e {
			return nil, e
		}
		userPoList = append(userPoList, userPo)
	}
	return userPoList, nil
}

// 更新用户信息
// 参数：
// tx： 事务对象
// updatePo：更新参数信息
// 返回：
// int64：数据库影响行
// error: 异常信息
func UpdateUser(tx *sqlx.Tx, updatePo *UpdatePo) (int64, error) {
	// 1 构建sql
	sql := strings.Builder{}
	sql.WriteString(" update user set ")
	if updatePo.Name.Valid {
		sql.WriteString(fmt.Sprintf(" name = '%s', ", updatePo.Name.String))
	}
	if updatePo.Sex.Valid {
		sql.WriteString(fmt.Sprintf(" sex = '%d', ", updatePo.Sex.Int16))
	}
	if updatePo.Birthday.Valid {
		sql.WriteString(fmt.Sprintf(" birthday = '%s', ", updatePo.Birthday.Time))
	}
	sql.WriteString(" gmt_modified = now() ")
	sql.WriteString(fmt.Sprintf(" where is_deleted = 0 and id = '%s' ", updatePo.Id))
	loginfo.Log.Info("sql:", sql.String())
	// 2 执行
	result, err := tx.Exec(sql.String())
	// 3 返回结果
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
