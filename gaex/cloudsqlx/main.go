package cloudsqlx

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	projectID := os.Getenv("PROJECT_ID")
	instanceName := os.Getenv("INSTANCE_NAME")
	databaseName := os.Getenv("DATABASE_NAME")

	db, err := gorm.Open("mysql", fmt.Sprintf(
		"root@cloudsql(%s:%s)/%s", projectID, instanceName, databaseName,
	))
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
