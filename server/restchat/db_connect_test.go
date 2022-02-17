package restchat

import (
	"context"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnectDB(t *testing.T) {
	godotenv.Load(".env")
	var inter IConnectDB = NewConnectDB(5432)
	pgxpool, err := inter.Use().Acquire(context.Background())
	if err != nil {
		t.Errorf(err.Error())
	}
	err = pgxpool.Conn().Close(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}
