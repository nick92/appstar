package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/nick92/appstar/common"
)

var starsTable = pq.QuoteIdentifier("stars")
var actionTable = pq.QuoteIdentifier("action")

// Stars table structure
type Stars struct {
	ID       int
	Appname  string
	Stars    int
	Datetime string
}

type StarsResponse struct {
	AppName string `json:"app_name"`
	Stars   int    `json:"stars"`
}

func GetAllAppStars(appName string) (*sql.Rows, error) {
	db := common.GetDB()
	var rows *sql.Rows

	rows, err := db.Query(fmt.Sprintf(`select * from %s where appname='%s'`, starsTable, appName))

	if err != nil {
		return rows, err
	}

	return rows, nil
}

func GetMostPopularApps() (*sql.Rows, error) {
	db := common.GetDB()
	var rows *sql.Rows

	rows, err := db.Query(fmt.Sprintf(`select * from %s order by stars desc limit 20`, starsTable))

	if err != nil {
		return rows, err
	}

	return rows, nil
}

func GetAllStars() (*sql.Rows, error) {
	db := common.GetDB()
	var rows *sql.Rows

	rows, err := db.Query(fmt.Sprintf(`select * from %s where stars > 0`, starsTable))

	if err != nil {
		return rows, err
	}

	return rows, nil
}

func AddStarToApp(appName string) error {
	db := common.GetDB()
	var counter int

	rows, err := db.Query(fmt.Sprintf(`select * from %s where appname='%s'`, starsTable, appName))

	if err != nil {
		return err
	}

	if rows == nil {
		res, inserterr := db.Exec(fmt.Sprintf(`insert into %s values('%s', 1, %s)`, starsTable, appName, time.Now()))

		if inserterr != nil {
			return inserterr
		}

		if res != nil {
			return nil
		}
	}

	defer rows.Close()
	for rows.Next() {
		var appname string
		var stars int
		var updatetime string

		if err := rows.Scan(&appname, &stars, &updatetime); err != nil {
			return err
		}

		_, err = db.Exec(fmt.Sprintf(`update %s set stars=%d where appname='%s'`, starsTable, stars+1, appname))

		if err != nil {
			return err
		}

		db.Exec(fmt.Sprintf(`insert into %s values('%s', 1, '%s')`, actionTable, appname, time.Now().Format(time.RFC3339)))
		counter++
	}

	return nil
}
