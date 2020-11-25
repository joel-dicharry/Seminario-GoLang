package main

import "fmt"

func main() {

	fmt.Println("Array\n")
	var arr [2]int
	arr[0] = 1
	arr[1] = 2

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println("------------------------------------------------")
	fmt.Println("Slice\n")

	var s []int
	// s := make([]int, 10) // make() va a devolve un slice int que va a tener capacity de 10
	s = append(s, 2) //append devuelve s + 2 en su ultima posicion
	// fmt.Println("%p\n", s) //con esto veo la posicio de memoria y contato como se crea un nuevo slice cada ves que agrego un dato
	s = append(s, 3)
	// fmt.Println("%p\n", s)
	s = append(s, 4)
	// fmt.Println("%p\n", s)
	s = append(s, 5)
	// fmt.Println("%p\n", s)

	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("Map\n")

	m := make(map[int]string)
	m[0] = "h"
	m[1] = "o"
	m[2] = "l"
	m[3] = "a"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
