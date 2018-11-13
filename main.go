package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/tothmate90/go-plugins/handlers"
	"github.com/tothmate90/go-plugins/models"
)

func main() {
	path := flag.String("path", "./plugins/regular.so", "regular for simple json, and encrypt for encrypted")
	flag.Parse()

	req := models.Request{
		Name:     "John Doe",
		Username: "JDcoke",
		Email:    "john.doe@coke.org",
		Password: "iwonttellyou",
		Status:   "Active",
		Role:     "Admin",
	}

	res, err := handlers.Run(path, req)
	if err != nil {
		panic(err)
	}

	json, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
