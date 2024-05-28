package connection_test

import (
	"context"
	"fmt"
	"testing"

	connection "github.com/Zeroaril7/learn-go/database"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetConnection(t *testing.T) {
	connection.GetConnection()
}

func TestInsertSQL(t *testing.T) {

	var db = connection.GetConnection()
	var ctx = context.Background()

	query := "INSERT INTO presence(name) VALUES ('Wawan');"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}
}

func TestGetSQL(t *testing.T) {
	var db = connection.GetConnection()
	var ctx = context.Background()

	query := "SELECT * FROM presence;"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {

		var id, name, datetime string

		err = rows.Scan(&id, &name, &datetime)
		if err != nil {
			panic(err)
		}

		fmt.Println(datetime)
	}

	defer rows.Close()

}

func TestGetLastIdSQL(t *testing.T) {
	var db = connection.GetConnection()
	var ctx = context.Background()

	query := "SELECT * FROM presence;"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

}
