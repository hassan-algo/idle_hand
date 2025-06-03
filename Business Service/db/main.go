package db

import (
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"example.com/extras"
)

type DatabaseConnection struct {
	psqlInfo      string
	Con           *gorm.DB
	lastConnected time.Time
}

func NewDatabaseConnection() *DatabaseConnection {
	psqlInfo := ""
	isLocal, err := strconv.ParseBool(extras.GetEnv("isLocal"))
	if err != nil {
		log.Fatal("ERROR: unable to convert isLocal to boolean", err)
	}
	if isLocal {
		var (
			host      = extras.GetEnv("localHost")
			port, err = strconv.Atoi(extras.GetEnv("localPort"))
			user      = extras.GetEnv("localUser")
			password  = extras.GetEnv("localPassword")
			dbname    = extras.GetEnv("localDbname")
		)
		if err != nil {
			log.Fatal("ERROR: Port unable to convert port to int", err)
		}

		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=require",
			host, port, user, password, dbname)
		log.Println("Connected to local DB")
	} else {
		var (
			host     = extras.GetEnv("liveHost")
			port     = extras.GetEnv("livePort")
			user     = extras.GetEnv("liveUser")
			password = extras.GetEnv("livePassword")
			dbname   = extras.GetEnv("liveDbname")
		)
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=require",
			host, port, user, password, dbname)
		log.Println("Connected to production DB")
	}

	// Use GORM to open the database connection
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Get the underlying *sql.DB and ping it
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	return &DatabaseConnection{
		psqlInfo:      psqlInfo,
		Con:           db,
		lastConnected: time.Now(),
	}
}
func (db *DatabaseConnection) HandleReconnect() error {
	previousDb := db.Con
	tempdb, err := gorm.Open(postgres.Open(db.psqlInfo), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := tempdb.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		return err
	}
	db.Con = tempdb
	db.lastConnected = time.Now()
	if previousDb != nil {
		if prevSqlDB, err := previousDb.DB(); err == nil {
			prevSqlDB.Close()
		}
	}
	return nil
}
func (db *DatabaseConnection) CheckTimeOut() error {
	passedTime := time.Since(db.lastConnected).Seconds()
	if (passedTime) > 10 || db.Con == nil {
		err := db.HandleReconnect()
		if err != nil {
			return err
		}
	}
	return nil
}
