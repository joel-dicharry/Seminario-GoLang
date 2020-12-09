package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Read ...
func Read() {
	// abro el archivo y si hay un error se corta la ejecución
	file, err := os.Open("./file.f")

	// defer para ejecutar close() una linea antes de cualquier return
	defer file.Close()

	Check(err)

	// genero un escaner, me da elmetodo scan
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}

// Check ...
func Check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Write ...
func Write() {
	var data []byte

	// busca el archivo
	f, _ := os.Stat("file.f")
	// si existe el archivo lo leo
	if f != nil {
		data, _ = ioutil.ReadFile("file.f")
		data = append(data, []byte("content to file file.f\n")...)
	} else {
		// slide de byte que se va a escribir en el archivo file.f
		data = []byte("content to file file.f\n")
	}

	// ejecuto la escritura
	// le paso url del archivo, el dato que es un slide de bytes
	// y un permiso (644) que da lectura para todos y escritura solo al admin(si no existe el archivo lo crea)
	err := ioutil.WriteFile("./file.f", data, 0644)
	Check(err)
	fmt.Println("supuestamente se escribió:\n" + string(data))
}

func main() {
	Write()
	fmt.Println("\n")
	fmt.Println("Realmente se escibió:")
	fmt.Println("\n")
	Read()
}
