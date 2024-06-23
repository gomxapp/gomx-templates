package main

import (
	"github.com/gomxapp/gomx"
	//_ "{{template "appName"}}/app/api"
	"net/http"
)

func main() {
	r := gomx.DefaultRouter()
	r.AddStaticFiles("static")
	s := gomx.NewServer(&http.Server{
		Addr: "localhost:8080",
	}, r)
	_ = s.ListenAndServe()
}
