package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chuckkainrath/SensorProject/middleware/errors"
)

type response struct {
	Err  string      `json:"error"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func WriteResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		ctx := r.Context()
		errVal := ctx.Value(ErrorKey)
		if errVal != nil {
			err := errVal.(*errors.AppError)
			writeResponse(w, err.Code, nil, &err.Message)
			return
		}

		data := ctx.Value(OutputDataKey)
		if data != nil {
			writeResponse(w, http.StatusOK, &data, nil)
			return
		}

		writeResponse(w, http.StatusOK, nil, nil)
	})
}

func writeResponse(w http.ResponseWriter, code int, data *interface{}, err *string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	//if this doesn't work we want the application to shutdown and show you the error message with a panic

	res := response{Code: code}
	if data != nil {
		res.Data = *data
	}
	if err != nil {
		res.Err = *err
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func AddResultToContext(r *http.Request, data interface{}, dataType ContextKey) {
	ctx := r.Context()
	req := r.WithContext(context.WithValue(ctx, dataType, data))
	*r = *req
}
