package req

import (
	"net/http"
	"temporary/pkg/res"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)

	if err != nil {
		res.WriteJSON(*w, ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return nil, err
	}

	err = IsValid(body)

	if err != nil {
		res.WriteJSON(*w, ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return nil, err
	}

	return &body, nil
}