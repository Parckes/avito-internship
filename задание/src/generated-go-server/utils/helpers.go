package openapi

import (
	"reflect"
)

func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{
		Code: code,
		Body: body,
	}
}

func IsZeroValue(val interface{}) bool {
	return val == nil || reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
}

func AssertRecurseInterfaceRequired[T any](obj interface{}, callback func(T) error) error {
	return AssertRecurseValueRequired(reflect.ValueOf(obj), callback)
}

func AssertRecurseValueRequired[T any](value reflect.Value, callback func(T) error) error {
	switch value.Kind() {
	case reflect.Struct:
		obj, ok := value.Interface().(T)
		if !ok {
			return ErrTypeAssertionError
		}

		if err := callback(obj); err != nil {
			return err
		}

	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if err := AssertRecurseValueRequired(value.Index(i), callback); err != nil {
				return err
			}
		}
	}
	return nil
}
