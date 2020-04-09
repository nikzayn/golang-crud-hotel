// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

/* This file contains the definition of AppContext */

const (
	//DbHost is the environment variable storing the database access url
	DbHost = "DB_HOST"
	//DbPort is the environment variable storing the database access port
	DbPort = "DB_PORT"
	//DbDatabaseName is the environment variable storing the database name
	DbDatabaseName = "DB_DATABASE_NAME"
	//DbUsername is the environment variable storing the database username
	DbUsername = "DB_USERNAME"
	//DbPassword is the environment variable storing the database password
	DbPassword = "DB_PASSWORD"
)

//DbConfig is the database configuration to connect to it
type DbConfig struct {
	//Host to be used to connect to the database
	Host string
	//Port with which the database can be accessed
	Port string
	//Database to connect
	Database string
	//Username to access the connection
	Username string
	//Password to access the connection
	Password string
}

//NewDbConfig will read the db config from the os environment variables and set it in the config
func NewDbConfig() *DbConfig {
	dbC := &DbConfig{
		Host:     "localhost",
		Port:     "3306",
		Database: "hotel",
		Username: "root",
		Password: "hotel123",
	}
	if len(os.Getenv(DbHost)) > 0 {
		dbC.Host = os.Getenv(DbHost)
	}
	if len(os.Getenv(DbPort)) > 0 {
		dbC.Port = os.Getenv(DbPort)
	}
	if len(os.Getenv(DbDatabaseName)) > 0 {
		dbC.Database = os.Getenv(DbDatabaseName)
	}
	if len(os.Getenv(DbUsername)) > 0 {
		dbC.Username = os.Getenv(DbUsername)
	}
	if len(os.Getenv(DbPassword)) > 0 {
		dbC.Password = os.Getenv(DbPassword)
	}
	return dbC
}

//Connect will connect the database. Will return an error if anything comes up else nil
func (d DbConfig) Connect() (*gorm.DB, error) {
	/*
	 * We will build the connection string
	 * Then will connect to the database
	 */
	cStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.Username, d.Password, d.Host, d.Port, d.Database)

	return gorm.Open("mysql", cStr)
}

//AppContext contains the
type AppContext struct {
	//Db is the database connection
	Db *gorm.DB
	//Log for logging purposes
	Log Logger
}

var rootAppContext *AppContext

func init() {
	/*
	 * We will initialize the context
	 * We will connect to the database
	 */
	rootAppContext = &AppContext{}

	err := rootAppContext.ConnectToDB()
	if err != nil {
		log.Fatal("Error while creating the root app context. Connecting to DB failed. ", err)
	}
	log.Println("INFO: Connected to the database")
}

//NewAppContext returns an initlized app context
func NewAppContext(l Logger) *AppContext {
	return &AppContext{Log: l, Db: rootAppContext.Db}
}

//ConnectToDB connects the database and updates the Db property of the context as new connection
//If any error happens in between , it will be returned and connection won't be set in the context
func (a *AppContext) ConnectToDB() error {
	/*
	 * We will get the db config
	 * Connect to it
	 * If no error then set the database connection
	 */
	c := NewDbConfig()
	d, err := c.Connect()
	if err == nil {
		a.Db = d
	}
	return err
}

//AddModels adds an model to database
func AddModels(model interface{}) {
	rootAppContext.Db.AutoMigrate(model)
}
