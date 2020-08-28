package handlers

import (
	"fmt"
	"github.com/1infras/go-kit/logger"
	"github.com/1infras/go-kit/transport"
	"github.com/Shopify/sarama"
	"github.com/ducmeit1/kafka-client/cmd/app/pkg"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//ProducerHandler Kafka Client
type ProducerHandler struct {
	Producer sarama.SyncProducer
}

func (h *ProducerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		vars  = mux.Vars(r)
		topic = vars["topic"]
	)

	if topic == "" {
		transport.BadRequest(w, map[string]interface{}{
			"Error": fmt.Errorf("topic was empty"),
		})
		return
	}

	pr, err := pkg.ProducerRequestBodyParser(r)
	if err != nil {
		logger.Warnf("The request body wasn't parsed, has error: %v", err.Error())
		transport.BadRequest(w, map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	m := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(pr),
	}

	partition, offset, err := h.Producer.SendMessage(m)
	if err != nil {
		transport.BadRequest(w, map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	log.Infof("Sent 1 message to topic: %v with offset: %v, partition: %v", topic, offset, partition)

	transport.OKJson(w, map[string]interface{}{
		"Partition": partition,
		"Offset":    offset,
	})
}
