package model

import "time"

type Article struct {
	Title string
	Link  string
	Date  time.Time
	Source string
}
