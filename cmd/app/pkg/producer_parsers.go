package pkg

import (
	"io/ioutil"
	"net/http"
)

func ProducerRequestBodyParser(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return "", err
	}

	return string(body), nil
}
