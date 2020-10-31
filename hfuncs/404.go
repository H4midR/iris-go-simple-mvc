package hfuncs

import (
	"log"

	"github.com/kataras/iris/v12"
)

func P404(ctx iris.Context) {
	log.Printf("ip %s \t notFound \t path %s \n", ctx.GetHeader("X-Forwarded-For"), ctx.Path())
	ctx.View("404.html")
	ctx.StatusCode(iris.StatusNotFound)
}
