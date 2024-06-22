package api

import (
	"github.com/gomxapp/gomx"
	"log"
	"net/http"
	"strconv"

	"gomx-examples.foo_shop/data"
)

func init() {
	gomx.RegisterOnPath("/item/{id}", http.MethodGet,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			arg := r.PathValue("id")
			id, err := strconv.Atoi(arg)
			if err != nil {
				gomx.ReturnBadRequestSimple(w, err)
				return
			}
			log.Println(id)
			item, err := data.GetItem(id)
			log.Println(item)
			err = gomx.ReturnGoHTMLFromFiles(w,
				[]string{"../index.gohtml", "item.gohtml"}, "", item)
			if err != nil {
				gomx.ReturnBadRequestSimple(w, err)
				return
			}
		}),
	)
}
