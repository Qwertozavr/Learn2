package main

import (
	"fmt"
)

func del_elem_1(slc []int, i int) ([]int, error) {
	if i > len(slc) || i < 0 {
		return nil, fmt.Errorf("Index out of range!")
	}

	slc[i] = slc[len(slc)-1]
	return slc[:len(slc)-1], nil
}

func del_elem_2(slc []int, i int) ([]int, error) {
	result := make([]int, 0)
	result = append(result, slc[:i]...)
	return append(result, slc[i+1:]...), nil
}

func Ex23() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex23")
	fmt.Println("---------------------------------------")
	var i int
	slc := []int{1, 2, 3, 4, 5}

	fmt.Println("Write index to delete: ")
	fmt.Scan(&i)

	ans, err := del_elem_1(slc, i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("First answer: %d\n", ans)

	ans, err = del_elem_2(slc, i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("First answer: %d\n", ans)

	fmt.Println("=======================================\n")
}
