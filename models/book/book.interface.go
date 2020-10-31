package book

import (
	"database/sql"

	"github.com/go-redis/redis"
)

type Book interface {
	UID() string
	Ganre() string // or type of Struct Ganre
	//Authers() []Person // the perosons who writes
	Create(sqc *sql.DB, rdc *redis.Client) error
	// Update - Delete - disable ...
	Lock()
	Unlock()
}
