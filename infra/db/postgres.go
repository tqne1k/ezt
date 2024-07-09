package db

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dbHost string, dbPort string, dbUser string, dbPasswd string, dbName string) error {
	encodedPassword := url.QueryEscape(dbPasswd)
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s", dbUser, encodedPassword, dbHost, dbPort)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	db.Exec("CREATE DATABASE " + dbName)
	connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, encodedPassword, dbHost, dbPort, dbName)
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	return err
}

func NewPgConnection(dbHost string, dbPort string, dbUser string, dbPasswd string, dbName string) (*gorm.DB, error) {
	encodedPassword := url.QueryEscape(dbPasswd)
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, encodedPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		defer log.Println("Connection to Database Failed")
		log.Fatal(err.Error())
	} else {
		log.Println("Connection to Database Successfully")
	}
	return db, err
}
