package main

import (
	"fmt"
)

type Dict struct {
	k string
	v interface{}
}

func FlattenDictionary(dict map[string]interface{}) map[string]string {
	if dict == nil {
		return nil
	}
	output := make(map[string]string)
	helper(dict, output, "")
	return output
}

func helper(data map[string]interface{}, output map[string]string, pKey string) {
	for key, val := range data {
		switch v := val.(type) {
		case string:
			newKey := ""
			if pKey != "" {
				newKey = fmt.Sprintf("%s.", pKey)
			}
			if key == "" {
				newKey = pKey
			}
			output[fmt.Sprintf("%s%s", newKey, key)] = fmt.Sprintf("%v", v)
		default:
			m := v.(map[string]interface{})
			if pKey != "" {
				key = fmt.Sprintf("%s.%s", pKey, key)
			}
			helper(m, output, key)
		}
	}
}

func main() {
	dect := map[string]interface{}{"Key1": "1", "key2": map[string]interface{}{"a": "2", "b": "3", "c": map[string]interface{}{"d": "3", "e": map[string]interface{}{"": "1"}}}}

	output := FlattenDictionary(dect)
	for k, v := range output {
		fmt.Print(k + "  ->  ")
		fmt.Println(v + " ")
	}
}
