package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go_v/common"
	"go_v/router"
)

func main() {
	common.InitDB()
	db := common.GetDB()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
