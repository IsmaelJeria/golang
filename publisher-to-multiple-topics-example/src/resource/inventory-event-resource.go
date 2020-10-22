package resource

import (
	"eventadapter/src/publisher"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//InventoryEventResource struct
type InventoryEventResource struct {
	inventoryEventPublisher publisher.Publisher
}

//NewInventoryEventResource constructor
func NewInventoryEventResource(p publisher.Publisher, r *fasthttprouter.Router) Resource {
	var rsc = InventoryEventResource{inventoryEventPublisher: p}
	r.POST("/api/v1/event-adapter/inventory-event/publish", rsc.Publish)
	return &rsc
}

//Publish sends to gcp topic
func (rsc *InventoryEventResource) Publish(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Enqueque"))
	rsc.inventoryEventPublisher.Publish(ctx.PostBody(), buildEventHeader(&ctx.Request.Header))
}
