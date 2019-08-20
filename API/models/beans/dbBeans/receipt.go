package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `receipt` (
//  `id` bigint(16) unsigned NOT NULL AUTO_INCREMENT COMMENT '单据号',
//  `money` decimal(20,3) NOT NULL COMMENT '交易金额',
//  `bank_id` int(10) unsigned NOT NULL COMMENT '银行ID',
//  `description` text NOT NULL COMMENT '备注',
//  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
//  `lastmodify` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后修改时间',
//  `operator` int(10) unsigned NOT NULL COMMENT '操作者',
//  `type` int(10) unsigned NOT NULL COMMENT '单据类型(receipt_type表)',
//  `money_type` tinyint(3) unsigned NOT NULL DEFAULT '2',
//  PRIMARY KEY (`id`),
//  KEY `type_bank_time` (`bank_id`,`createtime`,`type`) USING BTREE
//) ENGINE=InnoDB AUTO_INCREMENT=1526 DEFAULT CHARSET=utf8 COMMENT='收入支出单据'
const TableReceipt = "receipt"
const ReceiptColumns = "`id`,`money`,`bank_id`,`description`,`createtime`,`lastmodify`,`operator`,`type`,`money_type`"

type ReceiptDB struct {
	Id          int64   `description:"单据号" db:"id" primary:"true"`
	Money       float64 `description:"交易金额" db:"money"`
	BankId      int64   `description:"银行ID" db:"bank_id"`
	Description string  `description:"备注" db:"description"`
	Createtime  string  `description:"生成时间" db:"createtime"`
	Lastmodify  string  `description:"最后修改时间" db:"lastmodify"`
	Operator    int64   `description:"操作者" db:"operator"`
	Type        int64   `description:"单据类型(receipt_type表)" db:"type"`
	MoneyType   int     `description:"" db:"money_type"`
}
type Receipt struct {
	Id          int64   `description:"单据号" db:"id" primary:"true"`
	Money       float64 `description:"交易金额" db:"money"`
	BankId      int64   `description:"银行ID" db:"bank_id"`
	Description string  `description:"备注" db:"description"`
	Createtime  string  `description:"生成时间" db:"createtime"`
	Lastmodify  string  `description:"最后修改时间" db:"lastmodify"`
	Operator    int64   `description:"操作者" db:"operator"`
	Type        int64   `description:"单据类型(receipt_type表)" db:"type"`
	MoneyType   int     `description:"" db:"money_type"`
}

//Receipt数据转换
func (r *Receipt) SetMapValue(m map[string]interface{}) {
	r.Id = db.Int64Default(m["id"])
	r.Money = db.Float64Default(m["money"])
	r.BankId = db.Int64Default(m["bank_id"])
	r.Description = db.StringDefault(m["description"])
	r.Createtime = db.StringDefault(m["createtime"])
	r.Lastmodify = db.StringDefault(m["lastmodify"])
	r.Operator = db.Int64Default(m["operator"])
	r.Type = db.Int64Default(m["type"])
	r.MoneyType = db.IntDefault(m["money_type"])
}

//Receipt单个查询
func ReceiptGetOne(where string, q db.Query, args ...interface{}) *Receipt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ReceiptColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableReceipt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Receipt{}
	ret.SetMapValue(m)
	return ret
}

//Receipt列表查询
func ReceiptGetList(where string, q db.Query, args ...interface{}) []*Receipt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ReceiptColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableReceipt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Receipt
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Receipt{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Receipt列表查询
func ReceiptGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Receipt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ReceiptColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableReceipt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Receipt
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Receipt{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
