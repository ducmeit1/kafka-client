# KAFKA CLIENT

## Introduction

A simple client helps to produce message to any topic kafka in a random partition number.

## API Usages
```
POST: /v1/producer
{
  "topic": "topic_name",
  "message": "message"
}
```

## Setups with Kafka

### Kafka with TLS

- Add your pem Certificate and Private Key path to the `config.toml`
- If you are having a PKCS12 certificate or JKS, please export it to pem. For example:
```bash
# EXPORT TO PKCS12 FROM JKS (JAVA KEYSTORE)
keytool -importkeystore -srckeystore client.jks -destkeystore client.p12 -srcstoretype jks -deststoretype pkcs12

# EXPORT TO PEM FROM P12
openssl pkcs12 -in client.p12 -out client.pem

# EXPORT CERTIFICATE
openssl x509 -in client.pem -out client_cert.crt

# EXPORT PRIVATE KEY
openssl rsa -in client.pem -out client_key.key
```

## Docker
```bash
docker pull ducmeit1/kafka-client:latest 
```

## Kubernetes

- Configure at `kubernetes/deployment.yml`
- Create namespace: `kafka`: `kubectl create namespace kafka`
- Apply scripts: `kubectl apply -f deployment.yml`