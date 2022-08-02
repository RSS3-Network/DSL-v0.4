package utils

func IfInterfaceValueIsNil(v interface{}) bool {
	if v == nil {
		return true
	}

	switch v.(type) {
	case int:
		return v.(int) == 0
	case bool:
		return v.(bool) == false
	case string:
		return v.(string) == "" || v.(string) == "0001-01-01T00:00:00Z"
	}

	return false
}
