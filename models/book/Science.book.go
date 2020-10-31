package book

import (
	"database/sql"
	"sync"

	"github.com/go-redis/redis"
)

type ScienceBook struct {
	Uid   string `json:"uid"`
	Xid   string `json:"xid"` // external Id , check schema.org
	Title string `json:"title"`
	// Content []Page

	References []string `json:"references"`

	sync.Mutex
}

func (b *ScienceBook) UID() string {
	return b.Uid
}

func (b *ScienceBook) Create(sqc *sql.DB, rdc *redis.Client) (err error) {
	// create a Science book in database
	return
}
func (b *ScienceBook) Ganre() string {
	return "Science"
}
