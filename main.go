package main

import (
	"github.com/ducmeit1/kafka-client/cmd"
	"github.com/ducmeit1/kafka-client/cmd/app"
)

func main() {
	api := cmd.KafkaClient
	//Setup new server
	s, err := app.NewServer(api.Name, api.PathPrefix)
	if err != nil {
		panic(err)
	}
	//Run server
	s.Run()
}
