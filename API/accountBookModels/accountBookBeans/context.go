package accountBookBeans

import (
	"code.aliyun.com/zhizaofang/zfgoutil"

	"github.com/kinwyb/go/db"
)

//上下文
type Context struct {
	*zfgoutil.Context
	Token string   //token
	Query db.Query //数据库连接
}

//ChildOf
func (t *Context) Child() *Context {
	zfgoutil.ContextChild(t)
	return t
}

//FollowsFrom
func (t *Context) Follows() *Context {
	zfgoutil.ContextFollows(t)
	return t
}

func (t *Context) QueryTransaction(tx db.Query) *Context {
	ret := &Context{
		Token:   t.Token,
		Query:   t.Query,
		Context: t.Context,
	}
	ret.Query = tx
	return ret
}

//初始化上下文
func NewContext(companyID int, operationName string) *Context {
	ctx := &Context{}
	if operationName != "" && zfgoutil.Tracing {
		ctx.Context = zfgoutil.NewContext(operationName)
	}
	return ctx
}

//初始化上下文
func NewContextWithTracing(tracingSpan zfgoutil.TracingSpan) *Context {
	if !zfgoutil.Tracing {
		tracingSpan = nil
	}
	return &Context{
		Context: zfgoutil.NewContextWithTracing(tracingSpan),
	}
}
