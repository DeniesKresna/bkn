package config

import (
	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/gohelper/utstring"
)

func New() *Config {
	var cfg Config

	cfg.Service = &models.Service{
		Name:      utstring.GetEnv(models.AppNameENV, "main-api-service"),
		Version:   utstring.GetEnv(models.AppVersionENV),
		Host:      utstring.GetEnv(models.AppHostENV, "0.0.0.0"),
		Port:      utstring.GetEnv(models.AppPortENV, "9010"),
		Env:       utstring.GetEnv(models.AppENV, "dev"),
		Namespace: utstring.GetEnv(models.AppNamespaceENV),
		WebURL:    utstring.GetEnv(models.AppWebURL, "https://dev.jobhun.id"),
	}

	utlog.Info("Init App Service Done")

	return &cfg
}
