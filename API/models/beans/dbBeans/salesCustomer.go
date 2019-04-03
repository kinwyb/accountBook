package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `sales_customer` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
//  `name` varchar(255) NOT NULL COMMENT '客户名称',
//  PRIMARY KEY (`id`)
//) ENGINE=MyISAM AUTO_INCREMENT=39 DEFAULT CHARSET=utf8
const TableSalesCustomer = "sales_customer"
const SalesCustomerColumns = "`id`,`name`"

type SalesCustomerDB struct {
	Id   int64  `description:"自增ID" db:"id" primary:"true"`
	Name string `description:"客户名称" db:"name"`
}
type SalesCustomer struct {
	Id   int64  `description:"自增ID" db:"id" primary:"true"`
	Name string `description:"客户名称" db:"name"`
}

//SalesCustomer数据转换
func (s *SalesCustomer) SetMapValue(m map[string]interface{}) {
	s.Id = db.Int64Default(m["id"])
	s.Name = db.StringDefault(m["name"])
}

//SalesCustomer单个查询
func SalesCustomerGetOne(where string, q db.Query, args ...interface{}) *SalesCustomer {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesCustomerColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesCustomer))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &SalesCustomer{}
	ret.SetMapValue(m)
	return ret
}

//SalesCustomer列表查询
func SalesCustomerGetList(where string, q db.Query, args ...interface{}) []*SalesCustomer {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesCustomerColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesCustomer))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalesCustomer
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalesCustomer{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//SalesCustomer列表查询
func SalesCustomerGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*SalesCustomer {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesCustomerColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesCustomer))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalesCustomer
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalesCustomer{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
