package putil

import "bytes"

// Concat will attempt to put multiple values together if possible or return a map otherwise
func Concat(obj ...any) any {
	canString := true
	for _, v := range obj {
		if _, ok := v.(string); !ok {
			canString = false
			break
		}
	}
	if canString {
		res := &bytes.Buffer{}
		for _, v := range obj {
			res.WriteString(v.(string))
		}
		return res.String()
	}

	return map[string]any{
		"@concat": obj,
	}
}
