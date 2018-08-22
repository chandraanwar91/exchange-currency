package helpers

import (
	"strings"
    "unicode"
)

func CamelCaseToSnakeCase(camelCase string) (inputUnderScoreStr string) {
    //snake_case to camelCase
    for k, v := range camelCase {
        if(isUpperCase(string(v)) && k > 0 && unicode.IsLetter(v)){
            inputUnderScoreStr += "_"
        }
        inputUnderScoreStr += strings.ToLower(string(v))
    }
    return inputUnderScoreStr
}

func isUpperCase(str string) bool {
    if str == strings.ToUpper(str) {
        return true
    }
    return false
}