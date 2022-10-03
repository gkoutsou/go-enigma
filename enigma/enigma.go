package enigma

import (
	"fmt"
)

type Machine struct {
	RotorA    *Rotor
	RotorB    *Rotor
	RotorC    *Rotor
	Reflector Reflector
}

func (e *Machine) rotate() {
	if e.RotorA.rotate() {
		fmt.Println("rotor B should rotate!")

		if e.RotorB.rotate() {
			fmt.Println("rotor C should rotate!")
			e.RotorC.rotate()
		}
	}

	// fmt.Printf("rotor position: %c%c%c\n", int2rune(e.RotorC.currentPos), int2rune(e.RotorB.currentPos), int2rune(e.RotorA.currentPos))
}

func (e *Machine) Init(c, b, a int8) {
	e.RotorA.init(a)
	e.RotorB.init(b)
	e.RotorC.init(c)

	e.Reflector.init()
}

func (e *Machine) Press(inputChar rune) rune {
	input := rune2Int(inputChar)

	e.rotate()

	outputA := e.RotorA.Pass(input)
	outputB := e.RotorB.Pass(outputA)
	outputC := e.RotorC.Pass(outputB)
	outputR := e.Reflector.Pass(outputC)
	outputC2 := e.RotorC.PassBack(outputR)
	outputB2 := e.RotorB.PassBack(outputC2)
	outputA2 := e.RotorA.PassBack(outputB2)

	fmt.Printf("rotor encryption: %c->%c->%c->%c->%c->%c->%c\n", int2rune(outputA), int2rune(outputB), int2rune(outputC), int2rune(outputR), int2rune(outputC2), int2rune(outputB2), int2rune(outputA2))

	return int2rune(outputA2)
}

func (e *Machine) Type(text string) string {
	var chars []rune
	for _, c := range text {
		chars = append(chars, e.Press(c))
	}

	return string(chars)
}

func int2rune(i int8) rune {
	return rune(i) + 64
}

func rune2Int(c rune) int8 {
	return int8(c - 64)
}
