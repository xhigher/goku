package main

import (
	"goku.net/framework/commons"
	"goku.net/framework/network/http"
	"goku.net/services/user"
)

func main() {

	commons.InitLogger()
	server := http.NewServer(8989)

	server.AddModule("user", &user.UserFactory{})

	server.Start()
}
