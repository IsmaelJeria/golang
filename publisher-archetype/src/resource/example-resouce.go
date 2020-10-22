package resource

import (
	"publisher-archetype/src/publisher"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//ExampleResource is a struct
type ExampleResource struct {
	examplePublisher publisher.Publisher
}

//NewExampleResource constructor
func NewExampleResource(p publisher.Publisher, r *fasthttprouter.Router) Resource {
	var rsc = ExampleResource{examplePublisher: p}
	r.POST("/publishroute", rsc.Publish)
	return &rsc
}

//Publish sends to gcp topic
func (rsc *ExampleResource) Publish(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Enqueque"))
	rsc.examplePublisher.Publish(ctx.PostBody(), buildEventHeader(&ctx.Request.Header))
}
