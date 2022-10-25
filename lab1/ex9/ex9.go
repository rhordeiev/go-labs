package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false значение по умолчанию
	fmt.Println("second = ", second)      // false значение по умолчанию
	fmt.Println("third  = ", third)       // true присвоенное значение
	fmt.Println("fourth = ", fourth)      // false обратное значение переменной third
	fmt.Println("fifth  = ", fifth, "\n") // true присвоенное значение

	fmt.Println("!true  = ", !true)        // false обратное значение true
	fmt.Println("!false = ", !false, "\n") // true обратное значение false

	fmt.Println("true && true   = ", true && true)         // true знак && истинный, если два его операнда тоже истинны
	fmt.Println("true && false  = ", true && false)        // false знак && истинный, если два его операнда тоже истинны
	fmt.Println("false && false = ", false && false, "\n") // false знак && истинный, если два его операнда тоже истинны

	fmt.Println("true || true   = ", true || true)         // true знак || истинный, если хотя бы один из его операндов истинный
	fmt.Println("true || false  = ", true || false)        // true знак || истинный, если хотя бы один из его операндов истинный
	fmt.Println("false || false = ", false || false, "\n") // false знак || истинный, если хотя бы один из его операндов истинный

	fmt.Println("2 < 3  = ", 2 < 3)        // true 2 > 3 - верно, поэтому true
	fmt.Println("2 > 3  = ", 2 > 3)        // false 2 не больше 3, поэтому false
	fmt.Println("3 < 3  = ", 3 < 3)        // false 3 не меньше 3, поэтому false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true 3 = 3 - верно, поэтому true
	fmt.Println("3 > 3  = ", 3 > 3)        // false 3 не больше 3, поэтому false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true 3 = 3 - верно, поэтому true
	fmt.Println("2 == 3 = ", 2 == 3)       // false 2 и 3 не равны, поэтому false
	fmt.Println("3 == 3 = ", 3 == 3)       // true 3 = 3 - верно, поэтому true
	fmt.Println("2 != 3 = ", 2 != 3)       // true 3 = 3 - верно, поэтому true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false 3 и 3 равны, поэтому false

	//Задание.
	//1. Пояснить результаты операций
}
