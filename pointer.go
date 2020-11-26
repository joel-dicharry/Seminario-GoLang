package main

import "fmt"

type person struct {
	name string
}

//esta funcion no incorpora la referencia a la variable del reciver por lo que la modificacion se muere al terminar la funcion
func (pe person) changeName(n string) {
	pe.name = n
}

//esta funcion hace referencia al puntero de la 'person' que ejecuta la funcion, modificando el valor original
func (pe *person) changeNameOk(n string) {
	pe.name = n
}

func main() {
	p := person{"carlos"}
	//con printf se imprimen las variables con distintos formatos, segun los formatos de impresión
	//%p, imprime la representacion de la posicion de memoria de p(&p), y %v imprime el valor de la variable(p) en un formato predeterminado
	fmt.Printf("valor original\n%p %v \n", &p, p)

	p.changeName("pedro")
	fmt.Printf("Resultado de la funcion sin referencia al puntero\n%p %v \n", &p, p)

	p.changeNameOk("juan")
	fmt.Printf("Resultado de la funcion con referencia al puntero\n%p %v \n", &p, p)
}
