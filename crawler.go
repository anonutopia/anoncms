package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly"
)

func crawl(url string) {
	links := make(map[string]bool)
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasPrefix(href, "/") {
			links[href] = true
			c.Visit(fmt.Sprintf("http://0.0.0.0:5000%s", href))
		}
	})

	c.Visit("http://0.0.0.0:5000")

	for link, _ := range links {
		c := colly.NewCollector()
		c.OnResponse(func(r *colly.Response) {
			if link == "/" {
				link = "index"
			}
			f, _ := os.Create(fmt.Sprintf("docs/%s.html", link))
			// fmt.Println(string(r.Body))
			f.Write(r.Body)
			copyDir(fmt.Sprintf("content/%s/static", conf.Theme), "docs/static")
		})
		c.Visit(fmt.Sprintf("http://0.0.0.0:5000%s", link))
	}
}

func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

func copyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = copyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}
