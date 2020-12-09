package main

func main() {
	ticketO := NewTicketOffice() // Create

	p1 := Passage{1, 24, "El Rápido"}
	p2 := Passage{2, 9, "Chevallier"}
	p3 := Passage{3, 17, "Pullman"}
	p4 := Passage{4, 13, "General Belgrano"}
	p5 := Passage{5, 1, "Ñandú Del Sur"}

	ticketO.AddPassage(p1)
	ticketO.AddPassage(p2)
	ticketO.AddPassage(p3)
	ticketO.AddPassage(p4)
	ticketO.AddPassage(p5)

	ticketO.Print()

	ticketO.PrintByID(54) // Read
	ticketO.PrintByID(5)

	ticketO.Update(Passage{4, 27, "Gral Belgrano"}) // Update

	ticketO.Delete(3) //Delete
	ticketO.Print()

	ticketO.Update(Passage{1, 27, "cambio de ej"})
	ticketO.Delete(1)

}
