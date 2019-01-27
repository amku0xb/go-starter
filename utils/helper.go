package utils

import (
	"encoding/json"
)

func FillStructFromMap(data map[string]interface{}, result interface{}) {
	jsonData, _ := json.Marshal(data)
	json.Unmarshal(jsonData, result)
}
