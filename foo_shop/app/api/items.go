package api

import (
	"github.com/gomxapp/gomx"
	"log"

	"foo_shop/data"

	"net/http"
	"strconv"
)

func init() {
	gomx.RegisterOnPath("/items", http.MethodGet,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			items := data.GetItems()
			log.Println(items)
			err := gomx.ReturnGoHTMLFromFiles(w,
				[]string{"items.gohtml"}, "items", items)
			if err != nil {
				gomx.ReturnBadRequestSimple(w, err)
				return
			}
		}),
	)

	gomx.RegisterOnPath("/items", http.MethodPost,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()
			if err != nil {
				gomx.ReturnBadRequestSimple(w, err)
				return
			}
			name := r.FormValue("name")
			price, err := strconv.ParseFloat(r.FormValue("price"), 32)
			if err != nil {
				gomx.ReturnBadRequestSimple(w, err)
				return
			}

			err = data.AddItem(name, float32(price))
			if err != nil {
				gomx.ReturnBadRequestSimple(w, err)
				return
			}

			items := data.GetItems()
			item := items[len(items)-1]
			log.Println(item)

			err = gomx.ReturnGoHTMLFromFiles(w,
				[]string{"new-item.gohtml"}, "new-item", item)
			if err != nil {
				gomx.ReturnBadRequestSimple(w, err)
				return
			}
		}),
	)
}
