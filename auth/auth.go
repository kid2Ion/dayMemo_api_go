package auth

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

	// defer client.Close()
}

func AuthFirebase(ctx echo.Context) (string, error) {
	authTokenFromHeader := ctx.Request().Header.Get("Authorization")
	// TODO: もしから文字だったら400を返す
	idToken := strings.Replace(authTokenFromHeader, "Bearer ", "", 1)
	fmt.Println("idtoken:", idToken)

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		// TODO: 無効なとくんなら401を返す
		fmt.Println(err)
		return "", err
	}
	uid := token.UID

	return uid, nil
}
