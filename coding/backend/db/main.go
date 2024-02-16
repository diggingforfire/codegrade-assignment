package db

import (
	"context"
	"log"

	"codegrade.com/take-home/ent"
	_ "github.com/mattn/go-sqlite3"
)

var Client *ent.Client

func Connect(ctx context.Context) {
	var err error
	Client, err = ent.Open("sqlite3", "app.db:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed to create ent: %v", err)
	}

	if err := Client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
