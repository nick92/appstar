package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nick92/appstar/models"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/ping", getPing)
	router.GET("/get_packages", getApps)
	router.GET("/get_popular_packages", getPopularApps)
	router.GET("/get_all_stars", getAllStars)

	// legacy support for old routes
	router.GET("/useful", getUsefulApps)
	router.GET("/office", getOfficeApps)
	router.GET("/development", getDevelopmentApps)
	router.GET("/multimedia", getMulitmediaApps)
	router.GET("/games", getGameApps)

	router.POST("/add_app", addApplication)
	router.POST("/appstar", starApplication)

	router.DELETE("/delete_app", deleteApplication)
}

func getPing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func getUsefulApps(c *gin.Context) {
	var applications []models.Apps

	rows, err := models.GetAllAppsForCategory("useful")

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
				fmt.Sprintf("Error scanning rows: %q", err))
			return
		}
		applications = append(applications, models.Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getOfficeApps(c *gin.Context) {
	var applications []models.Apps

	rows, err := models.GetAllAppsForCategory("office")

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
				fmt.Sprintf("Error scanning rows: %q", err))
			return
		}
		applications = append(applications, models.Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getDevelopmentApps(c *gin.Context) {
	var applications []models.Apps

	rows, err := models.GetAllAppsForCategory("development")

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
				fmt.Sprintf("Error scanning rows: %q", err))
			return
		}
		applications = append(applications, models.Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getMulitmediaApps(c *gin.Context) {
	var applications []models.Apps

	rows, err := models.GetAllAppsForCategory("multimedia")

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
				fmt.Sprintf("Error scanning rows: %q", err))
			return
		}
		applications = append(applications, models.Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getGameApps(c *gin.Context) {
	var applications []models.Apps

	rows, err := models.GetAllAppsForCategory("games")

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
				fmt.Sprintf("Error scanning rows: %q", err))
			return
		}
		applications = append(applications, models.Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getPopularApps(c *gin.Context) {
	var applications []models.Apps

	rows, err := models.GetMostPopularApps()

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var appname string
		var stars string
		var datetime time.Time

		if err := rows.Scan(&appname, &stars, &datetime); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning rows: %q", err))
			return
		}
		applications = append(applications, models.Apps{Appname: appname})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func getApps(c *gin.Context) {
	appType := c.Query("type")
	var applications []models.Apps

	rows, err := models.GetAllAppsForCategory(appType)

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
				fmt.Sprintf("Error scanning rows: %q", err))
			return
		}
		applications = append(applications, models.Apps{ID: id, Appname: appname, Category: catregory})
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func addApplication(c *gin.Context) {
	var appRequest models.AppRequest

	if err := c.BindJSON(&appRequest); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	err := models.InsertAppForCategory(appRequest.AppName, appRequest.Category)

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "app saved!"})
}

func starApplication(c *gin.Context) {
	name := c.Query("name")

	err := models.AddStarToApp(name)

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "app stared!"})
}

func getAllStars(c *gin.Context) {
	var starResponse []models.StarsResponse
	rows, err := models.GetAllStars()

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var appName string
		var appStar int
		var lastUpdated time.Time

		if err := rows.Scan(&appName, &appStar, &lastUpdated); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}

		starRes := models.StarsResponse{
			AppName: appName,
			Stars:   appStar,
		}

		starResponse = append(starResponse, starRes)
	}

	c.JSON(http.StatusOK, gin.H{"result": starResponse})
}

func deleteApplication(c *gin.Context) {
	name := c.Query("name")

	err := models.DeleteApp(name)

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "app deleted!"})
}
