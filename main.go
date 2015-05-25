package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-simple-rest/config"
	"go-simple-rest/resource"
	"go-simple-rest/route"
	"log"
	"net/http"
	"os"
)

func showUsage() {
	fmt.Println("Usage: go run main.go <command>")
	fmt.Println("=============================\n")
	fmt.Println("Avaialable commands:")
	fmt.Println("go run main.go start    # run server")
	fmt.Println("go run main.go test # run test")
	fmt.Println("go run main.go deploy # run as deploy on host")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid command usage\n")
		showUsage()
		os.Exit(1)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	arg := os.Args[1]
	switch arg {
	case "start":
		viper.Set("env", "development")
	case "deploy":
		viper.Set("env", "production")
	case "test":
		viper.Set("env", "testing")
	default:
		fmt.Println("Invalid command:", arg)
		showUsage()
		os.Exit(1)
	}

	router := route.Router()
	err = resource.InitDB()

	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(config.GetServer("host")+":"+config.GetServer("port"), router))
}
