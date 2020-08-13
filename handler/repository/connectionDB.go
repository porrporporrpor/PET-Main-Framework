package repository

import (
	"gitlab.com/pplayground/pet_tracking/main-framework/utils"
	"gopkg.in/mgo.v2"
	"log"
)

func ConnectMongoDB() (*mgo.Session, error) {
	mongoDB := utils.GetConfig("MONGODB_HOST")
	log.Printf("Connection to MongoDB %v connection\n", mongoDB)
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