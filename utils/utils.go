package utils

import (
	"encoding/json"
)

func StructToMap(s interface{}) map[string]string {
	// struct转json
	jsonTypeData, _ := json.Marshal(s)
	mapType := make(map[string]string)
	// json转map
	json.Unmarshal(jsonTypeData, &mapType)
	return mapType
}

func StructToReturnData(s interface{}) map[string]interface{} {
	// struct转json
	jsonTypeData, _ := json.Marshal(s)
	mapType := make(map[string]interface{})
	// json转map
	json.Unmarshal(jsonTypeData, &mapType)
	return mapType
}
