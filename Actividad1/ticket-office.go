package main

import (
	"fmt"
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
	to.passages[p.idTicket] = &p
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
	to.passages[p.idTicket] = &p
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
