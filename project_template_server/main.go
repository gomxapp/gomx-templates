package main

import (
	"github.com/gomxapp/gomx"
	"net/http"
	_ "project_template_server/app/api"
)

func main() {
	r := gomx.DefaultRouter()
	r.AddStaticFiles("templates")
	s := gomx.NewServer(&http.Server{
		Addr: ":8081",
	}, r)
	_ = s.ListenAndServe()
}
