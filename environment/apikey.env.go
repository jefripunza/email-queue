package environment

import "os"

func GetApiKey() string {
	key := os.Getenv("API_KEY")
	if key == "" {
		key = "test-key"
	}
	return key
}
