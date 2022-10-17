package main

import (
	"log"
	"testing"
)

func stuff(userRepository UserRepository) (*UserRecord, error) {
	myUser := UserRecord{
		UserID: 100,
		Name:   "Jordan Peterson",
	}

	userRepository.CreateUser(myUser)

	userData, err := userRepository.GetUserByID(100)
	if err != nil {
		return nil, err
	}
	return userData, nil

}

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
	userData, err := stuff(userRepository)
	if err != nil {
		log.Fatalf("ERROR : could not fetch user : %v", err.Error())
	}
	log.Printf("user-id : %v", userData.UserID)
}

func testStuff(t *testing.T) {
	var mockUserRepository UserRepository /* = ??? */

	userData, err := stuff(mockUserRepository)

	if err != nil {
		t.Fatalf("You fucked up: %#v", err)
	}

	got := userData.UserID
	const want = 1234

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
