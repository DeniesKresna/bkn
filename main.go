package main

import (
	"github.com/DeniesKresna/bkn/app"
	"github.com/DeniesKresna/bkn/config"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/joho/godotenv"
)

type Stock struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability int     `json:"availability"`
	IsActive     bool    `json:"is_active"`
}

var inventory = map[string]Stock{}

func main() {
	err := godotenv.Load()
	if err != nil {
		utlog.Error("Error loading .env file")
	}

	conf := config.New()

	err = conf.InitDB()
	if err != nil {
		utlog.Errorf("error while connecting DB. %+v", err)
		return
	}

	conf.InitNewValidator()
	if conf.Validator == nil {
		utlog.Errorf("Validator not found. %+v", err)
		return
	}

	conf.InitNewXenditObject()
	if conf.Xendit == nil {
		utlog.Errorf("Xendit Object not found. %+v", err)
		return
	}

	// start Http server
	app := app.InitApp(conf)
	err = app.GateOpen()
	if err != nil {
		utlog.Errorf("error while open gate. %+v", err)
		return
	}
}
