package handlers

import (
	"fmt"
	"plugin"

	"github.com/tothmate90/go-plugins/models"
)

type Handler interface {
	Magic(models.Request) (models.Response, error)
}

func Run(code string, req models.Request) (models.Response, error) {
	var res models.Response

	var mod string
	switch code {
	case "regular":
		mod = "./plugins/regular.so"
		break
	case "encrypt":
		mod = "./plugins/encrypt.so"
		break
	default:
		return res, fmt.Errorf("Invalid input")
	}

	plug, err := plugin.Open(mod)
	if err != nil {
		return res, err
	}

	symbol, err := plug.Lookup("Handler")
	if err != nil {
		return res, err
	}

	handler, ok := symbol.(Handler)
	if !ok {
		return res, fmt.Errorf("Unexpected type from module")
	}

	res, err = handler.Magic(req)
	if err != nil {
		return res, err
	}

	return res, nil
}
