package main

import (
	"github.com/jinzhu/gorm"
	macaron "gopkg.in/macaron.v1"
)

var m *macaron.Macaron
var conf *Config
var db *gorm.DB

func main() {
	conf = initConfig()
	db = initDb()
	m = initMacaron()

	m.Get("/", pageView)
	m.Get("/generate", generateView)
	m.Get("/:page", pageView)

	m.Run("0.0.0.0", 5000)
}
