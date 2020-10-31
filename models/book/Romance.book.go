package book

import (
	"database/sql"
	"sync"

	"github.com/go-redis/redis"
)

type RomanceBook struct {
	Uid   string `json:"uid"`
	Xid   string `json:"xid"` // external Id , check schema.org
	Title string `json:"title"`
	// Content []Page

	sync.Mutex
}

func (b *RomanceBook) UID() string {
	return b.Uid
}

func (b *RomanceBook) Create(sqc *sql.DB, rdc *redis.Client) (err error) {
	// create a romance book in database
	return
}
func (b *RomanceBook) Ganre() string {
	return "Ganre"
}
