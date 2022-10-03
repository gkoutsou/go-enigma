package enigma

import "fmt"

type Machine struct {
	RotorA *Rotor
	RotorB *Rotor
	RotorC *Rotor
}

func (e *Machine) rotate() {
	if e.RotorA.rotate() {
		fmt.Println("rotor B should rotate!")

		if e.RotorB.rotate() {
			fmt.Println("rotor C should rotate!")
			e.RotorC.rotate()
		}
	}

	fmt.Printf("rotor position: %c%c%c\n", int2rune(e.RotorC.currentPos), int2rune(e.RotorB.currentPos), int2rune(e.RotorA.currentPos))
}

func (e *Machine) Init(a, b, c int8) {
	e.RotorA.init(a)
	e.RotorB.init(b)
	e.RotorC.init(c)
}

func (e *Machine) Press(inputChar rune) rune {
	input := rune2Int(inputChar)

	e.rotate()

	outputA := e.RotorA.Pass(input)
	outputB := e.RotorB.Pass(outputA)
	outputC := e.RotorC.Pass(outputB)

	fmt.Printf("rotor encryption: %c->%c->%c\n", int2rune(outputA), int2rune(outputB), int2rune(outputC))

	return int2rune(outputC)
}

func int2rune(i int8) rune {
	return rune(i) + 64
}

func rune2Int(c rune) int8 {
	return int8(c - 64)
}
