package main

import (
	"errors"
	"fmt"
)

// se especifican el/los parametros y retornos(siempre se retorna un error)
//si suma empiza con mayúscula va a ser un func publica para otros packages y necesita ser documentado
//(documentacion)
func suma(a int, b int) (int, error) { //se pueden nombrar los tipos de retorno, y no se necesitaria usar 'return'
	if a < b {
		return 0, errors.New("el 1er valor es menor que el 2do")
		//por convencion el error se empiza con minuscula
		//y no termina con un punto, porque se podria concatenar con otro error
	}
	return a + b, nil
}

func main() {
	fmt.Println("Hello World")
	r, err := suma(1, 2) // almaceno los retornos de suma()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
