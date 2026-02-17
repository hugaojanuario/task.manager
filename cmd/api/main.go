package main

import (
	"github.com/hugaojanuario/task.manager.api/database"
	"github.com/hugaojanuario/task.manager.api/http/routes"
)

func main() {
	database.ConectingOnDatabase()
	database.CreateTable()
	routes.HandlerRequest()
}
