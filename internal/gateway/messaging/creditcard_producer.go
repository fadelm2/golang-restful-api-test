package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"golang-restful-api-technical-test/internal/model"
)

type CreditcardProducer struct {
	Producer[*model.CreditcardEvent]
}

func NewCreditcardProducer(producer *kafka.Producer, log *logrus.Logger) *CreditcardProducer {
	return &CreditcardProducer{
		Producer: Producer[*model.CreditcardEvent]{
			Producer: producer,
			Topic:    "creditcard",
			Log:      log,
		},
	}
}
