package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Vinícius",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro_nome": "Vinícius",
			"ultimo_nome":   "Melo",
		},

		"curso": {
			"curso_nome":     "Ciência da Computação",
			"faculdade_nome": "IFSULDEMINAS - Muzambinho",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
}
