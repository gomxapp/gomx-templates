package main

import (
	"github.com/gomxapp/gomx"
	_ "gomx-examples.project_template_server/app/api"
	"net/http"
)

func main() {
	r := gomx.DefaultRouter()
	r.AddStaticFiles("templates")
	s := gomx.NewServer(&http.Server{
		Addr: ":8081",
	}, r)
	_ = s.ListenAndServe()
}
