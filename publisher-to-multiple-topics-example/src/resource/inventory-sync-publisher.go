package resource

import (
	"eventadapter/src/publisher"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//InventorySyncResource struct
type InventorySyncResource struct {
	inventorySyncPublisher publisher.Publisher
}

//NewInventorySyncResource constructor
func NewInventorySyncResource(p publisher.Publisher, r *fasthttprouter.Router) Resource {
	var rsc = InventorySyncResource{inventorySyncPublisher: p}
	r.POST("/api/v1/event-adapter/inventory-sync/publish", rsc.Publish)
	return &rsc
}

//Publish sends to gcp topic
func (rsc *InventorySyncResource) Publish(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Enqueque"))
	rsc.inventorySyncPublisher.Publish(ctx.PostBody(), buildEventHeader(&ctx.Request.Header))
}
