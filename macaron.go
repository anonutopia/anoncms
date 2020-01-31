package main

import (
	"fmt"
	"html/template"
	"strings"
	"time"

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
		Directory: fmt.Sprintf("themes/%s/html", conf.Theme),
	}

	s := macaron.Static(
		fmt.Sprintf("themes/%s/static", conf.Theme),
		macaron.StaticOptions{
			Prefix:      "static",
			SkipLogging: true,
			IndexFile:   "index.html",
			// Expires defines which user-defined function to use for producing a HTTP Expires Header. Default is nil.
			// https://developers.google.com/speed/docs/insights/LeverageBrowserCaching
			Expires: func() string {
				return time.Now().Add(24 * 60 * time.Minute).UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
			},
		})

	m.Use(macaron.Renderer(ro))
	m.Use(s)
	m.Use(cache.Cacher())

	return m
}
