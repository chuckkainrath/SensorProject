package middleware

import (
	"SensorProject/middleware/errors"
	"SensorProject/util"
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func BindRequestBody(next http.Handler, jsonStruct interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonStruct := reflect.New(reflect.TypeOf(jsonStruct))
		jsonAny := jsonStruct.Interface()
		err := json.NewDecoder(r.Body).Decode(jsonAny)
		defer r.Body.Close()

		ctx := r.Context()
		var req *http.Request
		if err != nil {
			appErr := errors.NewValidationError("Unable to bind request body")
			writeResponse(w, appErr.Code, nil, &appErr.Message)
			return
		} else {
			req = r.WithContext(context.WithValue(ctx, InputBodyKey, jsonAny))
		}
		*r = *req

		next.ServeHTTP(w, r)
	})
}

func BindRequestParams(next http.Handler, paramStruct interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		imap := make(map[string]interface{})
		for k, v := range vars {
			imap[k] = v
		}
		paramStruct := reflect.New(reflect.TypeOf(paramStruct))

		paramAny := paramStruct.Interface()
		err := util.Decode(imap, paramAny)
		ctx := r.Context()
		var req *http.Request
		if err != nil {
			appErr := errors.NewValidationError("Unable to bind params")
			writeResponse(w, appErr.Code, nil, &appErr.Message)
			return
		} else {
			req = r.WithContext(context.WithValue(ctx, InputParamsKey, paramAny))
		}
		*r = *req

		next.ServeHTTP(w, r)
	})
}

func GetRequestBody(r *http.Request) interface{} {
	return r.Context().Value(InputBodyKey)
}

func GetRequestParams(r *http.Request) interface{} {
	return r.Context().Value(InputParamsKey)
}
