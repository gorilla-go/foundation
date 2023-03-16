package strings

import "strings"
import "unicode"

func Title(str string) string {
	var camelCaseBuilder strings.Builder
	for i, r := range str {
		if r == '_' {
			continue
		}
		if i > 0 && str[i-1] == '_' {
			camelCaseBuilder.WriteRune(unicode.ToUpper(r))
		} else {
			camelCaseBuilder.WriteRune(unicode.ToLower(r))
		}
	}
	return UcFirst(camelCaseBuilder.String())
}
