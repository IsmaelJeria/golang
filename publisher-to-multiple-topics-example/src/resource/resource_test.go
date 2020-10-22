package resource_test

import (
	"eventadapter/src/resource"
	"testing"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type mockPublisher struct{}

func (s *mockPublisher) Publish(payload []byte, headers map[string]string) {
}

func init() {
	publisher := &mockPublisher{}
	router := fasthttprouter.New()
	resource.NewAsnVerifyResource(publisher, router)
	resource.NewDOCancelResource(publisher, router)
	resource.NewDistributionOrderUpdateResource(publisher, router)
	resource.NewInventoryEventResource(publisher, router)
	resource.NewInventorySyncResource(publisher, router)
	resource.NewASNShippingResource(publisher, router)

	go fasthttp.ListenAndServe(":8081", router.Handler)
}

func makeRequest(uri, method string) int {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	req.Header.SetMethod(method)
	req.Header.Add("X-country", "CL")
	req.SetBodyString(`{"msg":"mocked-body"}`)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	return resp.StatusCode()
}

func TestASNVerifyResource(t *testing.T) {
	respCode := makeRequest("http://localhost:8081/api/v1/event-adapter/asn-verify/publish", "POST")
	if respCode != 200 {
		t.Errorf("Status code is not 200: %v", respCode)
	}
}

func TestDOCancellationResource(t *testing.T) {
	respCode := makeRequest("http://localhost:8081/api/v1/gsc/event-adapter/distribution-order-cancellation/publish", "POST")
	if respCode != 200 {
		t.Errorf("Status code is not 200: %v", respCode)
	}
}

func TestDOUpdateResource(t *testing.T) {
	respCode := makeRequest("http://localhost:8081/api/v1/event-adapter/distribution-order-update/publish", "POST")
	if respCode != 200 {
		t.Errorf("Status code is not 200: %v", respCode)
	}
}

func TestInvEventResource(t *testing.T) {
	respCode := makeRequest("http://localhost:8081/api/v1/event-adapter/inventory-event/publish", "POST")
	if respCode != 200 {
		t.Errorf("Status code is not 200: %v", respCode)
	}
}

func TestInvSyncResource(t *testing.T) {
	respCode := makeRequest("http://localhost:8081/api/v1/event-adapter/inventory-sync/publish", "POST")
	if respCode != 200 {
		t.Errorf("Status code is not 200: %v", respCode)
	}
}

func TestASNShippingResource(t *testing.T) {
	respCode := makeRequest("http://localhost:8081/api/v1/gsc/event-adapter/asn-shipping/publish", "POST")
	if respCode != 200 {
		t.Errorf("Status code is not 200: %v", respCode)
	}
}
