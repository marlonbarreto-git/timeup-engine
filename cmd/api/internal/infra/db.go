package infra

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseConn interface {
}

type firebaseConn struct {
	ctx    context.Context
	sa     option.ClientOption
	app    *firebase.App
	client *firestore.Client
}

func NewFirebaseConn() FirebaseConn {
	ctx := context.Background()
	sa := option.WithCredentialsFile("timeup-engine-369607-ff47b7391bc9.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	fmt.Println("Initializing firebase connection")
	return &firebaseConn{ctx, sa, app, client}
}
