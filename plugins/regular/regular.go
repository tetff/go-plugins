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

	res.Data.Name = req.Name
	res.Data.Username = req.Username
	res.Data.Email = req.Email
	res.Data.Status = req.Status
	res.Data.Role = req.Role
	res.Data.CreatedAt = time.Time.String(time.Now())

	return res, nil
}

var Handler handler
