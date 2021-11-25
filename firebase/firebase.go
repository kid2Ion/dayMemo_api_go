package myfirebase

import (
	"context"
	"fmt"
	"log"

	testfirebase "github.com/hiroki-kondo-git/dayMemo_api_go/test"

	firebase "firebase.google.com/go"
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
	// testfirebase.CreateUser(ctx, client)
	// testfirebase.UpdateUser(ctx, client, "3Bl7PjvITAXcgqMCBF5lGTP5k3g1")
	// testfirebase.DeleatUser(ctx, client, "8ofSDj2BFjUSnMaBg31UeQ2KZEl1")
	testfirebase.GetUser(ctx, client, "gy5uqTu10Pg3PnVoUC27pbjRYQq1")
	// defer client.Close()
}
