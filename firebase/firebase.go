package myfirebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	myfirebase "github.com/hiroki-kondo-git/dayMemo_api_go/firebase"
	"google.golang.org/api/option"
)

func InitFirebaseAuth() {
	ctx := context.Background()
	fmt.Println("aaaaaaaa")

	opt := option.WithCredentialsFile("daymemo-c5df2-firebase-adminsdk-5hwvd-7fc642a4a1.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panic(fmt.Errorf("error initializing app: %v", err))
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth clint: %v`\n", err)
	}
	myfirebase.CreateUser(ctx, client)
	// defer client.Close()
}
