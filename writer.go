package writer

import (
	"bufio"
	"fmt"
	"os"
)

func Input(message ...string) string {
	params := ":"
	if len(message) > 0 {
		params = message[0]
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(params)
	scanner.Scan()
	output := scanner.Text()
	return output
}
