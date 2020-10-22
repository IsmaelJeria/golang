package resource

import (
	"eventadapter/src/publisher"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//ASNShippingResource struct
type ASNShippingResource struct {
	asnShippingPublisher publisher.Publisher
}

//NewASNShippingResource constructor
func NewASNShippingResource(p publisher.Publisher, r *fasthttprouter.Router) Resource {
	var rsc = ASNShippingResource{asnShippingPublisher: p}
	r.POST("/api/v1/gsc/event-adapter/asn-shipping/publish", rsc.publish)
	return &rsc
}

//Publish sends to gcp topic
func (rsc *ASNShippingResource) publish(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Enqueque"))
	rsc.asnShippingPublisher.Publish(ctx.PostBody(), buildEventHeader(&ctx.Request.Header))
}
