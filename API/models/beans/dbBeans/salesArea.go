package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `sales_area` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
//  `parent_id` int(10) unsigned NOT NULL COMMENT '父类ID',
//  `name` varchar(255) NOT NULL COMMENT '地区名称',
//  PRIMARY KEY (`id`)
//) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='销售单地区信息'
const TableSalesArea = "sales_area"
const SalesAreaColumns = "`id`,`parent_id`,`name`"

type SalesAreaDB struct {
	Id       int64  `description:"自增ID" db:"id" primary:"true"`
	ParentId int64  `description:"父类ID" db:"parent_id"`
	Name     string `description:"地区名称" db:"name"`
}
type SalesArea struct {
	Id       int64  `description:"自增ID" db:"id" primary:"true"`
	ParentId int64  `description:"父类ID" db:"parent_id"`
	Name     string `description:"地区名称" db:"name"`
}

//SalesArea数据转换
func (s *SalesArea) SetMapValue(m map[string]interface{}) {
	s.Id = db.Int64Default(m["id"])
	s.ParentId = db.Int64Default(m["parent_id"])
	s.Name = db.StringDefault(m["name"])
}

//SalesArea单个查询
func SalesAreaGetOne(where string, q db.Query, args ...interface{}) *SalesArea {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesAreaColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesArea))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &SalesArea{}
	ret.SetMapValue(m)
	return ret
}

//SalesArea列表查询
func SalesAreaGetList(where string, q db.Query, args ...interface{}) []*SalesArea {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesAreaColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesArea))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalesArea
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalesArea{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//SalesArea列表查询
func SalesAreaGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*SalesArea {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesAreaColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesArea))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalesArea
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalesArea{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
