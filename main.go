package main

import (
	"ginchat/router"
	"ginchat/utils"
	// http01 "ginchat/test/http" // for calling http01.Http_template_main()
	// dbtest "ginchat/test/db"  // for calling dbtest.Test_gorm_mysql()
)

func main() {
	// dbtest.Test_gorm_mysql()
	// http01.Http_template_main()

	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	r.Run(":8081") // 1isten and serve on 0.0.0.0:8080 default（forwindows"localhost:8080")
}

// ------------------
