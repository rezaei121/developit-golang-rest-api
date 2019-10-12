package main

import (
	"developit-golang-rest-api/db"
	"fmt"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	fmt.Println("*** start migration ***")

	fmt.Println("--- createUserTableMigration")
	createUserTableMigration := &CreateUserTableMigration{connection}
	createUserTableMigration.Up()

	fmt.Println("*** end migration ***")
}
