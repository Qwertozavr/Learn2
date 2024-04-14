package main

import (
	"fmt"
)

func update(p *int) {
	b := 2
	p = &b
}

func Test2() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
