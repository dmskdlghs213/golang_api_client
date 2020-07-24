package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func TestClient() (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequest("GET", "http://localhost:8080/test", nil)
	if err != nil {
		return "", errors.Wrap(err, "request error")
	}

	req.Header.Set("Authorization", "testtesttesttest")

	rsp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", errors.Wrap(err, "response error")
	}

	defer func() error {
		err = rsp.Body.Close()
		if err != nil {
			return err
		}
		return nil
	}()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", errors.Wrap(err, "response body reading error")
	}

	fmt.Println(string(body))

	return string(body), nil
}
