package resource

import (
	"eventadapter/src/publisher"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//AsnVerifyResource struct
type AsnVerifyResource struct {
	asnVerifyPublisher publisher.Publisher
}

//NewAsnVerifyResource constructor
func NewAsnVerifyResource(p publisher.Publisher, r *fasthttprouter.Router) Resource {
	var rsc = AsnVerifyResource{asnVerifyPublisher: p}
	r.POST("/api/v1/event-adapter/asn-verify/publish", rsc.Publish)
	return &rsc
}

//Publish sends to gcp topic
func (rsc *AsnVerifyResource) Publish(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Enqueque"))
	rsc.asnVerifyPublisher.Publish(ctx.PostBody(), buildEventHeader(&ctx.Request.Header))
}
