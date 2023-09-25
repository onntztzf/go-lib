package convertor

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"regexp"
	"strconv"
)

// ToBytes converts an interface to bytes using JSON encoding.
func ToBytes(data interface{}) ([]byte, error) {
	return jsoniter.Marshal(data)
}

// ToString converts a value to a string.
func ToString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case float32, float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%v", v)
	case string:
		return v
	default:
		newValue, _ := jsoniter.MarshalToString(value)
		return newValue
	}
}

// ToFloat64 converts a value to a float64.
func ToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return float64(v.(int64)), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(v.(uint64)), nil
	case float32, float64:
		return v.(float64), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0.0, fmt.Errorf("invalid interface type %T", value)
	}
}

// ToInt64 converts a value to an int64.
func ToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return v.(int64), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(v.(uint64)), nil
	case float32, float64:
		return int64(v.(float64)), nil
	case string:
		return strconv.ParseInt(v, 0, 64)
	default:
		return 0, fmt.Errorf("invalid interface type %T", value)
	}
}

// ToJson converts a value to a valid JSON string.
func ToJson(value interface{}) (string, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// StructToMap converts a struct to a map, only converting exported struct fields.
// Map keys are specified by the same struct field tag `json` value.
func StructToMap(value interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data type %T not supported, should be struct or pointer to struct", value)
	}
	res := make(map[string]interface{})
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		result, err := regexp.MatchString("^[A-Z]", name)
		if err != nil || !result {
			continue
		}
		tag := t.Field(i).Tag.Get("json")
		if len(tag) == 0 {
			res[name] = v.Field(i).Interface()
		} else {
			res[tag] = v.Field(i).Interface()
		}
	}
	return res, nil
}
