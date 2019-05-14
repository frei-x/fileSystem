package webServer

import (
	. "fmt"
	"net/http"
	"reflect"

	"../fileOperation"
)

func Main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(fileOperation.StrName))
		r.ParseForm()
		Println(r.Form["filename"])
		Println(reflect.TypeOf(r.Form["filename"]))
	})
	http.ListenAndServe(":4000", nil)
}
