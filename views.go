package main

import (
	macaron "gopkg.in/macaron.v1"
)

func pageView(ctx *macaron.Context) {
	var page string

	if page = ctx.Params(":page"); page == "" {
		page = "home"
		ctx.HTML(
			200,
			page,
			map[string]interface{}{"Page": page},
			macaron.HTMLOptions{Layout: "layout_home"})
		return
	}

	ctx.HTML(200, page, map[string]interface{}{"Page": page})
}

func generateView(ctx *macaron.Context) {
	crawl("http://0.0.0.0:5000")
}
