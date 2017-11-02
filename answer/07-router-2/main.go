package main

import (
	"net/http"
)

func main() {
	r := router{}
	r.Get("/", http.HandlerFunc(index))
	r.Get("/about", http.HandlerFunc(about))

	http.ListenAndServe(":3333", &r)
}

type router struct {
	paths []*path
}

type path struct {
	Method  string
	Path    string
	Handler http.Handler
}

func (router *router) Add(m, p string, h http.Handler) {
	router.paths = append(router.paths, &path{
		Method:  m,
		Path:    p,
		Handler: h,
	})
}

func (router *router) Get(p string, h http.Handler) {
	router.Add(http.MethodGet, p, h)
}

func (router *router) Post(p string, h http.Handler) {
	router.Add(http.MethodPost, p, h)
}

func (router *router) Put(p string, h http.Handler) {
	router.Add(http.MethodPut, p, h)
}

func (router *router) Patch(p string, h http.Handler) {
	router.Add(http.MethodPatch, p, h)
}

func (router *router) Delete(p string, h http.Handler) {
	router.Add(http.MethodDelete, p, h)
}

func (router *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, p := range router.paths {
		if p.Method == r.Method && p.Path == r.URL.Path {
			p.Handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}
