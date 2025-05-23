package utils

import (
	"encoding/json"
	"fmt"
)

// StructToString converts any struct to a JSON string
func Marshal(v any) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("error marshaling struct: %w", err)
	}
	return string(bytes), nil
}

// StringToStruct converts a JSON string to a struct (pass a pointer to your struct)
func Unmarshal(data string, v any) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return fmt.Errorf("error unmarshaling string: %w", err)
	}
	return nil
}

func ToGeoPoint(lat float64, long float64) string {
	return fmt.Sprintf("POINT(%f %f)", long, lat)
}

func ToJSON(data string) map[string]any {
	var result map[string]any
	Unmarshal(data, result)
	return result
}
