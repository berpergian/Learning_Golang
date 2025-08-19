package helper

func JsonError(message string) string {
	return `{"message": "` + message + `"}`
}
