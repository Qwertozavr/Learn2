package main

import (
	"fmt"
)

// класс Алиса (солосовой помошник)
type Alice struct{}

// реакция собаки
func (alc *Alice) Welcome() {
	fmt.Println("Добро пожаловать!")
}

// целевой интерфейс - Target
type Global_out interface {
	Reaction()
}

// адаптер для Алисы
type AliceAdapt struct {
	*Alice
}

// реакция Алисы
func (adapter *AliceAdapt) Reaction() {
	adapter.Welcome()
}

// конструктор адаптера для собаки
func NewDogAdapter(alc *Alice) Global_out {
	return &AliceAdapt{alc}
}

// класс - чайник
type Kattle struct {
}

// реакция чайника - адаптер не нужен, нужный метод итак есть
func (w *Kattle) Reaction() {
	fmt.Println("Добрый день, чаю?")
}

func Ex21() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex21")
	fmt.Println("---------------------------------------")

	devices := [2]Global_out{NewDogAdapter(&Alice{}), &Kattle{}}

	fmt.Println("Открываете дверь и заходите домой\n")
	for _, member := range devices {
		member.Reaction()
	}

	fmt.Println("=======================================\n")
}
