package dataBase

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"bytes"
	"strings"

	"github.com/kinwyb/go/err1"

	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/db/mysql"
)

// 收支列表
func ReceiptList(req *customer.ReceiptListReq, pg *db.PageObj, ctx *beans.Context) []*dbBeans.Receipt {
	defer ctx.Start("db.ReceiptList").Finish()
	sqlString := bytes.NewBufferString("")
	sqlString.WriteString(" 1=1 ")
	var args []interface{}
	sqlString, args = mysql.WhereIN("AND `type`", req.ReceiptType, args, sqlString)
	if req.BankID > 0 {
		sqlString.WriteString(" AND `bank_id` = ? ")
		args = append(args, req.BankID)
	}
	if req.Search != "" {
		sqlString.WriteString(" AND description LIKE ? ")
		args = append(args, "%"+req.Search+"%")
	}
	if req.StartTime != "" {
		sqlString.WriteString(" AND createtime >= ? ")
		args = append(args, req.StartTime)
	}
	if req.EndTime != "" {
		sqlString.WriteString(" AND createtime <= ? ")
		args = append(args, req.EndTime)
	}
	if req.BillType == 1 {
		sqlString.WriteString(" AND money >= 0 ")
	} else if req.BillType == 2 {
		sqlString.WriteString(" AND money < 0")
	}
	sqlString.WriteString(" ORDER BY id DESC ")
	return dbBeans.ReceiptGetPageList(sqlString.String(), ctx.Query, pg, args...)
}

// 时间范围内容金额统计
func ReceiptEndTimeMoneyCount(endTime string, disInOut bool, ctx *beans.Context) []*dbBeans.Receipt {
	defer ctx.Start("db.ReceiptTimeRangeMoneyCount").Finish()
	sqlString := strings.Builder{}
	sqlString.WriteString("SELECT ")
	sqlString.WriteString("`id`,SUM(money) money,`bank_id`,`description`,`createtime`,`lastmodify`,`operator`,`type`,`money_type`")
	sqlString.WriteString(" FROM ")
	sqlString.WriteString(dbBeans.TableReceipt)
	sqlString.WriteString(" WHERE 1=1 ")
	var args []interface{}
	if endTime != "" {
		sqlString.WriteString(" AND createtime <= ? ")
		args = append(args, endTime)
	}
	sqlString.WriteString(" GROUP BY ")
	if disInOut {
		sqlString.WriteString(" sign(money),")
	}
	sqlString.WriteString(" bank_id,money_type ")
	var ret []*dbBeans.Receipt
	ctx.Query.QueryRows(sqlString.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &dbBeans.Receipt{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

// 最后一个ID
func ReceiptLastID(ctx *beans.Context) int64 {
	defer ctx.Start("db.ReceiptLastID").Finish()
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT id FROM ")
	sqlBuilder.WriteString(dbBeans.TableReceipt)
	sqlBuilder.WriteString(" ORDER BY id DESC ")
	return db.Int64Default(ctx.Query.QueryRow(sqlBuilder.String()).Get("id"))
}

// 新增收支单
func ReceiptAdd(req *dbBeans.ReceiptDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("db.ReceiptAdd").Finish()
	return Insert(req, dbBeans.TableReceipt, ctx.Query)
}
