package oppai

import (
	"net/http"

	"github.com/otiai10/marmoset"
)

func init() {

	marmoset.LoadViews("./views")

	router := marmoset.NewRouter()
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		marmoset.Render(w).HTML("index", nil)
	})
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		marmoset.Render(w).HTML("hello", map[string]interface{}{
			"name": r.FormValue("name"),
		})
	})
	http.Handle("/", router)
}
