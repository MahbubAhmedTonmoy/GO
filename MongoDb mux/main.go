package main

import (
	"fmt"
	router "go-rest-api-db/Router"

	//service "go-rest-api-db/Service"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	// conf, err := service.NewConfig("./service/config.yaml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	r := router.Router()
	fmt.Println("Server is getting started...")
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.Database.Port), r))
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at port 8000 ...")
}
