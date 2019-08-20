package beans

import (
	"github.com/kinwyb/go/db/mysql"
	opentracinglog "github.com/opentracing/opentracing-go/log"

	"github.com/kinwyb/go/db"
)

type IContext interface {
	TracingContext

	Copy() IContext

	SetTracingType(int)

	Logf(fields ...opentracinglog.Field) IContext

	TracingSpan() TracingSpan
}

//上下文
type Context struct {
	tracing     TracingSpan   //追踪数据
	tracingType int           //追踪类型[1=ChildOf,2=FollowsFrom]
	Token       string        //token
	Query       db.Query      //数据库连接
	RequestArgs []interface{} //请求参数
}

//ChildOf
func ContextChild(parentCtx IContext) IContext {
	if parentCtx == nil {
		return NewContext("ContextChild")
	}
	if Tracing && parentCtx.TracingSpan() != nil {
		ret := parentCtx.Copy()
		ret.SetTracingType(TracingChild)
		return ret
	}
	return parentCtx
}

//FollowsFrom
func ContextFollows(parentCtx IContext) IContext {
	if parentCtx == nil {
		return NewContext("ContextFollows")
	}
	if Tracing && parentCtx.TracingSpan() != nil {
		ret := parentCtx.Copy()
		ret.SetTracingType(TracingFollowsFrom)
		return ret
	}
	return parentCtx
}

//Copy
func (t *Context) Copy() IContext {
	return &Context{
		tracing:     t.tracing,
		tracingType: t.tracingType,
	}
}

func (t *Context) SetTracingType(tracingType int) {
	t.tracingType = tracingType
}

//Finish
func (t *Context) Finish() {
	if t.tracing != nil {
		t.tracing.Finish()
	}
}

//日志
func (t *Context) Logf(fields ...opentracinglog.Field) IContext {
	if t.tracing != nil && Tracing {
		t.tracing.Logf(fields...)
	}
	return t
}

//追踪信息获取,可能返回nil
func (t *Context) TracingSpan() TracingSpan {
	return t.tracing
}

//Start
func (t *Context) Start(operationName string) TracingContext {
	//t.tracingStart = true
	if Tracing {
		if t.tracing == nil || t.tracing.Span() == nil { //没有父级span,生成根span
			t.tracing = NewTracingSpanStart(operationName)
		} else { //有父级span的按类型延伸子级,如果类型为空的不处理
			switch t.tracingType {
			case TracingChild:
				t.tracing = t.tracing.ChildOf(operationName)
			case TracingFollowsFrom:
				t.tracing = t.tracing.FollowsFrom(operationName)
			}
		}
	}
	t.tracingType = 0 //清空追踪类型,往后传递没有指定类型时按之前值往下扩展
	return t
}

//初始化上下文
func NewContext(operationName string) *Context {
	ctx := &Context{
		Query: mysql.GetDBConnect(),
	}
	if operationName != "" && Tracing {
		ctx.tracing = NewTracingSpanStart(operationName)
	}
	return ctx
}

//初始化上下文
func NewContextWithTracing(tracingSpan TracingSpan) *Context {
	if !Tracing {
		tracingSpan = nil
	}
	return &Context{
		tracing: tracingSpan,
		Query:   mysql.GetDBConnect(),
	}
}

//ChildOf
func (t *Context) Child() *Context {
	ContextChild(t)
	return t
}

//FollowsFrom
func (t *Context) Follows() *Context {
	ContextFollows(t)
	return t
}

func (t *Context) QueryTransaction(tx db.Query) *Context {
	ret := &Context{
		Token:       t.Token,
		Query:       t.Query,
		tracing:     t.tracing,
		tracingType: t.tracingType,
	}
	ret.Query = tx
	return ret
}
