package resource

import (
	"eventadapter/src/publisher"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//DOCancelResource struct
type DOCancelResource struct {
	doCancelPublisher publisher.Publisher
}

//NewDOCancelResource constructor
func NewDOCancelResource(p publisher.Publisher, r *fasthttprouter.Router) Resource {
	var rsc = DOCancelResource{doCancelPublisher: p}
	r.POST("/api/v1/gsc/event-adapter/distribution-order-cancellation/publish", rsc.Publish)
	return &rsc
}

//Publish sends to gcp topic
func (rsc *DOCancelResource) Publish(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Enqueque"))
	rsc.doCancelPublisher.Publish(ctx.PostBody(), buildEventHeader(&ctx.Request.Header))
}
