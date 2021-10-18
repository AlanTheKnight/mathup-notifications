package main

import (
	"context"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"google.golang.org/api/option"
)

const databaseURL = "https://mathup-55a32-default-rtdb.firebaseio.com/"

func InitApp(ctx context.Context) *firebase.App {
	// Get account key from JSON file
	opt := option.WithCredentialsFile("./firebase-config.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("⚠️ Couldn't inititalize Firebase: " + err.Error())
	}
	return app
}

func InitClient(ctx context.Context) *db.Client {
	app := InitApp(ctx)
	client, err := app.DatabaseWithURL(ctx, databaseURL)
	if err != nil {
		panic("⚠️ Couldn't create Database instance: " + err.Error())
	}
	return client
}

// Data stored in Firebase RTDB
type ExpoPushToken struct {
	// Shows if the notification subscription is active
	Active bool
	// Client language
	Language string
	// Client platform (ios or android)
	Platform string
	// Expo Push Notifications Token
	Token string
	// The last time when push notification was sent
	Time int
}

func RetrieveTokens() {
	ctx := context.Background()
	client := InitClient(ctx)
	ref := client.NewRef("subscriptions/")

	localDB, read, err := LoadLocalDB("data/data.json")
	if err != nil {
		errorLog.Println("⚠️ Couldn't load local data: " + err.Error())
	} else if read {
		infoLog.Println("✅ Local data was read from file")
	} else {
		errorLog.Println("⚠️ Local data was not read from file")
	}

	old := time.Unix(localDB.TimeStamp, 0)
	new := time.Now()
	needToRefresh := new.Sub(old).Hours() > 3 || !read

	if needToRefresh {
		err = ref.Get(ctx, &localDB.Data)
		if err != nil {
			errorLog.Println("⚠️ Failed to retrieve data from Firebase: " + err.Error())
		}
		infoLog.Printf("✅ Retrieved %d token(s) from Firebase RTDB", localDB.Size())
		localDB.TimeStamp = time.Now().Unix()
		err = SaveLocalDB(localDB)
		if err != nil {
			errorLog.Println("⚠️ Failed to save local data: " + err.Error())
		} else {
			infoLog.Println("✅ Saved local data")
		}
	}
}
