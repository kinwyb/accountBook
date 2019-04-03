package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `sales_user` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
//  `name` varchar(255) NOT NULL COMMENT '业务员',
//  PRIMARY KEY (`id`)
//) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8
const TableSalesUser = "sales_user"
const SalesUserColumns = "`id`,`name`"

type SalesUserDB struct {
	Id   int64  `description:"自增ID" db:"id" primary:"true"`
	Name string `description:"业务员" db:"name"`
}
type SalesUser struct {
	Id   int64  `description:"自增ID" db:"id" primary:"true"`
	Name string `description:"业务员" db:"name"`
}

//SalesUser数据转换
func (s *SalesUser) SetMapValue(m map[string]interface{}) {
	s.Id = db.Int64Default(m["id"])
	s.Name = db.StringDefault(m["name"])
}

//SalesUser单个查询
func SalesUserGetOne(where string, q db.Query, args ...interface{}) *SalesUser {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesUserColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesUser))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &SalesUser{}
	ret.SetMapValue(m)
	return ret
}

//SalesUser列表查询
func SalesUserGetList(where string, q db.Query, args ...interface{}) []*SalesUser {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesUserColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesUser))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalesUser
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalesUser{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//SalesUser列表查询
func SalesUserGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*SalesUser {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesUserColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesUser))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalesUser
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalesUser{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
