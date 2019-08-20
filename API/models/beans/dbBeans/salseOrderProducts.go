package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `salse_order_products` (
//  `order_id` varchar(255) NOT NULL COMMENT '订单号',
//  `product_id` int(10) unsigned NOT NULL COMMENT '产品ID',
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//  `nums` int(11) NOT NULL COMMENT '购买数量',
//  `price` decimal(10,3) NOT NULL COMMENT '单价',
//  `fregiht` decimal(10,3) NOT NULL COMMENT '快递费用',
//  `other` decimal(10,3) NOT NULL COMMENT '其他费用',
//  `cost` decimal(10,3) NOT NULL COMMENT '成本价',
//  `amount` decimal(10,3) NOT NULL DEFAULT '0.000',
//  `desc` text NOT NULL,
//  PRIMARY KEY (`id`),
//  KEY `order_id` (`order_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='销售订单产品'
const TableSalseOrderProducts = "salse_order_products"
const SalseOrderProductsColumns = "`order_id`,`product_id`,`id`,`nums`,`price`,`fregiht`,`other`,`cost`,`amount`,`desc`"

type SalseOrderProductsDB struct {
	OrderId   string  `description:"订单号" db:"order_id"`
	ProductId int64   `description:"产品ID" db:"product_id"`
	Id        int64   `description:"" db:"id" primary:"true"`
	Nums      int64   `description:"购买数量" db:"nums"`
	Price     float64 `description:"单价" db:"price"`
	Fregiht   float64 `description:"快递费用" db:"fregiht"`
	Other     float64 `description:"其他费用" db:"other"`
	Cost      float64 `description:"成本价" db:"cost"`
	Amount    float64 `description:"" db:"amount"`
	Desc      string  `description:"" db:"desc"`
}
type SalseOrderProducts struct {
	OrderId   string  `description:"订单号" db:"order_id"`
	ProductId int64   `description:"产品ID" db:"product_id"`
	Id        int64   `description:"" db:"id" primary:"true"`
	Nums      int64   `description:"购买数量" db:"nums"`
	Price     float64 `description:"单价" db:"price"`
	Fregiht   float64 `description:"快递费用" db:"fregiht"`
	Other     float64 `description:"其他费用" db:"other"`
	Cost      float64 `description:"成本价" db:"cost"`
	Amount    float64 `description:"" db:"amount"`
	Desc      string  `description:"" db:"desc"`
}

//SalseOrderProducts数据转换
func (s *SalseOrderProducts) SetMapValue(m map[string]interface{}) {
	s.OrderId = db.StringDefault(m["order_id"])
	s.ProductId = db.Int64Default(m["product_id"])
	s.Id = db.Int64Default(m["id"])
	s.Nums = db.Int64Default(m["nums"])
	s.Price = db.Float64Default(m["price"])
	s.Fregiht = db.Float64Default(m["fregiht"])
	s.Other = db.Float64Default(m["other"])
	s.Cost = db.Float64Default(m["cost"])
	s.Amount = db.Float64Default(m["amount"])
	s.Desc = db.StringDefault(m["desc"])
}

//SalseOrderProducts单个查询
func SalseOrderProductsGetOne(where string, q db.Query, args ...interface{}) *SalseOrderProducts {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalseOrderProductsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalseOrderProducts))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &SalseOrderProducts{}
	ret.SetMapValue(m)
	return ret
}

//SalseOrderProducts列表查询
func SalseOrderProductsGetList(where string, q db.Query, args ...interface{}) []*SalseOrderProducts {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalseOrderProductsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalseOrderProducts))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalseOrderProducts
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalseOrderProducts{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//SalseOrderProducts列表查询
func SalseOrderProductsGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*SalseOrderProducts {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalseOrderProductsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSalseOrderProducts))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*SalseOrderProducts
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &SalseOrderProducts{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
