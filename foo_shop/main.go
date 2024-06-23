package main

import (
	_ "foo_shop/app/api"
	"github.com/gomxapp/gomx"
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
