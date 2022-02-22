package main

import "fmt"

func main() {
	var variavel1 string = "variavel_1"
	variavel2 := "variavel_2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3  string = "teste_variável_3"
		vaariavel4 string = "teste_variável_4"
	)

	fmt.Println(variavel3, vaariavel4)

	variavel5, variavel6 := "variável_5", "variável_6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
