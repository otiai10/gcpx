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
	router.GET("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("おっぱい"))
	})
	// router.GET("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("パイスラッシュ"))
	// })
	http.Handle("/", router)
}
