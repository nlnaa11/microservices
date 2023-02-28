package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/pkg/errors"
)

type Validator interface {
	Validate() error
}

type Wrapper[Req any, Res any] struct {
	urlPath string
}

func New[Req any, Res any](path string) *Wrapper[Req, Res] {
	return &Wrapper[Req, Res]{
		urlPath: path,
	}
}

func (w *Wrapper[Req, Res]) Service(ctx context.Context, req Req) (Res, error) {
	var response Res

	rawJson, err := json.Marshal(req)
	if err != nil {
		return response, errors.Wrap(err, "marshaling json")
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, w.urlPath, bytes.NewBuffer(rawJson))
	if err != nil {
		return response, errors.Wrap(err, "creating http request")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return response, errors.Wrap(err, "calling http")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return response, fmt.Errorf("wrong status code: %d", httpResponse.StatusCode)
	}

	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return response, errors.Wrap(err, "decoding json")
	}

	return response, nil
}
