package main

import (
	"github.com/ducmeit1/kafka-client/cmd/producer/app"
	"github.com/ducmeit1/kafka-client/common"
)

func main() {
	kafkaProducer := common.KafkaProducer

	s, err := app.NewServer(kafkaProducer.Name, kafkaProducer.PathPrefix)
	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
