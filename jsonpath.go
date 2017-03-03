package jsonpath

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Marshal returns the JSON encoding of keysAndValues.
// Returns marshaled []byte and error in case if it can't be marshaled
func Marshal(keysAndValues map[string]string) ([]byte, error) {
	var result interface{}
	for key, value := range keysAndValues {
		buildEmbedded(&result, strings.Split(key, ".")[:], value)
	}
	return json.Marshal(result)
}

func buildEmbedded(result *interface{}, keys []string, value string) {
	if len(keys) == 0 {
		*result = value
		return
	}
	currentKey := keys[0]
	currentKeyIndex, err := strconv.Atoi(currentKey)
	if err == nil {
		if *result == nil {
			*result = []interface{}{}
		}
		if len((*result).([]interface{})) < currentKeyIndex+1 {
			for i := 0; i < currentKeyIndex+1; i++ {
				if len((*result).([]interface{})) < i+1 {
					a := (*result).([]interface{})
					*result = append(a, nil)
				}
			}
		}
		var nextLevelValue = (*result).([]interface{})[currentKeyIndex]
		buildEmbedded(&nextLevelValue, keys[1:], value)
		(*result).([]interface{})[currentKeyIndex] = nextLevelValue
	} else {
		if *result == nil {
			*result = map[string]interface{}{}
		}
		var nextLevelValue = (*result).(map[string]interface{})[currentKey]
		buildEmbedded(&nextLevelValue, keys[1:], value)
		switch currentKey {
		case "num()":
			{
				strNextLevelValue := nextLevelValue.(string)
				*result, _ = strconv.ParseFloat(strNextLevelValue, 64)
			}
		default:
			{
				(*result).(map[string]interface{})[currentKey] = nextLevelValue
			}
		}
	}
}
