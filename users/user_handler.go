package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/p7chkn/go_api/models"
)

func ReturnUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /all")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://postgres:postgres@127.0.0.1:5432/p7chkn")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	var users []*models.User
	err = pgxscan.Select(ctx, conn, &users, `SELECT * FROM users_user`)
	if err != nil {
		fmt.Println(err)
	}
	result := map[string][]*models.User{}
	result["result"] = users
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
