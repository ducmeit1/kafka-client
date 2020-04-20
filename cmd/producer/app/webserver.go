package app

import (
	"github.com/Shopify/sarama"
	"github.com/ducmeit1/kafka-client/cmd/producer/app/handlers"
	kafka_driver "github.com/ducmeit1/kafka-client/common/drivers/kafka"
	"github.com/ducmeit1/kafka-client/common/kits"
	"github.com/ducmeit1/kafka-client/common/transports"
	"github.com/prometheus/common/log"
)

var producer sarama.SyncProducer

func NewServer(name, prefix string) (*kits.Server, error) {
	p, err := kafka_driver.CreateSyncProducerFromDefaultConfig()
	if err != nil {
		return nil, err
	}

	producer = p

	s, err := kits.NewServer(name, OnClose)
	if err != nil {
		return nil, err
	}

	s.AddRoutes(InitRoutes(prefix))

	return s, nil
}

func OnClose() {
	err := producer.Close()
	if err != nil {
		log.Error(err.Error())
	}
}

func InitRoutes(prefix string) transports.Transport {
	return transports.Transport{
		PathPrefix: prefix,
		Routes: []transports.Route{
			{
				Path: "/",
				Handler: &handlers.SyncProducerHandler{
					Producer: producer,
				},
			},
		},
	}
}
