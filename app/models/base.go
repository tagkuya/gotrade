package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	_ "golang.org/x/tools/go/cfg"
	"gotrade/config"
	"log"
	"time"
)

const tabeleNameSignalEvents = "signal_events"

var DbConnection *sql.DB

func GetCandleTableName(productcode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productcode, duration)
}

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY KEY NOT NULL,
			product_code STRING,
			side	STRING,
			price FLOAT,
			size  FLOAT)`, tabeleNameSignalEvents)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalf("1コメ %v", err)
	}
	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
			CREATE 	TABLE IF NOT EXISTS %s (
			time  DATETIME PRIMARY KEY NOT NULL,
			open  FLOAT,
			close FLOAT,
			high  FLOAT,
			low open FLOAT,
			volume FLOAT )`, tableName)
		_, err := DbConnection.Exec(c)
		if err != nil {
			log.Fatalf("2コメ %v", err)
		}
	}
}
