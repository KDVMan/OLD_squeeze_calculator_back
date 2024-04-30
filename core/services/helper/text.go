package core_services_helper

import (
	"encoding/json"
	"log"
	"strings"
	"unicode"
)

func ModelToJson(data interface{}) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		log.Println("[ModelToJson] failed:", err)
		return "error"
	} else {
		return string(jsonData)
	}
}

func ToSnakeCase(str string) string {
	var result strings.Builder
	result.Grow(len(str) + 5)

	for i, r := range str {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}

			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
