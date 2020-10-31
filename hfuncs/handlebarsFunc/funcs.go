package handlebarsFunc

import (
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris/v12/view"
)

// next number
func Next(eng *view.HandlebarsEngine) {
	eng.AddFunc("Next", func(i string) string {
		var n int
		fmt.Sscanf(i, "%d", &n)
		return fmt.Sprintf("%d", n+1)
	})
}

// next number
func Previous(eng *view.HandlebarsEngine) {
	eng.AddFunc("Previous", func(i string) string {
		var n int
		fmt.Sscanf(i, "%d", &n)
		return fmt.Sprintf("%d", n-1)
	})
}

func PersianDateDay(eng *view.HandlebarsEngine) {
	eng.AddFunc("PersianDateDay", func(d string) string {
		n, err := time.Parse(time.RFC3339, d)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf("%d", n.Day())
	})
}
func PersianDateYear(eng *view.HandlebarsEngine) {
	eng.AddFunc("PersianDateYear", func(d string) string {
		n, err := time.Parse(time.RFC3339, d)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf("%d", n.Year())
	})
}
func PersianDateMon(eng *view.HandlebarsEngine) {
	eng.AddFunc("PersianDateMon", func(d string) string {
		n, err := time.Parse(time.RFC3339, d)
		if err != nil {
			log.Println(err)
		}
		return n.Month().String()
	})
}
