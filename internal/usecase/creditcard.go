package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang-restful-api-technical-test/internal/entity"
	"golang-restful-api-technical-test/internal/gateway/messaging"
	"golang-restful-api-technical-test/internal/model"
	"golang-restful-api-technical-test/internal/model/converter"
	"golang-restful-api-technical-test/internal/repository"
	"gorm.io/gorm"
)

type CreditcardUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	CreditcardRepository *repository.CreditcardRepository
	CreditcardProducer   *messaging.CreditcardProducer
}

func NewCreditcardUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, CreditcardRepository *repository.CreditcardRepository, CreditcardProducer *messaging.CreditcardProducer) *CreditcardUseCase {
	return &CreditcardUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		CreditcardRepository: CreditcardRepository,
		CreditcardProducer:   CreditcardProducer,
	}
}

func (c *CreditcardUseCase) Create(ctx context.Context, request *model.CreateCreditcardRequest) (*model.CreditcardResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}
	Creditcard := &entity.Creditcard{
		ID:      uuid.New().String(),
		Type:    request.Type,
		Name:    request.Name,
		Number:  request.Number,
		Expired: request.Expired,
		Cvv:     request.Cvv,
		UserId:  request.UserId,
	}

	if err := c.CreditcardRepository.Create(tx, Creditcard); err != nil {
		c.Log.WithError(err).Error("error creating Creditcard")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating Creditcard")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.CreditcardToEvent(Creditcard)
	if err := c.CreditcardProducer.Send(event); err != nil {
		c.Log.WithError(err).Error("error Publishing Creditcard")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CreditcardToResponse(Creditcard), nil
}

func (c *CreditcardUseCase) Update(ctx context.Context, request *model.UpdateCreditcardRequest) (*model.CreditcardResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	Creditcard := new(entity.Creditcard)
	if err := c.CreditcardRepository.FindByIdAndUserId(tx, Creditcard, request.ID, request.UserId); err != nil {
		c.Log.WithError(err).Error("error getting Creditcard")
		return nil, fiber.ErrNotFound
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	Creditcard.Name = request.Name
	Creditcard.Number = request.Number
	Creditcard.Expired = request.Expired
	Creditcard.Cvv = request.Cvv

	if err := c.CreditcardRepository.Update(tx, Creditcard); err != nil {
		c.Log.WithError(err).Error("error updating Creditcard")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Error updating Creditcard")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.CreditcardToEvent(Creditcard)
	if err := c.CreditcardProducer.Send(event); err != nil {
		c.Log.WithError(err).Error("error publishing Creditcard")
		return nil, fiber.ErrInternalServerError
	}
	return converter.CreditcardToResponse(Creditcard), nil
}

func (c *CreditcardUseCase) Get(ctx context.Context, request *model.GetCreditcardRequest) (*model.CreditcardResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	Creditcard := new(entity.Creditcard)
	if err := c.CreditcardRepository.FindByIdAndUserId(tx, Creditcard, request.ID, request.UserId); err != nil {
		c.Log.WithError(err).Error("error getting Creditcard")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Error getting Creditcard")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CreditcardToResponse(Creditcard), nil
}

func (c *CreditcardUseCase) Delete(ctx context.Context, request *model.DeleteCreditcardRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	Creditcard := new(entity.Creditcard)
	if err := c.CreditcardRepository.FindByIdAndUserId(tx, Creditcard, request.ID, request.UserId); err != nil {
		c.Log.WithError(err).Error("Error getting Creditcard")
		return fiber.ErrNotFound
	}

	if err := c.CreditcardRepository.Delete(tx, Creditcard); err != nil {
		c.Log.WithError(err).Error("error deleting Creditcard")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Error deleting Creditcard")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *CreditcardUseCase) Search(ctx context.Context, request *model.SearchCreditcardRequest) ([]model.CreditcardResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	Creditcards, total, err := c.CreditcardRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting Creditcards")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting Creditcards")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.CreditcardResponse, len(Creditcards))
	for i, Creditcard := range Creditcards {
		responses[i] = *converter.CreditcardToResponse(&Creditcard)
	}
	return responses, total, nil
}
