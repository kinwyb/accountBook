package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `bank_change` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
//  `out_bank_id` int(10) unsigned NOT NULL COMMENT '转出银行ID',
//  `in_bank_id` int(10) unsigned NOT NULL COMMENT '转入银行ID',
//  `in_receipt` int(10) unsigned NOT NULL COMMENT '转入单据',
//  `out_receipt` int(10) unsigned NOT NULL COMMENT '转出单据',
//  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '单据创建时间',
//  `opter` int(10) unsigned NOT NULL COMMENT '操作者ID',
//  `money` decimal(20,3) NOT NULL COMMENT '金额',
//  `desc` text COMMENT '备注',
//  `money_type` tinyint(3) unsigned NOT NULL DEFAULT '2',
//  PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8 COMMENT='银行转账ID'
const TableBankChange = "bank_change"
const BankChangeColumns = "`id`,`out_bank_id`,`in_bank_id`,`in_receipt`,`out_receipt`,`createtime`,`opter`,`money`,`desc`,`money_type`"

type BankChangeDB struct {
	Id         int64   `description:"自增ID" db:"id" primary:"true"`
	OutBankId  int64   `description:"转出银行ID" db:"out_bank_id"`
	InBankId   int64   `description:"转入银行ID" db:"in_bank_id"`
	InReceipt  int64   `description:"转入单据" db:"in_receipt"`
	OutReceipt int64   `description:"转出单据" db:"out_receipt"`
	Createtime string  `description:"单据创建时间" db:"createtime"`
	Opter      int64   `description:"操作者ID" db:"opter"`
	Money      float64 `description:"金额" db:"money"`
	Desc       *string `description:"备注" db:"desc"`
	MoneyType  int     `description:"" db:"money_type"`
}
type BankChange struct {
	Id         int64   `description:"自增ID" db:"id" primary:"true"`
	OutBankId  int64   `description:"转出银行ID" db:"out_bank_id"`
	InBankId   int64   `description:"转入银行ID" db:"in_bank_id"`
	InReceipt  int64   `description:"转入单据" db:"in_receipt"`
	OutReceipt int64   `description:"转出单据" db:"out_receipt"`
	Createtime string  `description:"单据创建时间" db:"createtime"`
	Opter      int64   `description:"操作者ID" db:"opter"`
	Money      float64 `description:"金额" db:"money"`
	Desc       string  `description:"备注" db:"desc"`
	MoneyType  int     `description:"" db:"money_type"`
}

//BankChange数据转换
func (b *BankChange) SetMapValue(m map[string]interface{}) {
	b.Id = db.Int64Default(m["id"])
	b.OutBankId = db.Int64Default(m["out_bank_id"])
	b.InBankId = db.Int64Default(m["in_bank_id"])
	b.InReceipt = db.Int64Default(m["in_receipt"])
	b.OutReceipt = db.Int64Default(m["out_receipt"])
	b.Createtime = db.StringDefault(m["createtime"])
	b.Opter = db.Int64Default(m["opter"])
	b.Money = db.Float64Default(m["money"])
	b.Desc = db.StringDefault(m["desc"])
	b.MoneyType = db.IntDefault(m["money_type"])
}

//BankChange单个查询
func BankChangeGetOne(where string, q db.Query, args ...interface{}) *BankChange {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BankChangeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBankChange))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &BankChange{}
	ret.SetMapValue(m)
	return ret
}

//BankChange列表查询
func BankChangeGetList(where string, q db.Query, args ...interface{}) []*BankChange {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BankChangeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBankChange))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BankChange
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &BankChange{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//BankChange列表查询
func BankChangeGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*BankChange {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BankChangeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBankChange))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BankChange
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &BankChange{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
