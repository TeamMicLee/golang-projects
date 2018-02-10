// Gin + Mysql 简单的 Restful 风格的 API
package main

import db "golang-projects/gin-api/database"

func main() {
	defer db.SqlDB.Close()

	router := initRouter()

	router.Run(":1112")
}
