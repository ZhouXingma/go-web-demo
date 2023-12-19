package content

import "github.com/jmoiron/sqlx"

var Db *sqlx.DB

// 开启一个事务
// 返回：
// sqlx.Tx : 事务操作对象
// error: 创建事务可能会失败，这是失败后的失败信息
func StartTx() (*sqlx.Tx, error) {
	tx, e := Db.Beginx()
	if nil != e {
		return nil, e
	}
	return tx, nil
}

// 结束事务
// 参数：
// tx：事务对象
// isCommit: 是否进行提交判断的对象
// 返回：
// error： 异常
func EndTx(tx *sqlx.Tx, isCommit bool) error {
	if nil == tx {
		return nil
	}
	var e error = nil
	if isCommit {
		e = tx.Commit()
	} else {
		e = tx.Rollback()
	}
	return e
}

// 结束事务，只不过这个方法会根据异常是否为空来判断进行回滚还是提交
// 参数：
// tx：事务对象
// err: 异常信息
// 返回：
// error： 异常
func EndTxByError(tx *sqlx.Tx, err *error) error {
	return EndTx(tx, *err != nil)
}
