package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang-restful-api-technical-test/internal/delivery/http"
	"golang-restful-api-technical-test/internal/delivery/http/middleware"
	"golang-restful-api-technical-test/internal/delivery/http/route"
	"golang-restful-api-technical-test/internal/gateway/messaging"
	"golang-restful-api-technical-test/internal/repository"
	"golang-restful-api-technical-test/internal/usecase"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Producer *kafka.Producer
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	creditCardRepository := repository.NewCreditcardRepository(config.Log)

	//setup producer
	userProducer := messaging.NewUserProducer(config.Producer, config.Log)
	creditCardProducer := messaging.NewCreditcardProducer(config.Producer, config.Log)

	//setup use cases
	userUseCase := usecase.NewUserCase(config.DB, config.Log, config.Validate, userRepository, userProducer)
	CreditcardUseCase := usecase.NewCreditcardUseCase(config.DB, config.Log, config.Validate, creditCardRepository, creditCardProducer)

	//setup controller
	userController := http.NewUserController(config.Log, userUseCase)
	CreditcardController := http.NewCreditcardController(CreditcardUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:                  config.App,
		UserController:       userController,
		CreditcardController: CreditcardController,
		AuthMiddleWare:       authMiddleware,
	}
	routeConfig.Setup()
}
