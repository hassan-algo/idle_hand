package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"example.com/extras"
)

type DatabaseConnection struct {
	psqlInfo      string
	Con           *sql.DB
	lastConnected time.Time
}

func NewDatabaseConnection() *DatabaseConnection {
	log.Println("Creating database connection")
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
		// psqlInfo = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname
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
		// psqlInfo = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname
		log.Println("Connected to production DB")
	}
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		// log.Fatal("Error Connecting to DB", err)
		panic(err)
	} else if err := db.Ping(); err != nil {
		panic(err)
	}
	// db.SetMaxOpenConns(50)
	// defer db.Close()
	return &DatabaseConnection{
		psqlInfo:      psqlInfo,
		Con:           db,
		lastConnected: time.Now(),
	}
}

func (db *DatabaseConnection) HandleReconnect() error {
	log.Println("Re-connecting database")
	previousDb := db.Con
	tempdb, err := sql.Open("postgres", db.psqlInfo)
	if err != nil {
		return err
	} else if err := tempdb.Ping(); err != nil {
		return err
	}
	db.Con = tempdb
	db.lastConnected = time.Now()
	previousDb.Close()
	return nil
}

func (db *DatabaseConnection) CheckTimeOut() error {
	passedTime := time.Since(db.lastConnected).Seconds()
	if (passedTime) > 60 || db.Con == nil {
		err := db.HandleReconnect()
		if err != nil {
			return err
		}
	}
	return nil
}
