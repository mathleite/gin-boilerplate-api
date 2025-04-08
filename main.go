package main

import (
	"github.com/mathleite/gin-boilerplate-api/database"
	"github.com/mathleite/gin-boilerplate-api/router"
)

func main() {
	database.Setup()
	router.Serve()
}
