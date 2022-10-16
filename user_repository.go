package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserRecord struct {
	UserID uint64 `gorm:"column:userid;"`
	Name   string `gorm:"column:name;"`
}

type IUserRecord interface {
	GetMongoRecord() Record
}

func (r UserRecord) GetMongoRecord() Record {
	record := Record{
		ID:     primitive.ObjectID{},
		UserID: r.UserID,
		Name:   r.Name,
	}
	return record
}

type UserRepository interface {
	GetUserByID(id uint64) (*UserRecord, error)
	CreateUser(user UserRecord)
	Initialize()
}

/* ******************************************************************************************************* */

type mySQLUserRepository struct {
	Database *MySQLDatabase
}

func (m mySQLUserRepository) Initialize() {
	err := m.Database.dbconnection.SQL().Debug().AutoMigrate(&UserRecord{})
	if err != nil {
		log.Printf("ERROR : could not auto-migrate : %v", err.Error())
	}
}

func (m mySQLUserRepository) CreateUser(user UserRecord) {
	m.Database.dbconnection.SQL().Debug().Create(&user)
}

func (m mySQLUserRepository) GetUserByID(id uint64) (*UserRecord, error) {
	user := UserRecord{
		UserID: 0,
		Name:   "",
	}
	m.Database.dbconnection.SQL().Debug().Where(map[string]interface{}{"userid": id}).Find(&user)
	log.Printf("user fetched : %v", user)
	return &user, nil
}

func NewMySQLUserRepository(d *MySQLDatabase) UserRepository {
	return &mySQLUserRepository{Database: d}
}

/* ******************************************************************************************************* */

type myMongoUserRepository struct {
	Database *MyMongoDatabase
}

func (m myMongoUserRepository) GetUserByID(id uint64) (*UserRecord, error) {
	record, err := MongoFetchRecord(id)
	if err != nil {
		log.Printf("ERROR : %v", err.Error())
		return nil, err
	}
	userRecord := record.GetPGRecord()
	return &userRecord, nil
}

func (m myMongoUserRepository) CreateUser(user UserRecord) {
	mongoRecord := user.GetMongoRecord()
	recordCreateOrUpdateStatus, err := MongoCreateOrUpdateRecord(mongoRecord.UserID, mongoRecord)
	if err != nil {
		log.Printf("ERROR : %v", err.Error())
	}
	log.Printf("recordCreateOrUpdateStatus : %v", recordCreateOrUpdateStatus)
}

func (m myMongoUserRepository) Initialize() {
	log.Printf("Nothing To Initialize For MongoDB...")
}

func NewMyMongoUserRepository(d *MyMongoDatabase) UserRepository {
	return &myMongoUserRepository{Database: d}
}

/* ******************************************************************************************************* */
