package zlog

import (
	"fmt"
	"strings"
)

func Format(str string, data ...interface{}) string {
	var out strings.Builder
	runes := []rune(str)
	dataIndex := 0

	for i := 0; i < len(runes); i++ {
		if runes[i] == '{' {
			i++
			if i >= len(runes) {
				out.WriteRune('{')
				break
			}

			if runes[i] == '{' {
				out.WriteRune('{')
				continue
			}

			if runes[i] == '}' {
				if dataIndex >= len(data) {
					out.WriteString("{<extra>}")
					continue
				}
				value := fmt.Sprintf("%v", data[dataIndex])
				out.WriteString(value)
				dataIndex++
			} else {
				var buf strings.Builder
				for runes[i] != '}' {
					buf.WriteRune(runes[i])
					i++
				}
				out.WriteString(buf.String())
			}
			i++
		}

		if i < len(runes) && runes[i] == '}' {
			if i + 1 < len(runes) && runes[i+1] == '}' {
				i++
			}
		}

		if i < len(runes) {
			out.WriteRune(runes[i])
		}
	}

	for dataIndex < len(data) {
		out.WriteRune(' ')
		value := fmt.Sprintf("%v", data[dataIndex])
		out.WriteString(value)
		dataIndex++
	}

	return out.String()
}
