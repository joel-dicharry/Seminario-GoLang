package main

import "fmt"

//declaro un funcion de tipo close
type close func()

//el parametro es tipo close el cual es una funcion
func invokeClose(c close) {
	c()
}

func main() {
	f := func() {
		fmt.Println("Hello world")
	}
	invokeClose(f)
}

// en main declaro a f como un tipo func() y con su comportamiento
// ejecuto invokeClose(f) pasansole f que contiene el comportamiento de la funcion
// invokeClose ejecuta una func que es la que solicita en su parametro de entrada
