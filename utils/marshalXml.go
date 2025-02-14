package utils

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
)

func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func XMLToJSON[T any](xmlData []byte, v *T) ([]byte, error) {
	if err := xml.Unmarshal(xmlData, v); err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(v)

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func XMLtoGoStruct(xmlData []byte, v interface{}) error {
	return xml.Unmarshal(xmlData, v)
}
