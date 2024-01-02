package converter

import (
	"golang-restful-api-technical-test/internal/entity"
	"golang-restful-api-technical-test/internal/model"
)

func CreditcardToResponse(Creditcard *entity.Creditcard) *model.CreditcardResponse {
	return &model.CreditcardResponse{
		ID:        Creditcard.ID,
		Type:      Creditcard.Type,
		Name:      Creditcard.Name,
		Number:    Creditcard.Number,
		Expired:   Creditcard.Expired,
		Cvv:       Creditcard.Cvv,
		CreatedAt: Creditcard.CreatedAt,
		UpdatedAt: Creditcard.UpdatedAt,
	}
}

func CreditcardToEvent(Creditcard *entity.Creditcard) *model.CreditcardEvent {
	return &model.CreditcardEvent{
		ID:        Creditcard.ID,
		UserID:    Creditcard.UserId,
		Type:      Creditcard.Type,
		Name:      Creditcard.Name,
		CreatedAt: Creditcard.CreatedAt,
		UpdatedAt: Creditcard.UpdatedAt,
	}
}
