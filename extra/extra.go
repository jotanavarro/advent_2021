package extra

import (
	"log"
	"os"
	"strconv"
)

func ConvertToInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	return value
}

// CloseFile just takes care of closing a file and handling any possible error in a less verbose way than closures.
// In case of an error, just logs it then kill the program.
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
