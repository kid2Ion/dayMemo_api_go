package testfirebase

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/auth"
)

func CreateUser(ctx context.Context, client *auth.Client) *auth.UserRecord {
	params := (&auth.UserToCreate{}).
		Email("useraa@example.coっazzasm5").
		Password("examplepasっsszzaaaa")
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo.UID)
	token, err := client.CustomToken(ctx, u.UserInfo.UID)
	if err != nil {
		fmt.Errorf("fatal get token from uid", err)
	}

	log.Printf("token: %v\n", token)
	return u
}

func UpdateUser(ctx context.Context, client *auth.Client, uid string) {
	params := (&auth.UserToUpdate{}).
		Email("user@exampleUpdate.com1")
	u, err := client.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error updating user:%v\n", err)
	}
	log.Printf("succesfully updating user:%v\n", u)
}

func DeleatUser(ctx context.Context, client *auth.Client, uid string) {
	err := client.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error deleat user:%v\n", err)
	}
	log.Printf("successfully deleat user:%s\n", uid)
}

func GetUser(ctx context.Context, client *auth.Client, uid string) *auth.UserRecord {
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error get user:%v\n", err)
	}
	log.Fatalf("successfully get user:%#v\n", u.UserInfo)
	return u
}
