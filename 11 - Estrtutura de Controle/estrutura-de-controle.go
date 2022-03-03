package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("\n O número é maior que 15")
	} else {
		fmt.Print("\n O número não é maior que 15")
	}

	if outro_numero := numero; outro_numero > 0 {
		fmt.Println("\n Este número é maior que ", outro_numero)
	} else if numero < -10 {
		fmt.Println("\n É menor que", outro_numero)
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
