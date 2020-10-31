package controllers

import (
	"crypto/rsa"
	"database/sql"

	"github.com/H4midR/iris-go-simple-mvc/models/book"
	"github.com/H4midR/iris-go-simple-mvc/services/response"
	"github.com/go-redis/redis"
	"github.com/kataras/iris/v12"
)

type BookController struct {
	rdc *redis.Client
	sqc *sql.DB
	pk  *rsa.PrivateKey // private key for enceription
}

func NewBookController(rdc *redis.Client, sqc *sql.DB, pk *rsa.PrivateKey) *BookController {
	bc := new(BookController)
	bc.rdc = rdc
	bc.sqc = sqc
	bc.pk = pk

	return bc
}

// Handle Get /book  			check main.go
func (c *BookController) Get(ctx iris.Context) response.Response {
	var res response.Response
	var err error
	// authenticate clinet using barier token or custom headers or use middlewares. and check acl

	// then

	// list of book
	bs := []book.Book{}
	// read books from database
	if res.HandleErrCtx(err, ctx) {
		return res
	}

	res.OK()
	res.Data = bs
	return res
}

func (c *BookController) Post(ctx iris.Context) response.Response {
	var res response.Response
	var err error
	// authenticate clinet using barier token or custom headers or use middlewares. and check acl

	// then

	// read data in json
	var req struct {
		Name string `json:"name"`
		//...
	}

	err = ctx.ReadJSON(&req)
	if res.HandleErrCtx(err, ctx) {
		return res
	}

	// create book and generate uid and xid and date created
	var b book.RomanceBook
	b.Title = req.Name
	err = b.Create(c.sqc, c.rdc)
	if res.HandleErrCtx(err, ctx) {
		return res
	}

	res.OK()
	res.Data = b
	return res
}

// handle Get /book/:uid ex /book/0x12c
func (c *BookController) GetBy(bookUid string, ctx iris.Context) response.Response {
	var res response.Response
	var err error
	// authenticate clinet using barier token or custom headers or use middlewares. and check acl

	// then

	// load the book form data base or cash
	b := book.ScienceBook{}
	// read books from database
	if res.HandleErrCtx(err, ctx) {
		return res
	}

	res.OK()
	res.Data = b
	return res
}

// or delete a book with DeleteBy
// func (c *BookController) DeleletBy(bookUid string, ctx iris.Context) response.Response
