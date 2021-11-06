package handler

import "strings"

func ExtractFieldFromValidationMsg(msg string) string {
	field := strings.Split(strings.Split(msg, "'")[1], ".")
	return field[len(field)-1]
}
