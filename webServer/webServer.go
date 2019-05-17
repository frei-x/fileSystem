package webServer

import (
	. "fmt"
	"net/http"
	"reflect"

	"../fileOperation"
)

// main
func Main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(fileOperation.StrName))
		r.ParseForm()
		Println(r.Form["filename"])
		Println(reflect.TypeOf(r.Form["filename"]))
	})
	a := 1 + sum(1, 2.2)
	Println(a)
	http.ListenAndServe(":4000", nil)
}

func sum(n1 int, n2 float32) float32 {
	return float32(n1) + n2
}
