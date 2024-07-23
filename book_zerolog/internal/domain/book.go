package domain

import _ "github.com/benebobaa/valo"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"  valo:"notblank"`
	Year   string `json:"year"   valo:"notblank"`
	Author Author `json:"author" valo:"valid"`
}
