package messaging

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"golang-restful-api-technical-test/internal/model"
)

type CreditcardConsumer struct {
	Log *logrus.Logger
}

func NewCreditcardConsumer(log *logrus.Logger) *CreditcardConsumer {
	return &CreditcardConsumer{
		Log: log,
	}
}

func (c CreditcardConsumer) Consume(message *kafka.Message) error {
	CreditcardEvent := new(model.CreditcardEvent)
	if err := json.Unmarshal(message.Value, CreditcardEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling Creditcard event")
		return err
	}

	//TODO process event
	c.Log.Infof("Received topic users with event : %v fron partition %d", CreditcardEvent, message.TopicPartition.Partition)
	return nil
}
