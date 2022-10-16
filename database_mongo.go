package main

type DBMongoConnection struct {
	MongoURLHost string
	User         string
	Password     string
}

type MyMongoDatabase struct {
	dbconnection *DBMongoConnection
}

func NewMongoDB(conn DBMongoConnection) *MyMongoDatabase {
	return &MyMongoDatabase{dbconnection: &conn}
}
