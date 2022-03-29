package main

import "fmt"

func main() {
	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 15

	soma2 := numero1 + numero2
	fmt.Println(soma2)

	

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	

	
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	

	
	fmt.Println(".............")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	

	
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)


}
