package utils

func IfInterfaceValueIsNil(v interface{}) bool {
	if v == nil {
		return true
	}

	switch t := v.(type) {
	case int:
		return t == 0
	case float64:
		return t == 0
	case bool:
		return !t
	case string:
		return t == "" || t == "0001-01-01T00:00:00Z"
	}

	return false
}
