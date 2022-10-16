package main

import "log"

func main() {

	// dependency-injection : using PostgreSQL backend

	dbConn := DBConnection{
		Host:         "postgres-instance.company.com",
		User:         "testuser",
		Password:     "testpassword",
		DatabaseName: "testdb",
	}

	databasePostgres := NewSQLDB(dbConn)

	userRepository := NewMySQLUserRepository(databasePostgres)

	userRepository.Initialize()

	myUser := UserRecord{
		UserID: 100,
		Name:   "Jordan Peterson",
	}

	userRepository.CreateUser(myUser)

	userData, err := userRepository.GetUserByID(100)
	if err != nil {
		log.Printf("ERROR : could not fetch user : %v", err.Error())
		return
	}
	log.Printf("user-id : %v", userData.UserID)

	// -------------- MongoDatabase Operations ---------------------

	// dependency-injection : using MongoDB backend

	mongoDBConn := DBMongoConnection{
		MongoURLHost: "mongodb://mongodb-instance.company.com:27017/",
		User:         "testuser",
		Password:     "testpassword",
	}

	databaseMongo := NewMongoDB(mongoDBConn)

	userRepository = NewMyMongoUserRepository(databaseMongo)

	userRepository.Initialize()

	myUser = UserRecord{
		UserID: 200,
		Name:   "Jordan Peterson",
	}

	userRepository.CreateUser(myUser)

	userData, err = userRepository.GetUserByID(200)
	if err != nil {
		log.Printf("ERROR : could not fetch user : %v", err.Error())
		return
	}
	log.Printf("user-id : %v", userData.UserID)

}
