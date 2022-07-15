package util

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/chuckkainrath/SensorProject/middleware/errors"
	"github.com/mitchellh/mapstructure"
)

func Decode(input map[string]interface{}, result interface{}) *errors.AppError {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc(),
			toUintHookFunc()),
		Result: result,
	})
	if err != nil {
		fmt.Printf("Err: %v\n", err)

		return errors.NewValidationError("Unable to bind params")
	}

	if err := decoder.Decode(input); err != nil {
		fmt.Printf("Err: %v\n", err)

		return errors.NewValidationError("Unable to bind params")
	}
	return nil
}

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		layout := "2006-01-02T15:04:05"
		return time.Parse(layout, data.(string))
	}
}

func toUintHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t.Kind() != reflect.Uint {
			return data, nil
		}

		res, err := strconv.ParseUint(data.(string), 10, 32)
		if err != nil {
			return nil, err
		}
		return uint(res), nil
	}
}
