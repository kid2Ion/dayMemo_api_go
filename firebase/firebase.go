package myfirebase

import (
	"context"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

var client *auth.Client

func init() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("daymemo-c5df2-firebase-adminsdk-5hwvd-7fc642a4a1.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panic(fmt.Errorf("error initializing app: %v", err))
	}

	var error error
	client, error = app.Auth(ctx)
	if error != nil {
		log.Fatalf("error getting Auth clint: %v`\n", err)
	}

	// token, err := client.VerifyIDToken(ctx, idToken)
	// testfirebase.CreateUser(ctx, client)
	// testfirebase.UpdateUser(ctx, client, "3Bl7PjvITAXcgqMCBF5lGTP5k3g1")
	// testfirebase.DeleatUser(ctx, client, "8ofSDj2BFjUSnMaBg31UeQ2KZEl1")
	// testfirebase.GetUser(ctx, client, "gy5uqTu10Pg3PnVoUC27pbjRYQq1")
	// defer client.Close()
}

func AuthFirebase(ctx echo.Context) (string, error) {
	authHeader := ctx.Request().Header.Get("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	fmt.Println("idtoken:", idToken)

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return "", err
	}
	fmt.Println("token:", token)
	uid := token.UID
	fmt.Println("uid:", uid)

	return uid, nil
}
