package main

import (
	"fmt"
	"net/http"

	"github.com/otiai10/marmoset"
)

func main() {
	router := marmoset.NewRouter()
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.FormValue("name"))
	})
	http.ListenAndServe(":8080", router)
}
