package main

import (
	"errors"
	"fmt"
)

func main() {
	//Número inteiros

	var numero int64 = -10000000000
	fmt.Println(numero)

	var numero2 int32 = 100000
	fmt.Println(numero2)

	//"rune" pode ser utilizado para substituir o "int32"
	var numero3 rune = 20000
	fmt.Println(numero3)

	//"byte" pode ser utilizado para substituir o "int8"
	var numero4 byte = 123
	fmt.Println(numero4)

	//Strings

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//Variável vazia
	texto := 4
	fmt.Println(texto)

	//Valor boleano
	var boleano bool
	fmt.Println(boleano)

	//Tipo error, declarada para uma indicação de algum erro interno
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
