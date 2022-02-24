package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivos Strucs")

	var u usuario
	u.nome = "Vinícius"
	u.idade = 23
	fmt.Println(u)

	enderecoExemplo := endereco{"Jardim dos Ipês", 0}
	fmt.Println(enderecoExemplo)

	usuario2 := usuario{"VInícius", 23}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Vinícius"}
	fmt.Println(usuario3)
}
