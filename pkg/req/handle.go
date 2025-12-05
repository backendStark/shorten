package req

import (
	"net/http"
	"shorten/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := Decode[T](req.Body)

	if err != nil {
		res.JSON(*w, 402, err.Error())
		return nil, err
	}

	if err := IsValid(body); err != nil {
		res.JSON(*w, 400, err.Error())
		return nil, err
	}

	return &body, err
}
