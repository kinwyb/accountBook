package customer

import "accountBook/models/beans/dbBeans"

type LogListReq struct {
	StartTime string
	EndTime   string
}

type LogAddReq struct {
	dbBeans.LogDB
	Args []interface{} //参数原型
}
