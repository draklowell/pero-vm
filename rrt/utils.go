package rrt

import (
	"reflect"

	"lab.draklowell.net/routine-runtime/word"
)

func TypeToString(t int) string {
	switch t {
	case word.TypeArray:
		return "array"
	case word.TypeContainer:
		return "container"
	case word.TypeFloat:
		return "float"
	case word.TypeInteger:
		return "integer"
	case word.TypeNone:
		return "none"
	case word.TypeBytes:
		return "bytes"
	default:
		return "unknown"
	}
}

func ConvertToWord(value interface{}) word.Word {
	switch typedValue := value.(type) {
	case string:
		return []uint8(typedValue)
	case []interface{}:
		result := make([]word.Word, len(typedValue))
		for i, element := range typedValue {
			result[i] = ConvertToWord(element)
		}
		return result
	}

	if reflectedValue := reflect.ValueOf(value); reflectedValue.Kind() == reflect.Slice {
		if reflectedValue.IsNil() {
			return nil
		}

		typedValue := make([]interface{}, reflectedValue.Len())

		for i := 0; i < reflectedValue.Len(); i++ {
			typedValue[i] = reflectedValue.Index(i).Interface()
		}

		return ConvertToWord(typedValue)
	}

	switch typedValue := value.(type) {
	case int8:
		return uint8(typedValue)
	case uint8:
		return uint8(typedValue)
	case int:
		return int64(typedValue)
	case int16:
		return int64(typedValue)
	case int32:
		return int64(typedValue)
	case int64:
		return int64(typedValue)
	case uint16:
		return int64(typedValue)
	case uint32:
		return int64(typedValue)
	case uint64:
		return int64(typedValue)
	case float32:
		return float64(typedValue)
	case float64:
		return float64(typedValue)
	default:
		return nil
	}
}
