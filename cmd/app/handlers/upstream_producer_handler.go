package handlers

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ducmeit1/kafka-client/cmd/app/parsers"
	"github.com/ducmeit1/kafka-client/common/transports"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UpStreamProducerHandler struct {
	Producer sarama.SyncProducer
}

func (p *UpStreamProducerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	if topic == "" {
		transports.BadRequest(w, map[string]interface{}{
			"Error": fmt.Errorf("empty topic"),
		})
		return
	}

	pr, err := parsers.ProducerUpstreamParsers(r)
	if err != nil {
		transports.BadRequest(w, map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	m := &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: sarama.StringEncoder(pr),
	}

	partition, offset, err := p.Producer.SendMessage(m)
	if err != nil {
		transports.BadRequest(w, map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	log.Infof("Sent 1 message to topic: %v with offset: %v, partition: %v", topic, offset, partition)

	transports.OK(w, map[string]bool{
		"success": true,
	})
}
