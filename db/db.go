package db

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	dbName   = "daily-beauty"
	collName = "emails"
)

var ctx = context.Background()
var emailsCol *mongo.Collection

// init mongodb connection
func init() {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, pass, host, port, dbName)
	client, err := mongo.NewClient(url)
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	emailsCol = client.Database(dbName).Collection(collName)
}

func checkExist(email string) (bool, error) {
	filter := map[string]string{"email": email}
	n, err := emailsCol.Count(ctx, filter)
	return n != 0, err
}

// InsertAEmail save a email into the database
func InsertAEmail(email string) error {
	// check email duplicate
	isExist, err := checkExist(email)
	if err != nil {
		return err
	}
	if isExist {
		return nil
	}

	data := map[string]string{"email": email}
	_, err = emailsCol.InsertOne(ctx, data)
	// if success, err will be nil
	return err
}

// RemoveAEmail removes a email from database
func RemoveAEmail(email string) error {
	filter := map[string]string{"email": email}
	_, err := emailsCol.DeleteOne(ctx, filter)
	return err
}

// GetEmails get all emails in the database
func GetEmails() ([]string, error) {
	cur, err := emailsCol.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	emails := []string{}
	for cur.Next(ctx) {
		raw, err := cur.DecodeBytes()
		if err != nil {
			return nil, err
		}
		email := raw.Lookup("email").StringValue()
		emails = append(emails, email)
	}
	return emails, nil
}
