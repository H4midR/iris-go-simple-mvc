package models

import (
	"database/sql"

	"github.com/H4midR/iris-go-simple-mvc/models/book"

	"github.com/go-redis/redis"
)

type Person struct {

	//Uniq Id of User
	UID string `json:"uid"`

	// Personal Information
	// Name : First Name : represent the name or first name of a person
	Name     string `json:"name,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Age      int    `json:"age" validate:"max=100,min=18"` // see go validate https://godoc.org/gopkg.in/go-playground/validator.v9
	//...
}

func (p *Person) Create(sqc *sql.DB, rdc *redis.Client) (err error) {

	//
	//CREATE A PEROSN IN DATA BASE

	if err != nil {
		return
	}

	return
}
func (p *Person) Update(sqc *sql.DB, rdc *redis.Client) (err error) {
	return
}

// also for any other action a user may do
func (p *Person) Own(b *book.Book) (err error) {
	// b.UID()
	// add require record to data base to user own this book
	return
}
