package main

import "log"

func main() {
	dbConn := DBConnection{
		Host:         "postgres-instance.company.com",
		User:         "testuser",
		Password:     "testpassword",
		DatabaseName: "testdb",
	}

	database := NewSQLDB(dbConn)

	userRepository := NewMySQLUserRepository(database)

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

}
