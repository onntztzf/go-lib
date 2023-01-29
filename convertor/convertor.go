package convertor

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"regexp"
	"strconv"
)

// ToBytes convert interface to bytes
func ToBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToString convert value to string
func ToString(value interface{}) string {
	res := ""
	if value == nil {
		return res
	}
	v := reflect.ValueOf(value)
	switch value.(type) {
	case float32, float64:
		res = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		return res
	case int, int8, int16, int32, int64:
		res = strconv.FormatInt(v.Int(), 10)
		return res
	case uint, uint8, uint16, uint32, uint64:
		res = strconv.FormatUint(v.Uint(), 10)
		return res
	case string:
		res = v.String()
		return res
	case []byte:
		res = string(v.Bytes())
		return res
	default:
		newValue, _ := json.Marshal(value)
		res = string(newValue)
		return res
	}
}

// ToFloat64 convert value to a float64, if input is not a float return 0.0 and error
func ToFloat64(value interface{}) (float64, error) {
	v := reflect.ValueOf(value)
	res := 0.0
	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = float64(v.Int())
		return res, nil
	case uint, uint8, uint16, uint32, uint64:
		res = float64(v.Uint())
		return res, nil
	case float32, float64:
		res = v.Float()
		return res, nil
	case string:
		res, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			res = 0.0
		}
		return res, err
	default:
		err := fmt.Errorf("ToInt: unvalid interface type %T", value)
		return res, err
	}
}

// ToInt64 convert value to a int64, if input is not a numeric format return 0 and error
func ToInt64(value interface{}) (int64, error) {
	v := reflect.ValueOf(value)
	var res int64
	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = v.Int()
		return res, nil
	case uint, uint8, uint16, uint32, uint64:
		res = int64(v.Uint())
		return res, nil
	case float32, float64:
		res = int64(v.Float())
		return res, nil
	case string:
		res, err := strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			res = 0
		}
		return res, err
	default:
		err := fmt.Errorf("ToInt: invalid interface type %T", value)
		return res, err
	}
}

// ToJson convert value to a valid json string
func ToJson(value interface{}) (string, error) {
	res, err := jsoniter.MarshalToString(value)
	if err != nil {
		return "", err
	}
	return res, nil
}

// StructToMap convert struct to map, only convert exported struct field
// map key is specified same as struct field tag `json` value
func StructToMap(value interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", value)
	}
	res := make(map[string]interface{})
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		result, err := regexp.MatchString("^[A-Z]", name)
		if err != nil || result == false {
			continue
		}
		tag := t.Field(i).Tag.Get("json")
		if len(tag) == 0 {
			res[name] = v.Field(i).Interface()
			continue
		}
		res[tag] = v.Field(i).Interface()
	}
	return res, nil
}
