package conf

import (
	"go-echo/internal/models"
)

func Get(path string) (config models.Conf) {

	config.Host = "0.0.0.0"
	config.Port = "8090"
	return config
}
