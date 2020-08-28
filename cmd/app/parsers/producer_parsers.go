package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/ducmeit1/kafka-client/cmd/app/pkg"
	"io/ioutil"
	"log"
	"net/http"
)

func ProducerParsers(r *http.Request) (pkg.ProduceRequest, error) {
	var (
		pr      = pkg.ProduceRequest{}
		decoder = json.NewDecoder(r.Body)
	)

	err := decoder.Decode(&pr)
	if err != nil {
		return pr, err
	}

	if pr.Topic == "" {
		return pr, fmt.Errorf("topic must not be empty")
	}

	if pr.Message == "" {
		return pr, fmt.Errorf("message must not be empty")
	}
	return pr, nil
}

func ProducerUpstreamParsers(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return "", err
	}

	return string(body), nil
}
