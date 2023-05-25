package middleware

import (
	"net/http"
)

type MethodMap struct {
	Connect func(http.ResponseWriter, *http.Request)
	Delete  func(http.ResponseWriter, *http.Request)
	Get     func(http.ResponseWriter, *http.Request)
	Head    func(http.ResponseWriter, *http.Request)
	Options func(http.ResponseWriter, *http.Request)
	Patch   func(http.ResponseWriter, *http.Request)
	Post    func(http.ResponseWriter, *http.Request)
	Put     func(http.ResponseWriter, *http.Request)
	Trace   func(http.ResponseWriter, *http.Request)
}

func MethodRouter(methodmap MethodMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodConnect:
			methodmap.Connect(w, r)
		case http.MethodDelete:
			methodmap.Delete(w, r)
		case http.MethodGet:
			methodmap.Get(w, r)
		case http.MethodHead:
			methodmap.Head(w, r)
		case http.MethodOptions:
			methodmap.Options(w, r)
		case http.MethodPatch:
			methodmap.Patch(w, r)
		case http.MethodPost:
			methodmap.Post(w, r)
		case http.MethodPut:
			methodmap.Put(w, r)
		case http.MethodTrace:
			methodmap.Trace(w, r)
		}
	}
}
