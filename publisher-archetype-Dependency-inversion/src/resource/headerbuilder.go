package resource

import (
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

var headerKeys = map[string]string{
	"X-Txref":    "eventId",
	"Eventtype":  "eventType",
	"Entityid":   "entityId",
	"Entitytype": "entityType",
	"X-Cmref":    "consumer",
	"Version":    "version",
	"X-Country":  "country",
	"X-Commerce": "commerce",
	"Capability": "capability",
	"Domain":     "domain",
	"Mimetype":   "mimeType",
	"X-Chref":    "channel",
}

func buildEventHeader(reqHeaders *fasthttp.RequestHeader) map[string]string {
	m := make(map[string]string)

	reqHeaders.VisitAll(func(key, value []byte) {
		k, v := headerKeys[string(key)], string(value)
		if k != "" {
			m[k] = v
		}
	})

	m["timestamp"] = strconv.Itoa(int(time.Now().UnixNano()))

	return m
}
