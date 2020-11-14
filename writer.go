package writer

import (
	"bufio"
	"fmt"
	"os"
)

//Input receives user input in the command line
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

//OpenFile takes a file name and trys to read it
func Open(filename string) (int, interface{}) {
	var Err interface{}

	file, err := os.Open(filename)
	var read int
	Err = err
	buffer := make([]byte, 100)
	for n, e := file.Read(buffer); e == nil; n, e = file.Read(buffer) {
		if n > 0 {
			read, Err = os.Stdout.Write(buffer[0:n])
			return read, Err
		}
	}
	return read, Err
}
