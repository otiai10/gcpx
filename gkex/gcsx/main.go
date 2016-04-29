package main

import (
	"io"
	"net/http"

	"github.com/otiai10/marmoset"
	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
)

const (
	objectName = "foo/bar/baz.txt"                // っていうのがあったんでそれ使う
	projectID  = "otiai10-playground"             // プロジェクトID
	bucketName = "otiai10-playground.appspot.com" // GAEでつかったやつ流用
)

func main() {
	router := marmoset.NewRouter()

	// めんどいので / で全部すませちゃいます
	router.GET("/", readFileController)

	http.ListenAndServe(":8080", router)
}

func readFileController(w http.ResponseWriter, req *http.Request) {

	// contextの取得
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		marmoset.RenderJSON(w, 500, map[string]string{"E001": err.Error()})
		return
	}

	reader, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		marmoset.RenderJSON(w, 500, map[string]string{"E002": err.Error()})
		return
	}

	// めんどくせえそのまま流し込んでしまえ
	io.Copy(w, reader)
}
