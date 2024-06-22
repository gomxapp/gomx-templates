package main

import (
	"github.com/gomxapp/gomx"
	_ "gomx-examples.foo_shop/app/api"
	"log"
	"net/http"
)

func main() {
	r := gomx.DefaultRouter()
	r.AddStaticFiles("static")
	s := gomx.NewServer(&http.Server{
		Addr: "localhost:8080",
	}, r)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
