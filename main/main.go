package main

import (
	"goku.net/framework/commons"
	"goku.net/framework/database"
	"goku.net/framework/network/http"
	"goku.net/services/user"
)

func main() {

	commons.InitLogger()

	database.Init(configs []*config.Mysql)

	server := http.NewServer(8989)

	server.AddModule(user.NewExecutorFactory())

	server.Start()
}
