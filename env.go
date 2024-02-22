package gravity

import (
	"bufio"
	"os"
	"strings"
)

func readEnv() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
