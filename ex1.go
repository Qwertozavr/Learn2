package main

import "fmt"

type Human struct {
	age    uint8
	height uint8
	name   string
}

type Action struct {
	Human
}

func (h Human) height_dev_age() {
	fmt.Printf("heaight / age = %d\n", h.height/h.age)
}

func Ex1() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex1                  ")
	fmt.Println("---------------------------------------")

	h := Human{age: 20, height: 182, name: "Dima"}
	a := Action{h}
	a.height_dev_age()

	fmt.Println("=======================================\n")
}
