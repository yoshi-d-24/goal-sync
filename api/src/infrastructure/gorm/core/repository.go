package core

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func CreateDB() *gorm.DB {
	if db != nil {
		return db
	}

	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
			getHost(), getUser(), getPassword(), getDBName(), getPort())))

	if err != nil {
		panic(err)
	}

	return db
}

func getUser() string {
	user, userExists := os.LookupEnv("POSTGRES_USER")
	if userExists {
		return user
	}
	return "postgres"
}

func getPassword() string {
	password, passwordExists := os.LookupEnv("POSTGRES_PASSWORD")
	if passwordExists {
		return password
	}
	return "postgres"
}

func getHost() string {
	host, hostExists := os.LookupEnv("POSTGRES_HOST")
	if hostExists {
		return host
	}
	return "db"
}

func getPort() string {
	port, portExists := os.LookupEnv("POSTGRES_PORT")
	if portExists {
		return port
	}
	return "5432"
}

func getDBName() string {
	dbname, dbnameExists := os.LookupEnv("POSTGRES_DBNAME")
	if dbnameExists {
		return dbname
	}
	return "postgres"
}
