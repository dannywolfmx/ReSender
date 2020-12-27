package main

import (
	"fmt"
	"log"

	"github.com/dannywolfmx/ReSender/server"
	"github.com/spf13/viper"
)

type serverConfig struct{
	Port int
}

func main() {
	configServer := &serverConfig{}
	config(configServer)
	app := server.NewApp()
	app.Run(fmt.Sprintf(":%d", configServer.Port))
}

func config(serverConfig *serverConfig){
	configFile := "config.yaml"
	viper.SetConfigFile(configFile)
	//Read the content of the config file
	if err := viper.ReadInConfig(); err != nil{
		//Report error
		log.Fatalf("Error while reading the %s file: %s", configFile, err)
	}
	port, ok := viper.Get("server.port").(int);
	if	!ok{
		log.Fatal("Error reading the port value")
	}
	//Set port
	serverConfig.Port = port

}
