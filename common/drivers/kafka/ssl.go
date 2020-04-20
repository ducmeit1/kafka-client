package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// NewTLS - New TLS for Kafka client
func NewTLS(tlsClientCert, tlsClientKey, tlsClientCA string, skipVerify bool) (*tls.Config, error) {
	//Load client certificate
	if tlsClientCert == "" || tlsClientKey == "" {
		return nil, fmt.Errorf("client cert or client key must not be empty")
	}

	cert, err := tls.LoadX509KeyPair(tlsClientCert, tlsClientKey)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	//Load CA certificate
	if tlsClientCA != "" {
		caCert, err := ioutil.ReadFile(tlsClientCA)
		if err != nil {
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool
	}

	tlsConfig.BuildNameToCertificate()
	tlsConfig.InsecureSkipVerify = skipVerify

	return tlsConfig, nil
}
