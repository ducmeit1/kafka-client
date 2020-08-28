package app

import (
	"github.com/1infras/go-kit/api"
	"github.com/1infras/go-kit/driver/kafka"
	"github.com/1infras/go-kit/logger"
	"github.com/1infras/go-kit/transport"
	"github.com/Shopify/sarama"
	"github.com/ducmeit1/kafka-client/cmd/app/handlers"
	"net/http"
)

var syncProducer sarama.SyncProducer

//NewServer Kafka Client API
func NewServer(name, pathPrefix string) (*api.Server, error) {
	//Setup new server
	s := api.NewServer(name, onClose)

	producer, err := kafka.CreateSyncProducerFromDefaultConnection()
	if err != nil {
		return nil, err
	}

	syncProducer = producer

	//Add route to server
	s.AddRouter(initRoutes(pathPrefix))

	return s, nil
}

func initRoutes(pathPrefix string) transport.Transport {
	return transport.Transport{
		PathPrefix: pathPrefix,
		Routes: []transport.Route{
			{
				Path:    "/{topic}",
				Method:  http.MethodPost,
				Handler: &handlers.ProducerHandler{Producer: syncProducer},
			},
		},
	}
}

func onClose() {
	if err := syncProducer.Close(); err != nil {
		logger.Errorf("close connection from kafka has error: %v", err)
	}
}
