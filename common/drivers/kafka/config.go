package kafka

import "crypto/tls"

type Config struct {
	Servers   []string
	TLSConfig *tls.Config
}
