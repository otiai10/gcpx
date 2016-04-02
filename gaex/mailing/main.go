package oppai

import (
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"

	"github.com/otiai10/marmoset"
)

func init() {

	marmoset.LoadViews("./views")

	router := marmoset.NewRouter()

	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		marmoset.Render(w).HTML("index", nil)
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request) {

		mailto := r.FormValue("mailto")
		if mailto == "" {
			marmoset.Render(w).HTML("index", map[string]interface{}{
				"error": "Address Undefined",
			})
			return
		}

		ctx := appengine.NewContext(r)
		msg := &mail.Message{
			Sender:  os.Getenv("GAEX_MAIL_SENDER"),
			To:      []string{mailto},
			Subject: "GAE/Go Mailing Example",
			Body:    "This is plain text body",
			HTMLBody: `<h1>This is HTML body</h1><p>
			<a href='https://github.com/otiai10/gcpx/tree/master/gaex/mailing'>source code</a>`,
		}

		if err := mail.Send(ctx, msg); err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		marmoset.Render(w).HTML("index", map[string]interface{}{
			"success": mailto,
		})
	})

	http.Handle("/", router)
}
