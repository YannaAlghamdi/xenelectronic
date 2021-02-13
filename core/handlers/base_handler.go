package handler

import (
	"encoding/json"
	"io"
)

type BaseHandler struct{}

type GenericResponse struct {
	Result bool   `json:"result"`
	Error  string `json:"error,omitempty"`
}

func (handler *BaseHandler) EncodeRequest(request io.Reader, model interface{}) error {
	err := json.NewDecoder(request).Decode(model)
	return err
}

func (handler *BaseHandler) EncodeResponse(response io.Writer, model interface{}) {
	json.NewEncoder(response).Encode(model)
}

func buildGenericError(err error) *GenericResponse {
	return &GenericResponse{
		Result: false,
		Error:  err.Error(),
	}
}

func buildGenericSuccess() *GenericResponse {
	return &GenericResponse{
		Result: true,
	}
}

// func error(res http.ResponseWriter, err *response.Error) {
// 	http.Error(res, err.Code, err.Status)
// }
