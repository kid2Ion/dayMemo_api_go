package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type Credentials struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

var client *auth.Client

func init() {
	err := godotenv.Load("env/local.env")
	if err != nil {
		panic("error loading .env file")
	}
	ctx := context.Background()
	firebaseCredentials := Credentials{os.Getenv("FB_TYPE"), os.Getenv("FB_PROJECT_ID"), os.Getenv("FB_PRIVATE_KEY_ID"), os.Getenv("FB_PRIVATE_KEY"), os.Getenv("FB_CLIENT_EMAIL"), os.Getenv("FB_CLIENT_ID"), os.Getenv("FB_AUTH_URI"), os.Getenv("FB_TOKEN_URI"), os.Getenv("FB_AUTH_PROVIDER_X509_CERT_URL"), os.Getenv("FB_CLIENT_X509_CERT_URL")}
	firebaseCredentialsJSON, err := json.Marshal(firebaseCredentials)
	credentials, err := google.CredentialsFromJSON(ctx, []byte(firebaseCredentialsJSON))
	if err != nil {
		log.Panic(fmt.Errorf("error firebase credentials from json: %v\n", err))
	}
	opt := option.WithCredentials(credentials)
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
