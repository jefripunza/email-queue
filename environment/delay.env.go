package environment

import (
	"os"
	"strconv"
)

func GetDelay() int {
	key := os.Getenv("DELAY_SECOND")
	if key == "" {
		key = "5"
	}
	delay, _ := strconv.Atoi(key)
	return delay
}
