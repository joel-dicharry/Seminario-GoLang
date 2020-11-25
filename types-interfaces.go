package main

import "fmt"

//Printable arranca con mayúscula para que sea exportable
type Printable interface {
	print()
}

// recibe por parametro un tipo de la interface Printable y usa su metodo, aplicado a su manera en el struct
func invokePrint(p Printable) {
	p.print()
}

type person struct {
	name string
}

// '(pe person)' es un reciver que asocia la funcion a un struct
// esta funcion es void por eso no se especifican los valores de return
func (pe *person) print() { //al implementar print() "person" es "Printable"
	fmt.Println("[person]", pe.name)
}

//pe *person hace referencia a un puntero a la instancia de la persona
//pe person hace una copia, funciona como un metodo estático en java
func (pe *person) clean() {
	pe.name = ""
}

func main() {
	p := &person{name: "Juan"}
	r := &person{name: "Carlos"}
	p.print()
	p.clean()
	p.print()
	fmt.Println("-------------------")
	r.print()
	r.clean()
	r.print()
	p1 := &person{name: "persona 1"}
	p2 := &person{name: "persona 2"}
	invokePrint(p1) //el metodo toma como parameto a un tipo 'person' y lo castea a 'Printable', ya que adoptan el metodo print()
	invokePrint(p2)
}
