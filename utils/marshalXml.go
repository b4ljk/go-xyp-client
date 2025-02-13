package utils

import (
	"encoding/json"
	"encoding/xml"
)

func XMLToJSON[T any](xmlData []byte, v *T) ([]byte, error) {
	if err := xml.Unmarshal(xmlData, v); err != nil {
		return nil, err
	}

	return json.MarshalIndent(v, "", "  ")
}
