package cloudstoragex

import (
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/cloud/storage"

	"github.com/otiai10/marmoset"
)

const objname = "foo/bar/baz.txt"

func init() {

	marmoset.LoadViews("./views")

	router := marmoset.NewRouter()

	// READするほう
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {

		// Requestから、Contextを取得
		ctx := appengine.NewContext(r)

		// Contextから、default bucket nameを取得
		bucketname, err := file.DefaultBucketName(ctx)
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": "Failed to get bucket name: " + err.Error()})
			return
		}

		// Contextから、Clientの取得
		client, err := storage.NewClient(ctx)
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": "Failed to get client: " + err.Error()})
			return
		}

		// クライアントにバケット名を食わせてBucketHandleを取得し、
		// それにオブジェクト名を食わせてObjectHandleを取得し、
		// そこにContextを食わせてObjectへのReaderを取得
		reader, err := client.Bucket(bucketname).Object(objname).NewReader(ctx)
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": "Failed to get reader: " + err.Error()})
			return
		}
		defer reader.Close()

		// CloudStorage上のObjectの、コンテンツの読み込み
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": err.Error()})
			return
		}

		marmoset.Render(w).HTML("index", map[string]interface{}{
			"objname": objname,
			"content": string(body),
		})
	})

	// WRITEするほう
	router.POST("/", func(w http.ResponseWriter, r *http.Request) {

		// まずはリクエストパラメータのmultipartのパースとファイルオブジェクトの取得
		f, h, err := r.FormFile("foo")
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": "Failed to extract file : " + err.Error()})
			return
		}

		ctx := appengine.NewContext(r)

		// {{{ このへんはGETと同様なのでコメント割愛
		bucketname, err := file.DefaultBucketName(ctx)
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": "Failed to get bucket name: " + err.Error()})
			return
		}

		client, err := storage.NewClient(ctx)
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": "Failed to get client: " + err.Error()})
			return
		}
		// }}}

		// アップロードされたファイルのコンテンツを読む
		body, err := ioutil.ReadAll(f)
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": "Failed to read uploaded file: " + err.Error()})
			return
		}

		// Writer取得
		writer := client.Bucket(bucketname).Object(objname).NewWriter(ctx)
		writer.ContentType = "text/plain"
		defer writer.Close()

		// コンテンツを書き込む
		if _, err := writer.Write(body); err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": "Failed to write to object: " + err.Error()})
			return
		}

		marmoset.Render(w).HTML("index", map[string]interface{}{
			"success": "Successfully uploaded file: " + h.Filename,
			"objname": h.Filename,
			"content": string(body),
		})
	})

	http.Handle("/", router)
}
