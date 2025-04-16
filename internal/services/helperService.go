package services

func updateField(newValue, oldValue string) string {
	if newValue == "" {
		return oldValue
	}

	return newValue
}

func updateIntPointerField(newValue, oldValue *int32) *int32 {
	if newValue == nil {
		return oldValue
	}

	return newValue
}
