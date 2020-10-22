package resource

import "github.com/valyala/fasthttp"

//Resource interface
type Resource interface {
	Publish(ctx *fasthttp.RequestCtx)
}
