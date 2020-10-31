package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"github.com/H4midR/iris-go-simple-mvc/controllers"
	"github.com/H4midR/iris-go-simple-mvc/db/myredis"
	"github.com/H4midR/iris-go-simple-mvc/db/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()

	// ? )  request limiter, 30 req/second for each ip
	// lmt := tollbooth.NewLimiter(15, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	// app.Use(tollboothic.LimitHandler(lmt))

	app.UseRouter(logger.New())
	app.UseRouter(recover.New())
	app.Logger().SetLevel("debug")

	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	// uncomment it to active gunzip
	app.Use(iris.Compression) // gzip, deflate, brotli, snappy depending on the client needs.

	//
	// ──────────────────────────────────────────────────────────── REGINSTE VIEW ─────
	//

	//app.I18n.Load("./web/locales/*/*.ini", "fa-IR", "en-US") // use ctx.Tr("key",...)
	// tmp := iris.Handlebars("./web/views", ".html")
	hdbreng := iris.Handlebars("./web/views", ".handlebars")
	hdbreng.Reload(true)

	//tmp.Reload(true)
	app.RegisterView(hdbreng)
	// httemp := iris.HTML("./web/ProgramFiles/blog", ".handlebars")
	// httemp.Reload(true)
	// app.RegisterView(httemp)

	app.HandleDir("/public", iris.Dir("./web/public"))
	// app.HandleDir("/assets", iris.Dir("./web/public"))
	app.HandleDir("/static", iris.Dir("./web/public"))
	// app.HandleDir("/img", iris.Dir("./web/public/img"))
	app.HandleDir("/css", iris.Dir("./web/public/css"))
	app.HandleDir("/js", iris.Dir("./web/public/js"))
	// app.HandleDir("/vendors", iris.Dir("./web/public/vendors"))
	// app.HandleDir("/scss", iris.Dir("./web/public/scss"))
	// app.HandleDir("/fonts", iris.Dir("./web/public/fonts"))

	// crs := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE", "PATCH", "WS"},
	// 	AllowedHeaders:   []string{"Accept", "X-Cat", "X-User", "X-USER", "X-S2SToken", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Authorization-Token", "Screen"},
	// 	AllowCredentials: true,
	// })
	/*
		@ X-User : standard user uid
		@ X-Token : standard user Token
	*/

	app.Get("/favicon.png", func(ctx iris.Context) {
		ctx.SendFile("./web/public/favicon.png", "favicon.png")
		// ctx.SendFileWithRate() // you have the option to rate-limit the upload speed now.
	})

	//
	// ─────────────────────────────────────────────────── PRIVATE KEY GENERATION ─────
	//

	InternalPrivateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Panic(err)
		return
	}

	//
	// ────────────────────────────────────────────────────── DATABASE CONNECTION ─────
	//

	SQLClient, err := mysql.NewClient()
	if err != nil {
		log.Panic(err)
		return
	}

	RedisClient, err := myredis.NewRedisCliet()
	if err != nil {
		log.Panic(err)
		return
	}
	//
	// ────────────────────────────────────────────────── CONTROLLERS DEFINIATION ─────
	//

	// higher priority
	// Handle /book
	bc := controllers.NewBookController(RedisClient, SQLClient, InternalPrivateKey)
	mvc.New(app.Party("/book")).Handle(bc)

	// Handle /
	mvc.New(app.Party("/")).Handle(new(controllers.MainController))
	// ────────────────────────────────────────────────────────────────────────────────

	// app.Get("/file/{action:path}", func(ctx iris.Context) {

	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.View("header.handlebars")
	// 	ctx.View("trimedPages/index.handlebars")
	// 	ctx.View("footer.handlebars")
	// })

	// app.Run(iris.Addr(":9090"), iris.WithoutServerError(iris.ErrServerClosed))
	app.Listen(":9090") // Listen is a shortcut of app.Run(iris.Addr(...)) this is excluded by default now (on normal flow) ^
}
