package cmd

type App struct {
	Name string
	PathPrefix string
}

var (
	KafkaClient = App{Name: "Kafka Client", PathPrefix:  "/v1/produce"}
)
