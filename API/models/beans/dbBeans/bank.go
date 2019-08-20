package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `bank` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '银行ID',
//  `bank_name` varchar(255) NOT NULL COMMENT '银行名称',
//  `bank_type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '银行类型',
//  `bank_account` varchar(255) NOT NULL DEFAULT '' COMMENT '银行帐号',
//  `bank_people` varchar(255) NOT NULL COMMENT '联系人',
//  `bank_phone` varchar(20) NOT NULL DEFAULT '' COMMENT '联系电话',
//  `bank_money` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '初期余额',
//  `bank_money_usa` decimal(20,3) NOT NULL DEFAULT '0.000',
//  PRIMARY KEY (`id`),
//  KEY `type` (`bank_type`)
//) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 COMMENT='银行数据'
const TableBank = "bank"
const BankColumns = "`id`,`bank_name`,`bank_type`,`bank_account`,`bank_people`,`bank_phone`,`bank_money`,`bank_money_usa`"

type BankDB struct {
	Id           int64   `description:"银行ID" db:"id" primary:"true"`
	BankName     string  `description:"银行名称" db:"bank_name"`
	BankType     int     `description:"银行类型" db:"bank_type"`
	BankAccount  string  `description:"银行帐号" db:"bank_account"`
	BankPeople   string  `description:"联系人" db:"bank_people"`
	BankPhone    string  `description:"联系电话" db:"bank_phone"`
	BankMoney    float64 `description:"初期余额" db:"bank_money"`
	BankMoneyUsa float64 `description:"" db:"bank_money_usa"`
}
type Bank struct {
	Id           int64   `description:"银行ID" db:"id" primary:"true"`
	BankName     string  `description:"银行名称" db:"bank_name"`
	BankType     int     `description:"银行类型" db:"bank_type"`
	BankAccount  string  `description:"银行帐号" db:"bank_account"`
	BankPeople   string  `description:"联系人" db:"bank_people"`
	BankPhone    string  `description:"联系电话" db:"bank_phone"`
	BankMoney    float64 `description:"初期余额" db:"bank_money"`
	BankMoneyUsa float64 `description:"" db:"bank_money_usa"`
}

//Bank数据转换
func (b *Bank) SetMapValue(m map[string]interface{}) {
	b.Id = db.Int64Default(m["id"])
	b.BankName = db.StringDefault(m["bank_name"])
	b.BankType = db.IntDefault(m["bank_type"])
	b.BankAccount = db.StringDefault(m["bank_account"])
	b.BankPeople = db.StringDefault(m["bank_people"])
	b.BankPhone = db.StringDefault(m["bank_phone"])
	b.BankMoney = db.Float64Default(m["bank_money"])
	b.BankMoneyUsa = db.Float64Default(m["bank_money_usa"])
}

//Bank单个查询
func BankGetOne(where string, q db.Query, args ...interface{}) *Bank {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BankColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBank))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Bank{}
	ret.SetMapValue(m)
	return ret
}

//Bank列表查询
func BankGetList(where string, q db.Query, args ...interface{}) []*Bank {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BankColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBank))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Bank
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Bank{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Bank列表查询
func BankGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Bank {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BankColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBank))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Bank
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Bank{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
