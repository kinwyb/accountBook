package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `operator` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
//  `name` varchar(255) NOT NULL COMMENT '使用者名称',
//  PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='帐号'
const TableOperator = "operator"
const OperatorColumns = "`id`,`name`"

type OperatorDB struct {
	Id   int64  `description:"用户ID" db:"id" primary:"true"`
	Name string `description:"使用者名称" db:"name"`
}
type Operator struct {
	Id   int64  `description:"用户ID" db:"id" primary:"true"`
	Name string `description:"使用者名称" db:"name"`
}

//Operator数据转换
func (o *Operator) SetMapValue(m map[string]interface{}) {
	o.Id = db.Int64Default(m["id"])
	o.Name = db.StringDefault(m["name"])
}

//Operator单个查询
func OperatorGetOne(where string, q db.Query, args ...interface{}) *Operator {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(OperatorColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableOperator))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Operator{}
	ret.SetMapValue(m)
	return ret
}

//Operator列表查询
func OperatorGetList(where string, q db.Query, args ...interface{}) []*Operator {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(OperatorColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableOperator))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Operator
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Operator{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Operator列表查询
func OperatorGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Operator {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(OperatorColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableOperator))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Operator
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Operator{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
