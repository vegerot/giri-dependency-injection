package main

import "log"

type UserRecord struct {
	UserID uint64 `gorm:"column:userid;"`
	Name   string `gorm:"column:name;"`
}
type UserRepository interface {
	GetUserByID(id uint64) (*UserRecord, error)
	CreateUser(user UserRecord)
	Initialize()
}

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
