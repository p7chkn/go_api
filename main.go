package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"../models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

func main() {

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://postgres:postgres@127.0.0.1:5432/p7chkn")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	var users []*User
	err = pgxscan.Select(ctx, conn, &users, `SELECT * FROM users_user`)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range users {
		b, err := json.Marshal(value)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}

}
