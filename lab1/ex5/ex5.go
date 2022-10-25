package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println("Синонимы целых типов\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")

	//Задание.
	//1. Определить разрядность ОС

	switch bits.UintSize {
	case 32:
		fmt.Println("ОС 32-розрядна")
	case 64:
		fmt.Println("ОС 64-розрядна")
	default:
		fmt.Println("Розрядність ОС неможливо визначити")
	}
}
