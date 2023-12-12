package string_type

import "strings"

func stringBuilder(s []string) string {
	builder := strings.Builder{}
	for _, str := range s {
		builder.WriteString(str)
	}
	return builder.String()
}
