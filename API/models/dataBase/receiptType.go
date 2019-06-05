package dataBase

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"
	"strings"

	"github.com/kinwyb/go/err1"
)

// 收支类型

// 收支类型列表
func ReceiptTypeList(parentID int64, ctx *beans.Context) []*dbBeans.ReceiptType {
	defer ctx.Start("db.ReceiptTypeList").Finish()
	return dbBeans.ReceiptTypeGetList(" parent_id = ? ", ctx.Query, parentID)
}

// 收支类型列表按等级查询
func ReceiptTypeListByLevel(level int64, ctx *beans.Context) []*dbBeans.ReceiptType {
	defer ctx.Start("db.ReceiptTypeListByLevel").Finish()
	whereSQL := strings.Builder{}
	whereSQL.WriteString(" 1=1 ")
	var args []interface{}
	if level < 2 {
		whereSQL.WriteString(" AND parent_id = ? ")
		args = append(args, 0)
	} else {
		whereSQL.WriteString(" AND parent_id != ? ")
		args = append(args, 0)
	}
	whereSQL.WriteString(" GROUP BY name ")
	return dbBeans.ReceiptTypeGetList(whereSQL.String(), ctx.Query, args...)
}

// 查询所有收支类型
func ReceiptTypeListAll(ctx *beans.Context) []*dbBeans.ReceiptType {
	defer ctx.Start("db.ReceiptTypeListAll").Finish()
	return dbBeans.ReceiptTypeGetList("", ctx.Query)
}

// 收支类型列表
func ReceiptTypeQueryByID(id int64, ctx *beans.Context) *dbBeans.ReceiptType {
	defer ctx.Start("db.ReceiptTypeQueryByID").Finish()
	return dbBeans.ReceiptTypeGetOne(" id = ? ", ctx.Query, id)
}

// 收支类型列表
func ReceiptTypeQueryByName(name string, ctx *beans.Context) []*dbBeans.ReceiptType {
	defer ctx.Start("db.ReceiptTypeQueryByName").Finish()
	return dbBeans.ReceiptTypeGetList(" name = ? ", ctx.Query, name)
}

// 收支类型列表
func ReceiptTypeQueryByParentIDAndName(parentID int64, name string, ctx *beans.Context) *dbBeans.ReceiptType {
	defer ctx.Start("db.ReceiptTypeQueryByParentIDAndName").Finish()
	return dbBeans.ReceiptTypeGetOne(" parent_id = ? AND name = ? ", ctx.Query, parentID, name)
}

// 收支类型列表
func ReceiptTypeAdd(req *dbBeans.ReceiptTypeDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("db.ReceiptTypeAdd").Finish()
	return Insert(req, dbBeans.TableReceiptType, ctx.Query)
}
