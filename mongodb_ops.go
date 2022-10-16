package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	MongoCollection = "UserRecordsCollection"
	MongoDatabase   = "UserRecordsDB"
)

type Record struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID uint64             `bson:"userid"`
	Name   string             `bson:"name"`
}

type IRecord interface {
	GetPGRecord() UserRecord
}

func (r Record) GetPGRecord() UserRecord {
	record := UserRecord{
		UserID: r.UserID,
		Name:   r.Name,
	}
	return record
}

func MongoConnect() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	credential := options.Credential{
		Username: "testuser",
		Password: "testpassword",
	}

	clientOptions := options.Client().ApplyURI("mongodb://mongodb-instance.company.com:27017/").SetAuth(credential)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("unable to create a client %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("unable to appvars connection %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to mongoconnect %v", err)
	}
	return client
}

func MongoFetchRecord(userID uint64) (Record, error) {
	var filter interface{}
	var myData Record
	errorMessage := ""

	c := MongoConnect()
	ctx := context.Background()
	defer func() {
		_ = c.Disconnect(ctx)
	}()

	filter = bson.D{{"userid", userID}}

	dbCollection := c.Database(MongoDatabase).Collection(MongoCollection)

	var result Record

	err := dbCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		errorMessage = fmt.Sprintf("no record found in the DB")
		log.Printf(errorMessage)
		return myData, errors.New(errorMessage)
	}

	PrettyPrintData(result)

	return result, nil
}

func MongoListRecords() ([]Record, error) {

	dataList := make([]Record, 0)

	var filter interface{}

	c := MongoConnect()
	ctx := context.Background()
	defer func() {
		_ = c.Disconnect(ctx)
	}()

	filter = bson.D{}

	dbCollection := c.Database(MongoDatabase).Collection(MongoCollection)
	rs, err := dbCollection.Find(ctx, filter)
	if err != nil {
		return dataList, errors.New(fmt.Sprintf("failed to fetch all records : %v", err.Error()))
	}

	err = rs.All(ctx, &dataList)
	if err != nil {
		return dataList, errors.New(fmt.Sprintf("failed to fetch all records : %v", err.Error()))
	}
	if len(dataList) == 0 {
		emptyListOfRecords := make([]Record, 0)
		return emptyListOfRecords, nil
	}

	PrettyPrintData(dataList)

	return dataList, nil
}

func MongoCreateOrUpdateRecord(userID uint64, newValue Record) (bool, error) {
	c := MongoConnect()
	ctx := context.Background()
	defer func() {
		_ = c.Disconnect(ctx)
	}()

	dbCollection := c.Database(MongoDatabase).Collection(MongoCollection)
	filter := bson.D{{"userid", userID}}
	update := bson.D{{"$set", newValue}}

	_, err := dbCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return false, errors.New(fmt.Sprintf("failed to update record with userid (%v) : %v", userID, err.Error()))
	}
	return true, nil
}

func MongoDeleteRecord(userID uint64) (bool, error) {
	errorMessage := ""
	c := MongoConnect()
	ctx := context.Background()
	defer func() {
		_ = c.Disconnect(ctx)
	}()
	isssueCollection := c.Database(MongoDatabase).Collection(MongoCollection)
	filter := bson.D{{"userid", userID}}
	deletionResult, err := isssueCollection.DeleteOne(ctx, filter)
	if err != nil {
		errorMessage = fmt.Sprintf("failed to delete record : %v", userID)
		log.Printf(errorMessage)
		return false, errors.New(errorMessage)
	}
	numberOfRecordsDeleted := deletionResult.DeletedCount
	if numberOfRecordsDeleted == 0 {
		log.Printf("nothing was deleted, because nothing was found")
		return false, nil
	} else {
		log.Printf("number of records deleted : %v", numberOfRecordsDeleted)
		return true, nil
	}
}
