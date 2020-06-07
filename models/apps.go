package models

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/nick92/appstar/common"
)

var appsTable = pq.QuoteIdentifier("apps")

// Apps table structure
type Apps struct {
	ID       int
	Appname  string
	Category string
}

type AppRequest struct {
	AppName  string `json:"app_name"`
	Category string `json:"category"`
}

// Actions table structure
type Actions struct {
	ID       int
	Appname  string
	Count    int
	Datetime string
}

func GetAllAppsForCategory(category string) (*sql.Rows, error) {
	db := common.GetDB()
	var rows *sql.Rows

	rows, err := db.Query(fmt.Sprintf(`select * from %s where category='%s'`, appsTable, category))

	if err != nil {
		return rows, err
	}

	return rows, nil
}

func InsertAppForCategory(appName string, category string) error {
	db := common.GetDB()
	_, err := db.Exec(fmt.Sprintf(`insert into %s (appname, category) values('%s', '%s')`, appsTable, appName, category))

	if err != nil {
		return err
	}

	return nil
}

func DeleteApp(appName string) error {
	db := common.GetDB()
	_, err := db.Exec(fmt.Sprintf(`delete from %s where appname='%s'`, appsTable, appName))

	if err != nil {
		return err
	}

	return nil
}
