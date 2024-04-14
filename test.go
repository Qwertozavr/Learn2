package main

import (
	"bytes"
	"fmt"
)

func Test() {
	strings := []string{"This ", "is ", "even ", "more ", "performant "}

	bs := make([]byte, 100)
	bl := 0

	for _, val := range strings {
		bl += copy(bs[bl:], []byte(val))
	}

	fmt.Println(string(bs[:]))

	strings = []string{"This ", "is ", "even ", "more ", "performant "}
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}

	fmt.Println(buffer.String())
}
