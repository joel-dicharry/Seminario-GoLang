package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// TicketOffice ...
type TicketOffice struct {
	passages map[int]*Passage
}

// NewTicketOffice ...
func NewTicketOffice() TicketOffice {
	passages := make(map[int]*Passage)
	return TicketOffice{
		passages,
	}
}

// AddPassage ...
func (to TicketOffice) AddPassage(p Passage) {
	// to.passages[p.idTicket] = &p
	to.Write(p)
}

// Print ...
func (to TicketOffice) Print() {
	fmt.Printf("%3s | %6s | %s\n", "ID", "Silla", "Compañia")
	fmt.Println("----------------------------------")
	for _, v := range to.passages {
		fmt.Printf("%3d | %6d | %s\n", v.idTicket, v.numSeat, v.bus)
	}
}

/*
// Print ...
func (to TicketOffice) Print(ID int) {
	fmt.Printf("%3s | %6s | %s\n", "ID", "Silla", "Compañia")
	fmt.Println("----------------------------------")
	fmt.Printf("%3d | %6d | %s\n", Passage.idTicket, to.numSeat, to.bus)
}*/

// GetByID ...
func (to TicketOffice) GetByID(ID int) (*Passage, error) {
	ret := to.passages[ID]
	if ret != nil {
		return ret, nil
	}
	return &Passage{}, fmt.Errorf("\nno se encontró nada con el ID %d", ID)
}

// Update ...
func (to TicketOffice) Update(p Passage) {
	_, e := to.GetByID(p.idTicket)
	if e != nil {
		to.Write(p)
	}
}

// Delete ...
func (to TicketOffice) Delete(ID int) {
	delete(to.passages, ID)
}

// PrintByID ...
func (to TicketOffice) PrintByID(ID int) {
	res, err := to.GetByID(ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nSe encontró exitosamente el pasaje con ID %d\n", ID)
		fmt.Printf("%3s | %6s | %s\n", "ID", "Silla", "Compañia")
		fmt.Println("----------------------------------")
		fmt.Printf("%3d | %6d | %s\n\n\n", res.idTicket, res.numSeat, res.bus)
	}
}

// Read ...
func Read(file string) ([]byte, error) {
	f, _ := os.Stat(file)
	if f != nil {
		f, _ := ioutil.ReadFile(file)
		return f, nil
	}
	return nil, errors.New("el archivo no se encuentra")
}

// Write ...
func (to TicketOffice) Write(p Passage) {
	jsonData := make(map[int]*Passage)
	var data []byte
	file := "passages.txt"
	f, err := Read(file)
	if err != nil {
		data, _ = json.Marshal(p)
		err = ioutil.WriteFile(file, data, 0644)
		Check(err)
	} else {
		err = json.Unmarshal(f, &jsonData)
		Check(err)
		jsonData[p.idTicket] = &p

		data, _ = json.Marshal(jsonData)
		err = ioutil.WriteFile(file, data, 0644)
		Check(err)
	}
}

// Check ...
func Check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
