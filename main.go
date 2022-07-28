package main

import (
	"go-chat/database"
)

func main() {
	dbConn := database.OpenConnection()
	postgresDB, _ := dbConn.DB()
	defer postgresDB.Close()

	// if err := database.CreateTables(dbConn); err != nil {
	// 	log.Fatalln("Create table failue")
	// }

	// if err := database.MigrateTables(dbConn); err != nil {
	// 	log.Fatalln("Migrate table failue")
	// }
	// var user models.UserModel
	// user =
}
