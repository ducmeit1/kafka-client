package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ducmeit1/kafka-client/common/utils/io"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewKafkaConfig(config *Config) *sarama.Config {
	c := sarama.NewConfig()
	c.Producer.RequiredAcks = sarama.WaitForAll
	c.Producer.Return.Successes = true
	c.Producer.Return.Errors = true
	if config.TLSConfig != nil {
		c.Net.TLS.Enable = true
		c.Net.TLS.Config = config.TLSConfig
		for k := range config.TLSConfig.NameToCertificate {
			log.Infof("Connecting to Kafka with SSL Client Name: %v", k)
		}
	}
	c.Producer.Partitioner = sarama.NewRandomPartitioner
	return c
}

func GetClientFilePathByViper(key string) (string, error) {
	return io.GetAbsolutelyFilePath(viper.GetString(key))
}

func GetDefaultConfig() (*Config, error) {
	servers := viper.GetStringSlice("kafka.servers")
	if len(servers) == 0 {
		return nil, fmt.Errorf("kafka servers must be defined")
	}

	c := &Config{
		Servers: servers,
	}

	if viper.GetBool("kafka.tls") {
		tlsClientCert, err := GetClientFilePathByViper("kafka.tls_client_cert")
		if err != nil {
			return nil, err
		}

		tlsClientKey, err := GetClientFilePathByViper("kafka.tls_client_key")
		if err != nil {
			return nil, err
		}

		tlsClientCA := viper.GetString("kafka.tls_client_ca")
		tlsSkipVerify := viper.GetBool("kafka.tls_skip_verify")

		if tlsClientCA != "" {
			tlsClientCA, err = GetClientFilePathByViper("kafka.tls_client_ca")
			if err != nil {
				return nil, err
			}
		}

		tlsConfig, err := NewTLS(tlsClientCert, tlsClientKey, tlsClientCA, tlsSkipVerify)
		if err != nil {
			return nil, err
		}

		c.TLSConfig = tlsConfig
	}

	return c, nil
}

func CreateAsyncProducer(config *Config) (sarama.AsyncProducer, error) {
	if config == nil {
		return CreateAsyncProducerFromDefaultConfig()
	}

	c := NewKafkaConfig(config)
	return sarama.NewAsyncProducer(config.Servers, c)
}

func CreateSyncProducer(config *Config) (sarama.SyncProducer, error) {
	if config == nil {
		return CreateSyncProducerFromDefaultConfig()
	}

	c := NewKafkaConfig(config)
	return sarama.NewSyncProducer(config.Servers, c)
}

func CreateAsyncProducerFromDefaultConfig() (sarama.AsyncProducer, error) {
	c, err := GetDefaultConfig()
	if err != nil {
		return nil, err
	}
	return CreateAsyncProducer(c)
}

func CreateSyncProducerFromDefaultConfig() (sarama.SyncProducer, error) {
	c, err := GetDefaultConfig()
	if err != nil {
		return nil, err
	}
	return CreateSyncProducer(c)
}
