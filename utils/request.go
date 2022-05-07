package utils

// CheckAndSetString for update payload
func CheckAndSetString(existingField, newField string) string {
	if newField != "" {
		return newField
	}
	return existingField
}
