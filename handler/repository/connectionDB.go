package repository

import (
	"database/sql"
	"gitlab.com/pplayground/pet_tracking/main-framework/utils"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

func ConnectMongoDB() (*mgo.Session, error) {
	mongoDB := utils.GetConfig("MONGODB_HOST")
	log.Printf("Connecting to MongoDB %v connection\n", mongoDB)
	session, err := mgo.Dial(mongoDB)
	if err != nil {
		return nil, err
	}
	log.Println("Connected")
	return session, nil
}

func DisConnectMongoDB(session *mgo.Session) {
	log.Println("Closing MongoDB connection")
	session.Close()
	log.Println("Closed MongoDB connection")
}

func ConnectSqlDB(dbName string) (*sql.DB, error) {
	sqlDB := utils.GetConfig("MYSQLDB_HOST")
	dbUser := os.Getenv("MYSQLDB_USER")
	dbPwd := os.Getenv("MYSQLDB_PASSWORD")
	log.Printf("Connecting to SqlDB %v connection\n", sqlDB)
	db, err := sql.Open("mysql", dbUser+":"+dbPwd+"@tcp("+sqlDB+")/"+dbName)
	if err != nil {
		return nil, err
	}
	log.Println("Connected")
	return db, nil
}

func DisConnectSqlDB(db *sql.DB) {
	log.Println("Closing SqlDB connection")
	db.Close()
	log.Println("Closed SqlDB connection")
}
