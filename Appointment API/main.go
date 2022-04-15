package main

import (
	"AppointmentApi/config"
	"AppointmentApi/domain"
	"AppointmentApi/email"
	"AppointmentApi/handlers"
	"AppointmentApi/middleware"
	"AppointmentApi/redisconfig"
	"AppointmentApi/repository"
	"AppointmentApi/route"
	"AppointmentApi/service"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.NewConfig("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// mail config ..
	mailChan := make(chan domain.MailData, 2) // i make a buffered channel of capacity: 2 for 2 mails at once, doctor & patient
	email.ListenForMail(mailChan, conf.Mail.SmtpHost, conf.Mail.SmtpPort, conf.Mail.SenderEmail, conf.Mail.Password)

	// redis cache config ..
	redisClient := redisconfig.InitializeRedisServer(conf.Redis.Address, conf.Redis.Password, conf.Redis.DB)
	context := context.Background()

	// Initializing database ..
	db, _ := repository.InitializeDB(conf.Database.ConnectionString)

	// Initializing repos ..
	docRepo := repository.DoctorRepo.Initialize(db)
	accRepo := repository.AccountRepo.Initialize(db)

	// Initializing services ...
	baseSrv := service.InitializeBaseService(redisClient, &context, mailChan, conf)
	service.DoctorService.Initialize(baseSrv, docRepo)
	service.AccountService.Initialize(baseSrv, accRepo)

	// Initializing handlers ...
	handlers.Base.Initialize(conf)
	handlers.DoctorHandler.Initialize()
	handlers.AccountHandler.Initialize()

	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	router := gin.Default()

	// Initialize middlewares for this app
	middleware.InitializeMiddlewares(router)

	// Initialize routes for this app
	route.InitializeRoutes(router)

	// By default it serves on :8080 unless a PORT environment variable was defined.
	log.Fatal(router.Run()) // router.Run(":3000") for a hard coded port
}
