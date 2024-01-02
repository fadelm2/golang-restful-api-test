package route

import (
	"github.com/gofiber/fiber/v2"
	"golang-restful-api-technical-test/internal/delivery/http"
)

type RouteConfig struct {
	App                  *fiber.App
	UserController       *http.UserController
	CreditcardController *http.CreditcardController
	AuthMiddleWare       fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_Login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleWare)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("api/users/_current", c.UserController.Current)

	c.App.Get("/api/creditcards", c.CreditcardController.List)
	c.App.Post("/api/creditcards", c.CreditcardController.Create)
	c.App.Put("/api/creditcards/:CreditcardId", c.CreditcardController.Update)
	c.App.Get("/api/creditcards/:CreditcardId", c.CreditcardController.Get)
	c.App.Delete("/api/creditcards/:CreditcardId", c.CreditcardController.Delete)
}
