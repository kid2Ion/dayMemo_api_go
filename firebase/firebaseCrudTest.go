package firebase

import (
	"context"
	"log"

	"firebase.google.com/go/auth"
)

func CreateUser(ctx context.Context, client *auth.Client) *auth.UserRecord {
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
