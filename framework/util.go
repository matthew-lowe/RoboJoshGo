package framework

import (
	"regexp"
)

func TagToUserId(tag string) string {
	return tag[3:][:len(tag)-4]
}

// Is the given code a valid hex color code?
func VerifyHexColor(code string) (bool, error) {
	matched, err := regexp.MatchString(`#?([A-F]|[a-f]|\d){6}`, code)

	if err != nil {
		return false, err
	}

	return matched, nil
}
