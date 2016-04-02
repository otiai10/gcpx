package cloudsqlx

import (
	"net/http"
	"os"
	"strconv"

	"google.golang.org/appengine"

	_ "github.com/go-sql-driver/mysql"

	// _ "google.golang.org/appengine/cloudsql"

	"github.com/jinzhu/gorm"
	"github.com/otiai10/marmoset"
)

// User represents model in users table
type User struct {
	ID   int
	Name string `json:"name" gorm:"NOT NULL;UNIQUE"`
	Age  int    `json:"age"  gorm:"TYPE:int;NOT NULL"`
}

func init() {

	uri := os.Getenv("DATABASE_URI") // 最初はこれだけでいいと思ったんだけど
	// Devサーバの中まで環境変数を渡せないので、なんらかの直接指定が必要
	if appengine.IsDevAppServer() {
		uri = "root@tcp(localhost:3306)/hoge"
	}

	db, err := gorm.Open("mysql", uri)
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}

	marmoset.LoadViews("./views")
	router := marmoset.NewRouter()

	router.GET("/", func(w http.ResponseWriter, r *http.Request) {

		users := []User{}
		if err := db.Find(&users).Error; err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": err.Error()})
			return
		}

		marmoset.Render(w).HTML("index", map[string]interface{}{
			"users": users,
		})
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request) {

		user := &User{}
		user.Name = r.FormValue("name")
		user.Age, err = strconv.Atoi(r.FormValue("age"))
		if err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": err.Error()})
			return
		}

		if err := db.Save(user).Error; err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": err.Error()})
			return
		}

		users := []User{}
		if err := db.Find(&users).Error; err != nil {
			marmoset.Render(w).HTML("index", map[string]interface{}{"error": err.Error()})
			return
		}

		marmoset.Render(w).HTML("index", map[string]interface{}{
			"users": users,
		})

	})

	http.Handle("/", router)
}
