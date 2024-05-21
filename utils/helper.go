package utils

// Helper function to get string value from map
func GetStringValue(data map[string]interface{}, key string) string {
	value, ok := data[key].(string)
	if !ok {
		return "" // return empty string if not found or not a string
	}
	return value
}

// Helper function to get int value from map
func GetIntValue(data map[string]interface{}, key string) int {
	value, ok := data[key].(float64)
	if !ok {
		return 0 // return 0 if not found or not a number
	}
	return int(value)
}

// Helper function to get bool value from map
func GetBoolValue(data map[string]interface{}, key string) bool {
	value, ok := data[key].(bool)
	return ok && value
}