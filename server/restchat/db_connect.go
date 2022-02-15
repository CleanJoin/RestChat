package restchat

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func connectDB() (connect *pgxpool.Pool) {
	urlExample := "postgres://restChat:qweasd123@localhost:5432/restChatDB"
	connect, err := pgxpool.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return connect
}
