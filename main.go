package main

import (
	"ginchat/models"
	"ginchat/test"
)

func main() {
	models.Test_gorm_sqlite()
	test.DB_sqlite3()
}
