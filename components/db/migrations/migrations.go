package main

import (
	db2 "developit-golang-rest-api/components/db"
	"fmt"
)

func main() {
	connection := db2.Connect()
	defer connection.Close()

	fmt.Println("*** start migration ***")

	fmt.Println("--- createUserTableMigration")
	createUserTableMigration := &CreateUserTableMigration{connection}
	createUserTableMigration.Up()

	fmt.Println("*** end migration ***")
}
