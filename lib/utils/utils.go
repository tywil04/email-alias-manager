package utils

import (
	"encoding/base64"
	"encoding/json"
)

func EncodeToBase64(v any) string {
	raw, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(raw)
}

func DecodeFromBase64(encoded string, v any) {
	raw, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return
	}
	json.Unmarshal(raw, &v)
}

func InArray(search any, array []any) bool {
	for _, arrayItem := range array {
		if search == arrayItem {
			return true
		}
	}
	return false
}

func IndexInArray(item any, array []any) int {
	positionInArray := -1
	for index, arrayItem := range array {
		if item == arrayItem {
			positionInArray = index
		}
	}
	return positionInArray
}

func RemoveFromArray(item any, array []any) []any {
	index := IndexInArray(item, array)
	if index == -1 {
		return nil
	}
	return append(array[:index], array[index+1:]...)
}
