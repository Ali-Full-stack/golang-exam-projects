package models

import (
	"fmt"
	"time"
)

type Author struct {
	Author_id  int    `json:"author_id"`
	Name       string `json:"name"`
	Birth_date string `json:"birth_date"`
	Biography  string `json:"biography"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Books struct {
	Book_id          int    `json:"book_id"`
	Title            string `json:"title"`
	Category         string `json:"category"`
	Author_id        int    `json:"author_id"`
	Publication_date string `json:"publication_date"`
	Isbn             string `json:"isbn"`
	Description      string `json:"description"`
	Created_at       string `json:"created_at"`
	Updated_at       string `json:"updated_at"`
}

func CurrentStringtime() string {
	time := time.Now()
	return fmt.Sprintf("%s-%s-%s", time.Year(), time, time.Month(), time.Day())
}
