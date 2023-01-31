package putil

import "bytes"

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
