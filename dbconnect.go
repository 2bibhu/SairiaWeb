package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
	"log"
)

const (
	MongoDBHosts = "ds023078.mlab.com:23078"
	AuthDatabase = "siaria"
	AuthUserName = "siaria"
	AuthPassword = "Radhaswam1"
	TestDatabase = "siaria"
)

var db *sql.DB
var mongoSession *mgo.Session

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://pqivxbpkudvusc:2KItrACTlbbEJN93-L-sF1g-n3@ec2-107-21-101-67.compute-1.amazonaws.com:5432/d30vbn7l6qoba8")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal(err)
	}

}
