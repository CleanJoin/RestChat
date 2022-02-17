package restchat

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ConnectDB struct {
	user     string
	password string
	host     string
	port     int
	dbname   string
}
type IConnectDB interface {
	Use() (connect *pgxpool.Pool)
}

func NewConnectDB(port int) *ConnectDB {
	connectDB := new(ConnectDB)
	connectDB.user = os.Getenv("POSTGRES_USER")
	connectDB.password = os.Getenv("POSTGRES_PASSWORD")
	connectDB.host = os.Getenv("POSTGRES_HOST")
	connectDB.port = port
	connectDB.dbname = os.Getenv("POSTGRES_DB")
	return connectDB

}

func (connectDB *ConnectDB) Use() (connect *pgxpool.Pool) {
	port := strconv.Itoa(connectDB.port)
	var url string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", connectDB.user, connectDB.password, connectDB.host, port, connectDB.dbname)
	connect, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return connect
}
