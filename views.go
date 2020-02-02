package main

import (
	"log"

	macaron "gopkg.in/macaron.v1"
)

func pageView(ctx *macaron.Context) {
	var page string
	layout := "layout"

	if page = ctx.Params(":page"); page == "" {
		page = "index"
		layout = "layout_home"
	}

	p := &Page{URL: page}
	db.FirstOrCreate(p, p)
	ctx.Data["Page"] = p

	ctx.HTML(
		200,
		page,
		nil,
		macaron.HTMLOptions{Layout: layout})
}

func generateView(ctx *macaron.Context) {
	log.Println("fdsafas")
	crawl("http://0.0.0.0:5000")
}
