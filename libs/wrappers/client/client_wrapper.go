package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/pkg/errors"
)

type RequestMaker interface {
	MakeRequest() (any, error)
}

type ResponseDataMaker interface {
	MakeResponseData() (any, error)
}

type Wrapper[ReqData RequestMaker, Res ResponseDataMaker] struct {
	path string
}

func New[ReqData RequestMaker, Res ResponseDataMaker](path string) *Wrapper[ReqData, Res] {
	return &Wrapper[ReqData, Res]{
		path: path,
	}
}

func (w *Wrapper[ReqData, Res]) Service(ctx context.Context, reqData ReqData, path string) (any, error) {
	var responseData any

	request, err := reqData.MakeRequest()
	if err != nil {
		return responseData, errors.Wrap(err, "converting data to request")
	}

	rawJson, err := json.Marshal(request)
	if err != nil {
		return responseData, errors.Wrap(err, "marshaling json")
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewBuffer(rawJson))
	if err != nil {
		return responseData, errors.Wrap(err, "creating http request")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return responseData, errors.Wrap(err, "calling http")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return responseData, fmt.Errorf("wrong status code: %d", httpResponse.StatusCode)
	}

	var res Res
	err = json.NewDecoder(httpResponse.Body).Decode(&res)
	if err != nil {
		return responseData, errors.Wrap(err, "decoding json")
	}

	responseData, err = res.MakeResponseData()
	if err != nil {
		return responseData, errors.Wrap(err, "converting response to data")
	}

	return responseData, nil
}
