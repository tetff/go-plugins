package main

import (
	"time"

	"github.com/tothmate90/go-plugins/models"
)

type handler string

func (h handler) Magic(req models.Request) (models.Response, error) {
	var res models.Response
	res.Status = 200
	res.Message = "success"

	res.Data.Name = encrypt(req.Name)
	res.Data.Username = encrypt(req.Username)
	res.Data.Email = encrypt(req.Email)
	res.Data.Status = encrypt(req.Status)
	res.Data.Role = encrypt(req.Role)
	res.Data.CreatedAt = encrypt(time.Time.String(time.Now()))

	return res, nil
}

func encrypt(str string) string {
	oba := []byte(str)
	var mba []byte
	for _, c := range oba {
		c = c + 1
		mba = append(mba, c)
	}
	return string(mba)
}

var Handler handler
