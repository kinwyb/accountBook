package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `sales` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
//  `area` int(10) unsigned NOT NULL COMMENT '地区ID',
//  `order_id` varchar(255) NOT NULL COMMENT '订单编号',
//  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '订货日期',
//  `customer` int(10) unsigned NOT NULL COMMENT '客户ID',
//  `money_usa` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '美金金额',
//  `money` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '人民币金额',
//  `sales` int(10) unsigned NOT NULL COMMENT '业务员ID',
//  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '单据状况(0=未审核,1=审核成功,2=审核失败)',
//  `finish_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '完成情况(0=未完成,1=完成)',
//  `description` text NOT NULL COMMENT '描述',
//  `profit` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '利润',
//  `gd` varchar(255) NOT NULL COMMENT '跟单',
//  `mlv` double(255,0) NOT NULL DEFAULT '0' COMMENT '毛利率',
//  `hv` double NOT NULL COMMENT '汇率',
//  `inputTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//  `passTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '审核时间',
//  PRIMARY KEY (`id`,`order_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8 COMMENT='销售单'
const TableSales = "sales"
const SalesColumns = "`id`,`area`,`order_id`,`createtime`,`customer`,`money_usa`,`money`,`sales`,`status`,`finish_status`,`description`,`profit`,`gd`,`mlv`,`hv`,`inputTime`,`passTime`"

type SalesDB struct {
	Id           int64   `description:"序号" db:"id" primary:"true"`
	Area         int64   `description:"地区ID" db:"area"`
	OrderId      string  `description:"订单编号" db:"order_id" primary:"true"`
	Createtime   string  `description:"订货日期" db:"createtime"`
	Customer     int64   `description:"客户ID" db:"customer"`
	MoneyUsa     float64 `description:"美金金额" db:"money_usa"`
	Money        float64 `description:"人民币金额" db:"money"`
	Sales        int64   `description:"业务员ID" db:"sales"`
	Status       int     `description:"单据状况(0=未审核,1=审核成功,2=审核失败)" db:"status"`
	FinishStatus int     `description:"完成情况(0=未完成,1=完成)" db:"finish_status"`
	Description  string  `description:"描述" db:"description"`
	Profit       float64 `description:"利润" db:"profit"`
	Gd           string  `description:"跟单" db:"gd"`
	Mlv          float64 `description:"毛利率" db:"mlv"`
	Hv           float64 `description:"汇率" db:"hv"`
	InputTime    string  `description:"创建时间" db:"inputTime"`
	PassTime     string  `description:"审核时间" db:"passTime"`
}
type Sales struct {
	Id           int64   `description:"序号" db:"id" primary:"true"`
	Area         int64   `description:"地区ID" db:"area"`
	OrderId      string  `description:"订单编号" db:"order_id" primary:"true"`
	Createtime   string  `description:"订货日期" db:"createtime"`
	Customer     int64   `description:"客户ID" db:"customer"`
	MoneyUsa     float64 `description:"美金金额" db:"money_usa"`
	Money        float64 `description:"人民币金额" db:"money"`
	Sales        int64   `description:"业务员ID" db:"sales"`
	Status       int     `description:"单据状况(0=未审核,1=审核成功,2=审核失败)" db:"status"`
	FinishStatus int     `description:"完成情况(0=未完成,1=完成)" db:"finish_status"`
	Description  string  `description:"描述" db:"description"`
	Profit       float64 `description:"利润" db:"profit"`
	Gd           string  `description:"跟单" db:"gd"`
	Mlv          float64 `description:"毛利率" db:"mlv"`
	Hv           float64 `description:"汇率" db:"hv"`
	InputTime    string  `description:"创建时间" db:"inputTime"`
	PassTime     string  `description:"审核时间" db:"passTime"`
}

//Sales数据转换
func (s *Sales) SetMapValue(m map[string]interface{}) {
	s.Id = db.Int64Default(m["id"])
	s.Area = db.Int64Default(m["area"])
	s.OrderId = db.StringDefault(m["order_id"])
	s.Createtime = db.StringDefault(m["createtime"])
	s.Customer = db.Int64Default(m["customer"])
	s.MoneyUsa = db.Float64Default(m["money_usa"])
	s.Money = db.Float64Default(m["money"])
	s.Sales = db.Int64Default(m["sales"])
	s.Status = db.IntDefault(m["status"])
	s.FinishStatus = db.IntDefault(m["finish_status"])
	s.Description = db.StringDefault(m["description"])
	s.Profit = db.Float64Default(m["profit"])
	s.Gd = db.StringDefault(m["gd"])
	s.Mlv = db.Float64Default(m["mlv"])
	s.Hv = db.Float64Default(m["hv"])
	s.InputTime = db.StringDefault(m["inputTime"])
	s.PassTime = db.StringDefault(m["passTime"])
}

//Sales单个查询
func SalesGetOne(where string, q db.Query, args ...interface{}) *Sales {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSales))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Sales{}
	ret.SetMapValue(m)
	return ret
}

//Sales列表查询
func SalesGetList(where string, q db.Query, args ...interface{}) []*Sales {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSales))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Sales
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Sales{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Sales列表查询
func SalesGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Sales {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SalesColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSales))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Sales
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Sales{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
