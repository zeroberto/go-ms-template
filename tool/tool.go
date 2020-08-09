package tool

// Contains is responsible for checking if there is an equal value in an array
func Contains(value interface{}, values []interface{}) bool {
	for _, v := range values {
		if value == v {
			return true
		}
	}
	return false
}

// ContainsKey is responsible for checking if there is an equal key in an map
func ContainsKey(key interface{}, values map[interface{}]interface{}) bool {
	for k := range values {
		if key == k {
			return true
		}
	}
	return false
}

// ContainsString is responsible for checking if there is an equal string value in an array
func ContainsString(value string, values []string) bool {
	for _, v := range values {
		if value == v {
			return true
		}
	}
	return false
}

// ContainsStringKey is responsible for checking if there is an equal string key in an map
func ContainsStringKey(key string, values map[string]interface{}) bool {
	for k := range values {
		if key == k {
			return true
		}
	}
	return false
}
