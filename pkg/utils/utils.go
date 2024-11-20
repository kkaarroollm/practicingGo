package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func ParseValue(value string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.Bool:
		return value == "true", nil
	case reflect.Int:
		return strconv.Atoi(value)
	case reflect.String:
		return value, nil
	case reflect.Struct:
		parsedTime, err := time.Parse("2006-01-02", value)
		if err != nil {
			return nil, err
		}
		return parsedTime, nil
	default:
		return nil, fmt.Errorf("unsupported kind: %v", kind)
	}
}

func CompareValues(fieldValue, parsedValue interface{}) bool {
	return reflect.DeepEqual(fieldValue, parsedValue)
}
