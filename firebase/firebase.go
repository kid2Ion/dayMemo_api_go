package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
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
	createUser(ctx, client)
	// defer client.Close()
}

func createUser(ctx context.Context, client *auth.Client) *auth.UserRecord {
	params := (&auth.UserToCreate{}).
		Email("user@example.com2").
		Password("examplepass2")
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)
	return u
}
