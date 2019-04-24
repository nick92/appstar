package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

// Apps table structure
type Apps struct {
	ID       int
	Appname  string
	Category string
}

// Actions table structure
type Actions struct {
	ID       int
	Appname  string
	Count    int
	Datetime string
}

// Stars table structure
type Stars struct {
	ID       int
	Appname  string
	Stars    int
	Datetime string
}

var applications []Apps
var err error
var db *sql.DB
var starsTable string
var appsTable string
var actionTable string

func getUsefulApps(c *gin.Context) {
	//c.JSON(200, applications)
	applications = nil

	rows, err := db.Query(fmt.Sprintf(`select * from %s where category='%s'`, appsTable, "useful"))

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var appname string
		var catregory string

		if err := rows.Scan(&id, &appname, &catregory); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		applications = append(applications, Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getOfficeApps(c *gin.Context) {
	//c.JSON(200, applications)
	applications = nil

	rows, err := db.Query(fmt.Sprintf(`select * from %s where category='%s'`, appsTable, "office"))

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var appname string
		var catregory string

		if err := rows.Scan(&id, &appname, &catregory); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		applications = append(applications, Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getDevelopmentApps(c *gin.Context) {
	//c.JSON(200, applications)
	applications = nil

	rows, err := db.Query(fmt.Sprintf(`select * from %s where category='%s'`, appsTable, "development"))

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var appname string
		var catregory string

		if err := rows.Scan(&id, &appname, &catregory); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		applications = append(applications, Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getMulitmediaApps(c *gin.Context) {
	//c.JSON(200, applications)
	applications = nil

	rows, err := db.Query(fmt.Sprintf(`select * from %s where category='%s'`, appsTable, "multimedia"))

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var appname string
		var catregory string

		if err := rows.Scan(&id, &appname, &catregory); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		applications = append(applications, Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getGameApps(c *gin.Context) {
	//c.JSON(200, applications)
	//catregory := c.Query("category")

	applications = nil

	rows, err := db.Query(fmt.Sprintf(`select * from %s where category='%s'`, appsTable, "games"))

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var appname string
		var catregory string

		if err := rows.Scan(&id, &appname, &catregory); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		applications = append(applications, Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func starApplication(c *gin.Context) {
	name := c.Query("name")
	counter := 0

	rows, err := db.Query(fmt.Sprintf(`select * from %s where appname='%s'`, starsTable, name))

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	if rows == nil {
		res, inserterr := db.Exec(fmt.Sprintf(`insert into %s values('%s', 1, %s)`, starsTable, name, time.Now()))

		if inserterr != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error: %q", err))
			return
		}

		if res != nil {
			c.String(http.StatusOK,
				fmt.Sprintf("Star row inserted: %q", res))
			return
		}

	}

	defer rows.Close()
	for rows.Next() {
		var appname string
		var stars int
		var updatetime string

		if err := rows.Scan(&appname, &stars, &updatetime); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}

		_, err = db.Exec(fmt.Sprintf(`update %s set stars=%d where appname='%s'`, starsTable, stars+1, appname))

		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error: %q", err))
			return
		}

		db.Exec(fmt.Sprintf(`insert into %s values('%s', 1, '%s')`, actionTable, name, time.Now().Format(time.RFC3339)))
		counter++
	}

	if counter == 0 {
		_, err := db.Exec(fmt.Sprintf(`insert into %s values('%s', 1, '%s')`, starsTable, name, time.Now().Format(time.RFC3339)))

		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error: %q", err))
			return
		}

		_, err = db.Exec(fmt.Sprintf(`insert into %s values('%s', 1, '%s')`, actionTable, name, time.Now().Format(time.RFC3339)))

		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error: %q", err))
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "app stared!"})
}

func getPing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	/*recommendedJson := `{"applications":{"name":"firefox","description":"for browsing web stuffs", "type": "browser", "rating": 100}}`
	var result map[string]interface{}

	json.Unmarshal([]byte(recommendedJson), &result)*/

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	starsTable = pq.QuoteIdentifier("stars")
	appsTable = pq.QuoteIdentifier("apps")
	actionTable = pq.QuoteIdentifier("action")

	//rows, err := db.Query(`SELECT name FROM users WHERE favorite_fruit = $1 OR age BETWEEN $2 AND $2 + 3`, "orange", 64)

	//applications = append(applications, Application{Name: "firefox", Description: "for browsing awesome web stuff", Type: "Web Browser", Stars: 1})

	r := gin.Default()

	r.GET("/ping", getPing)
	r.GET("/packages/useful", getUsefulApps)
	r.GET("/packages/office", getOfficeApps)
	r.GET("/packages/development", getDevelopmentApps)
	r.GET("/packages/multimedia", getMulitmediaApps)
	r.GET("/packages/games", getGameApps)

	r.POST("/packages/appstar", starApplication)

	r.Run(":" + port) // listen and serve on 0.0.0.0:port
}
