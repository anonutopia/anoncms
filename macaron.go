package main

import (
	"html/template"
	"strings"

	"github.com/go-macaron/cache"
	macaron "gopkg.in/macaron.v1"

	_ "github.com/go-macaron/session/redis"
)

func initMacaron() *macaron.Macaron {
	m := macaron.Classic()

	ro := macaron.RenderOptions{
		Layout: "layout",
		Funcs: []template.FuncMap{map[string]interface{}{
			"obfuscate": func(args ...interface{}) template.HTML {
				email := args[0].(string)
				email = strings.Replace(email, "@", "<span style=\"display:none\">evilspam</span>@", 1)
				return template.HTML(email)

			},
		}},
	}

	m.Use(macaron.Renderer(ro))
	m.Use(cache.Cacher())

	return m
}
