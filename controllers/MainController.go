package controllers

import (
	"log"

	"github.com/H4midR/iris-go-simple-mvc/services/response"
	"github.com/kataras/iris/v12"
)

type MainController struct {
}

// some private function
func (c *MainController) handleError(err error, ctx iris.Context) {
	var res response.Response
	//some error loging services may be here
	res.HandleErrCtx(err, ctx)
	ctx.JSON(err) // Or ctx.ViewData("Error",err.Error())
}

// linked to GET : /
func (c *MainController) Get(ctx iris.Context) {
	ctx.ViewData("IsHome", "true")

	ctx.View("header.handlebars")
	ctx.View("index.handlebars")
	ctx.View("footer.handlebars")
}

// linked to POST : /
func (c *MainController) Post(ctx iris.Context) {
	var err error
	var req struct {
		Name                string `form:"name"`
		Email               string `form:"email"`
		Subject             string `form:"subject"`
		Message             string `form:"message"`
		IrisCaptchaResponse string `form:"irisCaptcha"`
	}
	err = ctx.ReadForm(&req)
	if err != nil {
		c.handleError(err, ctx)
		return
	}

	log.Println(req)
	ctx.ViewData("OPStatus", "OK")
	ctx.ViewData("MessageHeader", "Thanks,")
	ctx.ViewData("Message", "We will be in touch with you ♥")

	c.Get(ctx)
}

// linked to Get : /docs
func (c *MainController) GetDocs(ctx iris.Context) {
	ctx.ViewData("IsDocs", "true")
	ctx.View("header.handlebars")
	ctx.View("trimedPages/docs.handlebars")
	ctx.View("footer.handlebars")
}
