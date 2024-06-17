package helpers

import "strings"

func SanitizeCSSId(element string) string {
	return "#" + strings.TrimLeft(element, "#")
}
