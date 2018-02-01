package jsonpath

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Marshal returns the JSON encoding of keysAndValues.
// Returns marshaled []byte and error in case if it can't be marshaled
func Marshal(keysAndValues map[string]string) ([]byte, error) {
	var result interface{}
	var err error
	for key, value := range keysAndValues {
		if err = buildEmbedded(&result, strings.Split(key, ".")[:], value); err != nil {
			return nil, err
		}
	}
	return json.Marshal(result)
}

func buildEmbedded(result *interface{}, keys []string, value string) error {
	var err error
	if len(keys) == 0 {
		*result = value
		return err
	}
	currentKey := keys[0]
	currentKeyIndex, err := strconv.Atoi(currentKey)
	if err == nil {
		if *result == nil {
			*result = []interface{}{}
		}
		currentResult, ok := (*result).([]interface{})
		if !ok {
			return fmt.Errorf("Wrong key dimension %v", keys)
		}
		if len(currentResult) < currentKeyIndex+1 {
			for i := 0; i < currentKeyIndex+1; i++ {
				if len((*result).([]interface{})) < i+1 {
					a := (*result).([]interface{})
					*result = append(a, nil)
				}
			}
		}
		var nextLevelValue = (*result).([]interface{})[currentKeyIndex]
		if err = buildEmbedded(&nextLevelValue, keys[1:], value); err != nil {
			return err
		}
		(*result).([]interface{})[currentKeyIndex] = nextLevelValue
	} else {
		if *result == nil {
			*result = map[string]interface{}{}
		}
		_, ok := (*result).(map[string]interface{})
		if !ok {
			return fmt.Errorf("Wrong key dimension %v", keys)
		}
		var nextLevelValue = (*result).(map[string]interface{})[currentKey]
		if err = buildEmbedded(&nextLevelValue, keys[1:], value); err != nil {
			return err
		}
		switch currentKey {
		case "num()":
			{
				strNextLevelValue := nextLevelValue.(string)
				*result, _ = strconv.ParseFloat(strNextLevelValue, 64)
			}
		case "bool()":
			{
				strNextLevelValue := nextLevelValue.(string)
				*result = (strNextLevelValue == "true")
			}
		case "[]":
			{
				strNextLevelValue := nextLevelValue.(string)
				*result = strings.Split(strNextLevelValue, ",")
			}
		default:
			{
				(*result).(map[string]interface{})[currentKey] = nextLevelValue
			}
		}
	}
	return err
}
