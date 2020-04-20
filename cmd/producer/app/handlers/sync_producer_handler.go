package handlers

import (
	"github.com/Shopify/sarama"
	"github.com/ducmeit1/kafka-client/cmd/producer/app/parsers"
	"github.com/ducmeit1/kafka-client/common/transports"
	"net/http"
)

type SyncProducerHandler struct {
	Producer sarama.SyncProducer
}

func (p *SyncProducerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pr, err := parsers.ProducerParsers(r)
	if err != nil {
		transports.BadRequest(w, map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	m := &sarama.ProducerMessage{
		Topic: pr.Topic,
		Key:   nil,
		Value: sarama.StringEncoder(pr.Message),
	}

	partition, offset, err := p.Producer.SendMessage(m)
	if err != nil {
		transports.BadRequest(w, map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	transports.OK(w, map[string]interface{}{
		"Partition": partition,
		"Offset":    offset,
	})
}
