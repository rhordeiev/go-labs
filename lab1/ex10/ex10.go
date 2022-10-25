package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	//Задание.
	//1. Вывести украинскую букву 'Ї'
	//2. Пояснить назначение типа "rune"
	var chartype2 rune = 'Ї'

	fmt.Printf("Code '%c' - %d\n", chartype2, chartype2)
}

//rune - синоним int32
//rune используется для представления символа Unicode, в то время как символы ASCII может представлять только тип данных int32
