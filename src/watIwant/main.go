package main

import (
	"fmt"
	"log"
	"net/http"
	"watIwant/apiserver"
	"watIwant/config"
)

type WatIWant struct{}

func (watIWant WatIWant) Run() {
	appConfig := config.GetConfiguration()

	config.InitializeDatabase()

	fmt.Println("watIwant service listening at port: " + appConfig.Environment.Port)
	errListen := http.ListenAndServe(":"+appConfig.Environment.Port, apiserver.Handlers())

	if errListen != nil {
		log.Fatal(appConfig.Environment.Port, errListen)
	}
}

func main() {
	watIwant := new(WatIWant)
	watIwant.Run()
}
