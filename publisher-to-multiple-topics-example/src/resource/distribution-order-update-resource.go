package resource

import (
	"eventadapter/src/publisher"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//DistributionOrderUpdateResource struct
type DistributionOrderUpdateResource struct {
	distributionOrderUpdatePublisher publisher.Publisher
}

//NewDistributionOrderUpdateResource constructor
func NewDistributionOrderUpdateResource(p publisher.Publisher, r *fasthttprouter.Router) Resource {
	var rsc = DistributionOrderUpdateResource{distributionOrderUpdatePublisher: p}
	r.POST("/api/v1/event-adapter/distribution-order-update/publish", rsc.Publish)
	return &rsc
}

//Publish sends to gcp topic
func (rsc *DistributionOrderUpdateResource) Publish(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Enqueque"))
	rsc.distributionOrderUpdatePublisher.Publish(ctx.PostBody(), buildEventHeader(&ctx.Request.Header))
}
