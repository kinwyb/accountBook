package dataBase

import (
	"accountBook/models/log"

	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/db/mysql"
	"github.com/kinwyb/go/err1"
)

var (
	DatabaseNotInitialize = err1.NewError(101, "数据库连接尚未初始化")
	DatabaseConnectFail   = err1.NewError(102, "数据库连接失败")
	SQLError              = err1.NewError(103, "数据库操作异常")
	SQLEmptyChange        = err1.NewError(104, "数据无变化")
	LockFail              = err1.NewError(105, "锁失败")
	UnlockFail            = err1.NewError(106, "解锁失败")
)

//新增数据
func Insert(obj interface{}, table string, q db.Query) err1.Error {
	setSQL, args := mysql.SetSQL(obj)
	if setSQL == "" {
		return SQLError
	}
	sqlString := "INSERT INTO " + q.Table(table) + " SET " + setSQL
	result := q.Exec(sqlString, args...).Error(func(err err1.Error) {
		log.Error(log.DBTag, "新增[%s]错误:%s", table, err.Error())
	})
	return ExecResultHasError(result, true)
}

// 修改数据
func Update(obj interface{}, table string, q db.Query) err1.Error {
	updateSQL, args, err := mysql.Update(q.Table(table), obj)
	if err != nil {
		return err
	}
	result := q.Exec(updateSQL, args...).Error(func(err err1.Error) {
		log.Error(log.DBTag, "更新错误:%s", err.Error())
	})
	return ExecResultHasError(result, false)
}

//执行结果是否有错误
func ExecResultHasError(execresult db.ExecResult, reportZeroChange bool, param ...map[string]string) err1.Error {
	retError := execresult.HasError()
	if retError != nil {
		log.Error(log.DBTag, "[%d]%s", retError.Code(), retError.Msg())
		if retError.Code() == mysql.DuplicateErrorCode { //字段重复
			field := mysql.GetDuplicateField(retError.Msg())
			if len(param) < 1 {
				param = []map[string]string{}
			}
			if len(param) > 0 && param[0] != nil {
				if v, ok := param[0][field]; ok {
					return err1.NewError(retError.Code(), "["+v+"]重复")
				} else if field == mysql.PRIMARY {
					return err1.NewError(retError.Code(), "[主键]重复")
				}
			}
			return err1.NewError(retError.Code(), "唯一数据重复")
		}
		return SQLError
	}
	changrow, _ := execresult.RowsAffected()
	if changrow == 0 && reportZeroChange {
		return SQLEmptyChange
	}
	return nil
}
