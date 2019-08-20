package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `receipt_type` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '操作类型ID',
//  `name` varchar(255) NOT NULL COMMENT '名称',
//  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级ID',
//  PRIMARY KEY (`id`,`name`)
//) ENGINE=InnoDB AUTO_INCREMENT=184 DEFAULT CHARSET=utf8 COMMENT='单据类型'
const TableReceiptType = "receipt_type"
const ReceiptTypeColumns = "`id`,`name`,`parent_id`"

type ReceiptTypeDB struct {
	Id       int64  `description:"操作类型ID" db:"id" primary:"true"`
	Name     string `description:"名称" db:"name" primary:"true"`
	ParentId int64  `description:"父级ID" db:"parent_id"`
}
type ReceiptType struct {
	Id       int64  `description:"操作类型ID" db:"id" primary:"true"`
	Name     string `description:"名称" db:"name" primary:"true"`
	ParentId int64  `description:"父级ID" db:"parent_id"`
}

//ReceiptType数据转换
func (r *ReceiptType) SetMapValue(m map[string]interface{}) {
	r.Id = db.Int64Default(m["id"])
	r.Name = db.StringDefault(m["name"])
	r.ParentId = db.Int64Default(m["parent_id"])
}

//ReceiptType单个查询
func ReceiptTypeGetOne(where string, q db.Query, args ...interface{}) *ReceiptType {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ReceiptTypeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableReceiptType))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &ReceiptType{}
	ret.SetMapValue(m)
	return ret
}

//ReceiptType列表查询
func ReceiptTypeGetList(where string, q db.Query, args ...interface{}) []*ReceiptType {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ReceiptTypeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableReceiptType))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*ReceiptType
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &ReceiptType{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//ReceiptType列表查询
func ReceiptTypeGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*ReceiptType {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ReceiptTypeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableReceiptType))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*ReceiptType
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &ReceiptType{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
