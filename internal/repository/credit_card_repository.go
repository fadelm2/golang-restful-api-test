package repository

import (
	"golang-restful-api-technical-test/internal/entity"
	"golang-restful-api-technical-test/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreditcardRepository struct {
	Repository[entity.Creditcard]
	Log *logrus.Logger
}

func NewCreditcardRepository(log *logrus.Logger) *CreditcardRepository {
	return &CreditcardRepository{
		Log: log,
	}
}
func (r *CreditcardRepository) FindByIdAndUserId(db *gorm.DB, Creditcard *entity.Creditcard, id string, userId string) error {
	return db.Where("id = ? AND user_id = ?", id, userId).Take(Creditcard).Error
}

func (r *CreditcardRepository) Search(db *gorm.DB, request *model.SearchCreditcardRequest) ([]entity.Creditcard, int64, error) {
	var Creditcards []entity.Creditcard

	if err := db.Scopes(r.FilterCreditcard(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&Creditcards).Error; err != nil {
		return nil, 0, err
	}
	var total int64 = 0
	if err := db.Model(&entity.Creditcard{}).Scopes(r.FilterCreditcard(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return Creditcards, total, nil
}

func (r *CreditcardRepository) FilterCreditcard(request *model.SearchCreditcardRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", request.UserId)

		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ? ", name)
		}

		if number := request.Number; number != "" {
			number = "%" + number + "%"
			tx = tx.Where("number LIKE ?", number)
		}

		return tx

	}
}
