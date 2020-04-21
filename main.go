package main

import (
	"github.com/ducmeit1/kafka-client/cmd"
	"github.com/ducmeit1/kafka-client/common"
	log "github.com/sirupsen/logrus"
)

func main() {
	kafkaProducer := common.KafkaProducer

	s, err := cmd.NewServer(kafkaProducer.Name, kafkaProducer.PathPrefix)
	if err != nil {
		log.Errorf("Unable to new http server: %v", err)
		return
	}

	err = s.Run()
	if err != nil {
		log.Error("Run server has failed: %v", err)
	}
}
