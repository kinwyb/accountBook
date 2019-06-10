package controllers

import (
	"accountBook/models/beans"
	"fmt"
	"time"
	"unsafe"

	"github.com/modern-go/reflect2"

	jsoniter "github.com/json-iterator/go"

	"github.com/astaxie/beego"
	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/err1"
	"github.com/rcrowley/go-metrics"
)

var (
	Success         = "操作成功"
	ParamMissing    = err1.NewError(-1, "参数缺失")
	ParamDecodeFail = err1.NewError(-1, "参数解析失败")
)

var json jsoniter.API

func init() {
	json = jsoniter.ConfigFastest
	json.RegisterExtension(&extFloat64{})
}

type RespObj struct {
	Code   int64       `description:"错误编码" json:"code"`
	ErrMsg string      `description:"错误描述" json:"errmsg"`
	Err    string      `description:"错误内容" json:"err"`
	Data   interface{} `description:"返回结果内容" json:"data"`
	Page   *db.PageObj `description:"分页信息" json:"page,omitempty"`
}

//Controller 接口控制器
type RestController struct {
	beego.Controller
	st   time.Time
	resp *RespObj
	OCtx *beans.Context
}

func (ctl *RestController) Prepare() {
	ctl.resp = &RespObj{}
	ctl.OCtx = beans.NewContextWithTracing(
		beans.NewTracingSpanExtractHttpRequest(ctl.Ctx.Request))
	ctl.OCtx.Token = ctl.Ctx.Input.Header("token")
	ctl.st = time.Now()
}

// Render sends the response with rendered template bytes as text/html type.
func (ctl *RestController) Render() error {
	//性能统计处理
	defer func(ctl *RestController) {
		if enableMetrics {
			metrics.GetOrRegisterTimer(ctl.Ctx.Request.URL.Path, metricsRegistry).UpdateSince(ctl.st)
			//总数请求
			metrics.GetOrRegisterTimer(allRequestMetrics, metricsRegistry).UpdateSince(ctl.st)
		}
	}(ctl)
	jsondata, err := json.Marshal(ctl.resp)
	if err != nil {
		return err
	}
	ctl.Ctx.ResponseWriter.Write(jsondata)
	if ctl.OCtx != nil {
		if tracingspan := ctl.OCtx.TracingSpan(); tracingspan != nil {
			tracingspan.InjectHttpHeader(ctl.Ctx.ResponseWriter.Header())
		}
		ctl.OCtx.Finish()
	}
	return nil
}

//ResponseError 返回错误内容
func (ctl *RestController) ResponseError(code int64, msg string, err ...string) *RestController {
	ctl.resp.Code = code
	ctl.resp.ErrMsg = msg
	if err != nil && len(err) > 0 {
		ctl.resp.Err = err[0]
	}
	return ctl
}

//RespError 返回错误内容
func (ctl *RestController) RespError(err err1.Error, e ...error) *RestController {
	ctl.resp.Code = err.Code()
	ctl.resp.ErrMsg = err.Msg()
	ctl.resp.Err = err.Error()
	if len(e) > 0 && e[0] != nil {
		ctl.resp.Err = e[0].Error()
	}
	return ctl
}

//ResponseSUCC 返回成功内容
func (ctl *RestController) ResponseSUCC(data interface{}) *RestController {
	ctl.resp.Code = 0
	ctl.resp.Data = data
	return ctl
}

//ResponseSUCC 返回成功内容
func (ctl *RestController) Response() *RespObj {
	return ctl.resp
}

//Page 分页数据
func (ctl *RestController) Page(page *db.PageObj) {
	ctl.resp.Page = page
}

//错误日志
func (ctl *RestController) LogError(format string, args ...interface{}) {
	log.Error(format, args...)
}

type extFloat64 struct {
	jsoniter.DummyExtension
}

func (e *extFloat64) DecorateEncoder(typ reflect2.Type, encoder jsoniter.ValEncoder) jsoniter.ValEncoder {
	if typ.String() == "float64" {
		return e
	}
	return encoder
}

func (*extFloat64) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	f := *((*float64)(ptr))
	stream.WriteRaw(fmt.Sprintf("%.3f", f))
}

func (*extFloat64) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*float64)(ptr)) == 0
}
