package config

import (
	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/myqgen2/qgen"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Service        *models.Service
	DB             *sdb.DBInstance
	Q              *qgen.Obj
	Validator      *validator.Validate
	MessagerLogger IMessagerLogger
	Xendit         *Xendit
}
