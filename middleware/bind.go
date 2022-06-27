package middleware

import (
	"SensorProject/middleware/errors"
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

func BindModel(next http.Handler, jsonStruct interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonStruct := reflect.New(reflect.TypeOf(jsonStruct))
		err := json.NewDecoder(r.Body).Decode(jsonStruct.Interface())
		defer r.Body.Close()

		ctx := r.Context()
		var req *http.Request
		if err != nil {
			appErr := errors.NewValidationError("Unable to bind request body")
			writeResponse(w, appErr.Code, nil, &appErr.Message)
			return
		} else {
			req = r.WithContext(context.WithValue(ctx, InputBodyKey, jsonStruct))
		}
		*r = *req

		next.ServeHTTP(w, r)
	})
}

func BindParams(next http.Handler, paramStruct interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramStruct := reflect.New(reflect.TypeOf(paramStruct))
		err := mapstructure.Decode(vars, paramStruct.Interface())

		ctx := r.Context()
		var req *http.Request
		if err != nil {
			appErr := errors.NewValidationError("Unable to bind params")
			writeResponse(w, appErr.Code, nil, &appErr.Message)
			return
		} else {
			req = r.WithContext(context.WithValue(ctx, InputParamsKey, paramStruct))
		}
		*r = *req

		next.ServeHTTP(w, r)
	})
}

func GetBody(r *http.Request) interface{} {
	return r.Context().Value(InputBodyKey)
}

func GetParams(r *http.Request) interface{} {
	return r.Context().Value(InputParamsKey)
}
