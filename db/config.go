package db

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type dbConfig struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}

func NewDBConfig(userENV, passENV, hostENV, portENV, nameENV string) *dbConfig {
	dbConf := new(dbConfig)
	dbConf.dbUser = os.Getenv(userENV)
	dbConf.dbPass = os.Getenv(passENV)
	dbConf.dbHost = os.Getenv(hostENV)
	dbConf.dbPort = os.Getenv(portENV)
	dbConf.dbName = os.Getenv(nameENV)
	return dbConf
}

func dbInit() *sqlx.DB {
	db := NewDBConfig("DBUSER", "DBPASS", "DBHOST", "DBPORT", "DBNAME")
	creds := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=disable",
		db.dbUser, db.dbPass, db.dbHost, db.dbPort, db.dbName)

	return sqlx.MustOpen("pgx", creds)
}

// PingWithTimeout ...
func PingWithTimeout(db *sqlx.DB) error {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func Init() *sqlx.DB {
	PGDB := dbInit()
	return PGDB
}
