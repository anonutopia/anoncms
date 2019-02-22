package main

import (
	macaron "gopkg.in/macaron.v1"
)

var m *macaron.Macaron

func main() {
	m = initMacaron()

	m.Get("/", homeView)

	m.Run()
}
