package main

import (
	"context"
	"fmt"

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

type ExpoPushToken struct {
	// Shows if the notification subscription is active
	Active bool
	// Client language
	Language string
	// Client platform (ios or android)
	Platform string
	// Expo Push Notifications Token
	Token string
}

func GetTokens() {
	client := InitClient(context.Background())
	ref := client.NewRef("subscriptions/")
	var data map[string]ExpoPushToken
	err := ref.Get(context.Background(), &data)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(data)
	}
}
