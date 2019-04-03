package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `sales_products` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
//  `name` varchar(255) NOT NULL COMMENT '产品名称',
//  `money` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '商品单价',
//  `gys` varchar(255) NOT NULL COMMENT '供应商',
//  `cost` decimal(10,3) NOT NULL COMMENT '成本价',
//  PRIMARY KEY (`id`)
//) ENGINE=MyISAM AUTO_INCREMENT=116 DEFAULT CHARSET=utf8
const TableSalesProducts = "sales_products"
const SalesProductsColumns = "`id`,`name`,`money`,`gys`,`cost`"

type SalesProductsDB struct {
	Id    int64   `description:"自增ID" db:"id" primary:"true"`
	Name  string  `description:"产品名称" db:"name"`
	Money float64 `description:"商品单价" db:"money"`
	Gys   string  `description:"供应商" db:"gys"`
	Cost  float64 `description:"成本价" db:"cost"`
}
type SalesProducts struct {
	Id    int64   `description:"自增ID" db:"id" primary:"true"`
	Name  string  `description:"产品名称" db:"name"`
	Money float64 `description:"商品单价" db:"money"`
	Gys   string  `description:"供应商" db:"gys"`
	Cost  float64 `description:"成本价" db:"cost"`
}

//SalesProducts数据转换
func (s *SalesProducts) SetMapValue(m map[string]interface{}) {
	s.Id = db.Int64Default(m["id"])
	s.Name = db.StringDefault(m["name"])
	s.Money = db.Float64Default(m["money"])
	s.Gys = db.StringDefault(m["gys"])
	s.Cost = db.Float64Default(m["cost"])
}

//SalesProducts单个查询
func SalesProductsGetOne(where string, q db.Query, args ...interface{}) *SalesProducts {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesProductsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesProducts))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &SalesProducts{}
	ret.SetMapValue(m)
	return ret
}

//SalesProducts列表查询
func SalesProductsGetList(where string, q db.Query, args ...interface{}) []*SalesProducts {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesProductsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesProducts))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalesProducts
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalesProducts{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//SalesProducts列表查询
func SalesProductsGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*SalesProducts {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesProductsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalesProducts))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalesProducts
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalesProducts{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
