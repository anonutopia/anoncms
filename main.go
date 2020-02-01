package main

import (
	macaron "gopkg.in/macaron.v1"
)

var m *macaron.Macaron

var conf *Config

func main() {
	conf = initConfig()

	m = initMacaron()

	m.Get("/", pageView)
	m.Get("/generate", generateView)
	m.Get("/:page", pageView)

	m.Run("0.0.0.0", 5000)
}
