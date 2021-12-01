package extra

import (
	"log"
	"strconv"
)

func ConvertToInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	return value
}
