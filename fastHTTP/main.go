package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

type myHandler struct {
	foobar string
}

func main() {
	myHandler := &myHandler{
		foobar: "foobar",
	}
	fasthttp.ListenAndServe(":8080", myHandler.handleFastHTTP)
}

func (h *myHandler) handleFastHTTP(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/foo":
		holaMundo(ctx)
	case "/bar":
		chaoMundo(ctx)
		//	case "/baz":
		//		bazHandler.HandlerFunc(ctx)
	default:
		ctx.Error("not found", fasthttp.StatusNotFound)
	}
}

func holaMundo(ctx *fasthttp.RequestCtx) {
	log.Println("hola mundo")
	d := ctx.PostBody()
	log.Printf("%s", d)
	ctx.Write(d)
}

func chaoMundo(ctx *fasthttp.RequestCtx) {
	log.Println("chao mundo")
}
