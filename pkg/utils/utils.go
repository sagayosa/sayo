package utils

import "strings"

func StringPlus(segments ...string) string {
	var builder strings.Builder

	for _, seg := range segments {
		builder.WriteString(seg)
	}

	return builder.String()
}
