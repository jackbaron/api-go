package helpers

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrInternal = errors.New("internal validation error")
)

// FindStructFieldJSONName gets the value of the `json` tag for the given field in the given struct
// Implementation inspired by https://github.com/jellydator/validation/blob/v1.0.0/struct.go#L70
func FindStructFieldJSONName(structPtr interface{}, fieldPtr interface{}, errorTag string) string {
	value := reflect.ValueOf(structPtr)
	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		// must be a pointer to a struct
		panic("must be a pointer to a struct")
	}
	if value.IsNil() {
		// treat a nil struct pointer as valid
		panic("treat a nil struct pointer as valid")
	}
	value = value.Elem()

	fv := reflect.ValueOf(fieldPtr)
	if fv.Kind() != reflect.Ptr {
		panic("must be a pointer to a field")
	}
	ft := findStructField(value, fv)
	if ft == nil {
		panic("field not found")
	}
	return getErrorFieldName(ft, errorTag)
}

// findStructField looks for a field in the given struct.
// The field being looked for should be a pointer to the actual struct field.
// If found, the field info will be returned. Otherwise, nil will be returned.
// Implementation borrowed from https://github.com/jellydator/validation/blob/v1.0.0/struct.go#L134
func findStructField(structValue reflect.Value, fieldValue reflect.Value) *reflect.StructField {
	ptr := fieldValue.Pointer()
	for i := structValue.NumField() - 1; i >= 0; i-- {
		sf := structValue.Type().Field(i)
		if ptr == structValue.Field(i).UnsafeAddr() {
			// do additional type comparison because it's possible that the address of
			// an embedded struct is the same as the first field of the embedded struct
			if sf.Type == fieldValue.Elem().Type() {
				return &sf
			}
		}
		if sf.Anonymous {
			// delve into anonymous struct to look for the field
			fi := structValue.Field(i)
			if sf.Type.Kind() == reflect.Ptr {
				fi = fi.Elem()
			}
			if fi.Kind() == reflect.Struct {
				if f := findStructField(fi, fieldValue); f != nil {
					return f
				}
			}
		}
	}
	return nil
}

// getErrorFieldName returns the name that should be used to represent the validation error of a struct field.
// Implementation borrowed from https://github.com/jellydator/validation/blob/v1.0.0/struct.go#L162
func getErrorFieldName(f *reflect.StructField, errorTag string) string {
	if tag := f.Tag.Get(errorTag); tag != "" && tag != "-" {
		if cps := strings.SplitN(tag, ",", 2); cps[0] != "" {
			return cps[0]
		}
	}
	return f.Name
}
