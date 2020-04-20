package common

type App struct {
	Name       string
	PathPrefix string
}

var (
	KafkaProducer = App{Name: "Kafka - Producer", PathPrefix: "/v1/producer"}
)
