package main

import "fmt"

func clousure() func() {
	texto := "Dentro da funçaoes clousure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	texto := "Dentro da main"
	fmt.Println(texto)

	funcaoNova := clousure()
	funcaoNova()

}
