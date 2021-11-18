// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"net/http"
	"os"
	"time"
)

type Page struct {
	Hostname string
	Year string
}

func loadPage() (*Page, error) {
	t := time.Now()
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return &Page{Hostname: hostname, Year: t.Format("2006") }, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage()
	t, _ := template.ParseFiles("index.tpl")
	t.Execute(w, p)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/public/", resourceHandler)
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

