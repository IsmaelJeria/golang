package requestexecutor

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

type errorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

//POST request
func POST(c *fasthttp.Client, body []byte, headers map[string]string, uri string) {

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.Header.SetMethodBytes([]byte("POST"))

	for key, element := range headers {
		req.Header.SetBytesKV([]byte(key), []byte(element))
	}
	//	request.Header.VisitAll(func(key, value []byte) {
	//		req.Header.SetBytesKV(key, value)
	//	})
	req.URI().Update(uri)
	req.Header.SetContentType("application/json")

	req.SetBody(body)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	err := c.Do(req, res)

	if err != nil {
		log.Println(err)
	}

	if res.StatusCode() > 299 {
		var msg errorMessage
		msg.Code = res.StatusCode()
		msg.Path = req.URI().String()
		msg.Message = string(res.Body())
		if msg.Message == "" {
			msg.Message = "Error"
		}
	}
	fmt.Println(string(body))
	fmt.Println(string(res.Body()))

}

func PATCH(c *fasthttp.Client, body []byte, headers map[string]string, uri string) {

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.Header.SetMethodBytes([]byte("PATCH"))

	for key, element := range headers {
		req.Header.SetBytesKV([]byte(key), []byte(element))
	}
	//	request.Header.VisitAll(func(key, value []byte) {
	//		req.Header.SetBytesKV(key, value)
	//	})
	req.URI().Update(uri)
	req.Header.SetContentType("application/json")
	req.SetBody(body)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	err := c.Do(req, res)

	if err != nil {
		log.Println(err)
		log.Println("PATCH  ERROR  " + uri)
	}

	if res.StatusCode() > 299 {
		var msg errorMessage
		msg.Code = res.StatusCode()
		msg.Path = req.URI().String()
		msg.Message = string(res.Body())
		if msg.Message == "" {
			msg.Message = "Error"
		}
	}

}

func DELETE(c *fasthttp.Client, body []byte, headers map[string]string, uri string) {

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.Header.SetMethodBytes([]byte("DELETE"))

	for key, element := range headers {
		req.Header.SetBytesKV([]byte(key), []byte(element))
	}
	//	request.Header.VisitAll(func(key, value []byte) {
	//		req.Header.SetBytesKV(key, value)
	//	})
	req.URI().Update(uri)
	req.Header.SetContentType("application/json")
	req.SetBody(body)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	err := c.Do(req, res)

	if err != nil {
		log.Println(err)
		log.Println("PATCH  ERROR  " + uri)
	}

	if res.StatusCode() > 299 {
		var msg errorMessage
		msg.Code = res.StatusCode()
		msg.Path = req.URI().String()
		msg.Message = string(res.Body())
		if msg.Message == "" {
			msg.Message = "Error"
		}
	}

}
