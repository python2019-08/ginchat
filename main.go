package main

import (
	"ginchat/router"
	// dbtest "ginchat/test/db"  // for dbtest.Test_gorm_mysql()
)

func main() {
	// dbtest.Test_gorm_mysql()

	r := router.Router()
	r.Run(":8081") // 1isten and serve on 0.0.0.0:8080 default（forwindows"localhost:8080")
}

// ------------------
