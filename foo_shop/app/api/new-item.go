package api

import (
	"github.com/gomxapp/gomx"
	"net/http"
)

func init() {
	gomx.RegisterOnPath("/new-item", http.MethodDelete,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := gomx.ReturnJSON(w, "{}")
			if err != nil {
				gomx.ReturnBadRequestSimple(w, err)
				return
			}
		}),
	)
}
