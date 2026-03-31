package environment

import (
	"os"
)

func GetServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
