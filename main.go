package main

import (
	"bufio"
	//	"bytes"
	"fmt"
	"os"
)

func main() {
	buf := make([]byte, 10)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a cell: ")
	//reader := bytes.NewBufferString("45")
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	reader.Read(buf)
	//text := string(buf[:10])
	text := []byte(string(buf[:10]))
	writer.Write(text)
	fmt.Println("You entered: ", text)
}
