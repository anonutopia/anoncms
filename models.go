package main

import (
	"github.com/jinzhu/gorm"
)

const (
	// PAGE type of page
	PAGE = "page"
	// POST type of page
	POST = "post"
)

// Page model represents any kind of page
type Page struct {
	gorm.Model
	URL         string `sql:"size:255;unique_index"`
	Title       string `sql:"size:255"`
	Description string `sql:"size:255"`
	Type        string `sql:"size:255"`
}
