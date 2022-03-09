package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	slice = append(slice, 9)
	fmt.Println(slice)

	slice = array3[1:3]
	fmt.Println(slice)

	array2[1] = "Alterado com sucesso!"
	fmt.Println(array2)

	// ARRAYS INTERNO
	fmt.Println("-----------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

	slice3 = append(slice3, 23)
	slice3 = append(slice3, 28)

	fmt.Println(slice)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
