package rest

import (
	"internal-listener/rest/requestexecutor"

	"github.com/valyala/fasthttp"
)

type PriceLogACL struct {
	dns    *string
	client *fasthttp.Client
}

func NewPriceLogACL(dns *string, c *fasthttp.Client) *PriceLogACL {
	return &PriceLogACL{dns: dns, client: c}
}

func (d *PriceLogACL) Save(body []byte, headers map[string]string) {
	requestexecutor.POST(d.client, body, headers, *d.dns+"/api/v1/relational-datasource-adapter/price/save")
}

func (d *PriceLogACL) Update(body []byte, headers map[string]string) {
	requestexecutor.PATCH(d.client, body, headers, *d.dns+"/api/v1/relational-datasource-adapter/gsc/update")
}

func (d *PriceLogACL) Delete(body []byte, headers map[string]string) {
	requestexecutor.DELETE(d.client, body, headers, *d.dns+"/api/v1/relational-datasource-adapter/gsc/delete")
}
