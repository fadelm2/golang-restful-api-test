package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang-restful-api-technical-test/internal/delivery/http/middleware"
	"golang-restful-api-technical-test/internal/model"
	"golang-restful-api-technical-test/internal/usecase"
	"math"
)

type CreditcardController struct {
	UseCase *usecase.CreditcardUseCase
	Log     *logrus.Logger
}

func NewCreditcardController(useCase *usecase.CreditcardUseCase, log *logrus.Logger) *CreditcardController {
	return &CreditcardController{
		UseCase: useCase,
		Log:     log,
	}
}
func (c *CreditcardController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateCreditcardRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.UserId = auth.ID

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating Creditcard")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.CreditcardResponse]{Data: response})
}

func (c *CreditcardController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.SearchCreditcardRequest{
		UserId: auth.ID,
		Name:   ctx.Query("name", ""),
		Number: ctx.Query("email", ""),
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
	}
	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching Creditcards")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.CreditcardResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *CreditcardController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetCreditcardRequest{
		UserId: auth.ID,
		ID:     ctx.Params("CreditcardId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting Creditcard")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.CreditcardResponse]{Data: response})
}

func (c *CreditcardController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateCreditcardRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.UserId = auth.ID
	request.ID = ctx.Params("CreditcardId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating Creditcard")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreditcardResponse]{Data: response})

}

func (c *CreditcardController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	CreditcardId := ctx.Params("CreditcardId")

	request := &model.DeleteCreditcardRequest{
		UserId: auth.ID,
		ID:     CreditcardId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting Creditcard")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
