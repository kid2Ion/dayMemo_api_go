package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

var client *auth.Client

func init() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("firebase-sdk.json")
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
	// もしから文字だったら400を返す
	if authTokenFromHeader == "" {
		return "", &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "idToken is empty",
		}
	}
	idToken := strings.Replace(authTokenFromHeader, "Bearer ", "", 1)

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		// TODO: 無効なトークンなら401を返す
		return "", &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid idToken",
		}
	}
	uid := token.UID

	return uid, nil
}
