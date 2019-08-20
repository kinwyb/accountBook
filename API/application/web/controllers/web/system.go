package web

import (
	"accountBook/application/web/controllers"
	"accountBook/models/beans/customer"
	"accountBook/models/endpoints/web"
	"encoding/json"

	"github.com/kinwyb/go/db"
)

// 系统相关接口
type SystemController struct {
	controllers.RestController
	Serv web.ISystemEp
}

// @Title 系统日志列表
// @router /logList [post]
func (s *SystemController) LogList() {
	var req customer.LogListReq
	if len(s.Ctx.Input.RequestBody) > 0 {
		if err := json.Unmarshal(s.Ctx.Input.RequestBody, &req); err != nil {
			s.ResponseError(-1, err.Error())
			return
		}
	}
	page, _ := s.GetInt("page", 1)
	pageSize, _ := s.GetInt("pageSize", 20)
	pg := &db.PageObj{Page: page, Rows: pageSize}
	ret, err := s.Serv.LogList(&req, pg, s.OCtx)
	if err != nil {
		s.RespError(err)
		return
	}
	s.Page(pg)
	s.ResponseSUCC(ret)
}
