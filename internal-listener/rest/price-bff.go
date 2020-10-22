package rest

import (
	"internal-listener/rest/requestexecutor"

	"github.com/valyala/fasthttp"
)

//PriceBFF ...
type PriceBFF struct {
	dns    *string
	client *fasthttp.Client
}

//NewPriceBFF ...
func NewPriceBFF(dns *string, c *fasthttp.Client) *PriceBFF {
	return &PriceBFF{dns: dns, client: c}
}

//Update ...
func (d *PriceBFF) Update(body []byte, headers map[string]string) {
	requestexecutor.PATCH(d.client, body, headers, *d.dns+"/facilities/products/prices")
}
