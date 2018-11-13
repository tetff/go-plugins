package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tothmate90/go-plugins/handlers"
	"github.com/tothmate90/go-plugins/models"
)

func main() {
	var code string
	if len(os.Args) == 2 {
		code = os.Args[1]
	}

	req := models.Request{
		Name:     "John Doe",
		Username: "JDcoke",
		Email:    "john.doe@coke.org",
		Password: "iwonttellyou",
		Status:   "Active",
		Role:     "Admin",
	}

	res, err := handlers.Run(code, req)
	if err != nil {
		panic(err)
	}

	json, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
