package dataBase

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"bytes"

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
