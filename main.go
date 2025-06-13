package main

import "ginchat/router"

func main() {
	// dbtest.Test_gorm_mysql()

	r := router.Router()
	r.Run(":8081") // 1isten and serve on 0.0.0.0:8080 default（forwindows"localhost:8080")
}

// ------------------
