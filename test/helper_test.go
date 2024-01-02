package test

import (
	"github.com/google/uuid"
	"golang-restful-api-technical-test/internal/entity"
	"strconv"
)

func ClearAll() {
	ClearCreditcards()
	ClearUsers()
}

func ClearUsers() {
	err := db.Where("id is not null").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func ClearCreditcards() {
	err := db.Where("id is not null").Delete(&entity.Creditcard{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func CreateCreditCards(user *entity.User, total int) {
	for i := 0; i < total; i++ {
		contact := entity.Creditcard{
			ID:     uuid.NewString(),
			Type:   "visa",
			Name:   "visa" + strconv.Itoa(i),
			Number: "12222222" + strconv.Itoa(i),
			Cvv:    "08000" + strconv.Itoa(i),
			UserId: user.ID,
		}
		err := db.Create(&contact).Error
		if err != nil {
			log.Fatalf("Failed create contact data :%+v", err)
		}
	}
}
