package app

import (
	"github.com/DeniesKresna/bkn/config"
	"github.com/DeniesKresna/bkn/service/delivery/gate"
	"github.com/DeniesKresna/bkn/service/repository/mail"
	"github.com/DeniesKresna/bkn/service/repository/outgate"
	"github.com/DeniesKresna/bkn/service/repository/sql"
	"github.com/DeniesKresna/bkn/service/repository/storage"
	"github.com/DeniesKresna/bkn/service/usecase"
)

type Application struct {
	Conf    *config.Config
	AppGate *gate.Gate
}

func InitApp(conf *config.Config) *Application {
	userRepo := sql.InitUserSqlRepository(conf.DB, conf.Service, conf.Q)
	expertRepo := sql.InitExpertSqlRepository(conf.DB, conf.Service, conf.Q)
	courseRepo := sql.InitCourseSqlRepository(conf.DB, conf.Service, conf.Q)
	paymentRepo := sql.InitPaymentSqlRepository(conf.DB, conf.Service, conf.Q)
	paymentOutgate := outgate.InitPaymentOutgateRepository(conf.Xendit)
	storageRepo := storage.InitS3Storage()
	mailRepo, _ := mail.InitGmail()

	authUsecase := usecase.InitAuthUsecase(conf.DB, userRepo)
	userUsecase := usecase.InitUserUsecase(conf.DB, conf.Service, authUsecase, userRepo, storageRepo, mailRepo)
	expertUsecase := usecase.InitExpertUsecase(conf.DB, expertRepo, mailRepo, userUsecase, storageRepo, authUsecase)
	courseUsecase := usecase.InitCourseUsecase(conf.DB, courseRepo, mailRepo, userUsecase, storageRepo, authUsecase)
	paymentUsecase := usecase.InitPaymentUsecase(conf.DB, paymentRepo, paymentOutgate, courseUsecase, mailRepo, userUsecase, storageRepo, authUsecase, conf)

	appGate := gate.InitGate(conf.Validator, conf.MessagerLogger, userUsecase, authUsecase, expertUsecase, courseUsecase, paymentUsecase)

	return &Application{
		Conf:    conf,
		AppGate: appGate,
	}
}
