package main

import (
	"errors"
	"fmt"
	"github.com/Omnia-Wahid/sql-client.git/sqlclient"
)

const (
	queryGetUser = "SELECT id,email from users where id=%d;"
)

var (
	dbClient sqlclient.SqlClient
)

type User struct {
	Id    int64
	Email string
}

func init() {
	var err error
	dbClient, err = sqlclient.Open("mysql", "this is the connection string")
	if err != nil {
		panic(err)
	}
}

func main() {
	user, err := GetUser(123)
	if err != nil {
		panic(err)
	}
	fmt.Print(user.Email)
}

func GetUser(id int64) (*User, error) {
	rows, err := dbClient.Query(fmt.Sprintf(queryGetUser, id))
	if err != nil {
		return nil, err
	}
	var user User
	defer rows.Close()
	for rows.HasNext() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("User not found")
}
