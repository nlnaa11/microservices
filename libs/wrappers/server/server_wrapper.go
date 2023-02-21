package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Validator interface {
	Validate() error
}

type Wrapper[Req Validator, Res any] struct {
	fn func(ctx context.Context, req Req) (Res, error)
}

func New[Req Validator, Res any](fn func(ctx context.Context, req Req) (Res, error)) *Wrapper[Req, Res] {
	return &Wrapper[Req, Res]{
		fn: fn,
	}
}

func (w *Wrapper[Req, Res]) ServeHTTP(resWriter http.ResponseWriter, httpReq *http.Request) {
	ctx := httpReq.Context()

	var request Req
	err := json.NewDecoder(httpReq.Body).Decode(&request)
	if err != nil {
		processsError(resWriter, http.StatusBadRequest, ErrDecodingJson, err)
		return
	}

	err = request.Validate()
	if err != nil {
		processsError(resWriter, http.StatusBadRequest, ErrValidatingRequest, err)
		return
	}

	response, err := w.fn(ctx, request)
	if err != nil {
		processsError(resWriter, http.StatusInternalServerError, ErrRunningHandler, err)
		return
	}

	rawJSON, err := json.Marshal(response)
	if err != nil {
		processsError(resWriter, http.StatusInternalServerError, ErrEncodingJson, err)
		return
	}

	resWriter.Header().Add("Content-Type", "application/json")
	_, _ = resWriter.Write(rawJSON)
}

func processsError(resWriter http.ResponseWriter, status int, text string, err error) {
	buf := bytes.NewBufferString(text)
	buf.WriteString(": ")
	buf.WriteString(err.Error())
	buf.WriteByte('\n')

	resWriter.WriteHeader(status)
	resWriter.Write(buf.Bytes())
}
