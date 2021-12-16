package main

import (
	"flag"
	"io/ioutil"

	"goku.net/framework/cache"
	"goku.net/framework/commons"
	"goku.net/framework/config"
	"goku.net/framework/database"
	"goku.net/framework/network/http"
	"goku.net/services/user"
	"gopkg.in/yaml.v2"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "", "config file path")
	flag.StringVar(&configFile, "f", "", "config file path")
}

func main() {

	flag.Parse()
	configFileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	config := config.ServerConfig{}
	err = yaml.Unmarshal(configFileContent, &config)
	if err != nil {
		panic(err)
	}

	commons.InitLogger()

	database.Init(config.Mysql)

	cache.Init(config.Redis)

	server := http.NewServer(8989)

	server.AddModule(user.NewExecutorFactory())

	server.Start()
}
