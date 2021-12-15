package utils

import (
	"encoding/json"
	"log"
)

func ToJSONString(data interface{}) (string, error) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return BytesToString(jsonString), nil
}

func ParseStruct(jsonString string, result interface{}) (err error) {
	err = json.Unmarshal(StringToBytes(jsonString), result)
	if err != nil {
		return
	}
	return
}
