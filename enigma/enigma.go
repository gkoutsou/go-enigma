package enigma

import (
	"github.com/pkg/errors"
)

type Machine struct {
	Reflector Reflector
	RotorA    *Rotor
	RotorB    *Rotor
	RotorC    *Rotor
	plugboard Plugboard
}

func (e *Machine) rotate() {
	if e.RotorA.rotate() {
		if e.RotorB.rotate() {
			e.RotorC.rotate()
		}
	}
}

func (e *Machine) Init(settings Settings) error {
	e.RotorA.init(settings.InitialPosition.a, settings.RingSetting.a)
	e.RotorB.init(settings.InitialPosition.b, settings.RingSetting.b)
	e.RotorC.init(settings.InitialPosition.c, settings.RingSetting.c)
	e.Reflector.init()

	err := e.plugboard.init(settings.PlugboardConnections)
	if err != nil {
		return errors.Wrap(err, "failed initialising plugboard")
	}

	return nil
}

func (e *Machine) Press(inputChar rune) rune {
	input := rune2Int(inputChar)

	e.rotate()

	outputP := e.plugboard.Pass(input)
	outputA := e.RotorA.Pass(outputP)
	outputB := e.RotorB.Pass(outputA)
	outputC := e.RotorC.Pass(outputB)
	outputR := e.Reflector.Pass(outputC)
	outputC2 := e.RotorC.PassBack(outputR)
	outputB2 := e.RotorB.PassBack(outputC2)
	outputA2 := e.RotorA.PassBack(outputB2)
	outputP2 := e.plugboard.Pass(outputA2)

	// fmt.Printf("rotor encryption: %c->%c->%c->%c->%c->%c->%c->%c->%c\n",
	// 	int2rune(outputP),
	// 	int2rune(outputA),
	// 	int2rune(outputB),
	// 	int2rune(outputC),
	// 	int2rune(outputR),
	// 	int2rune(outputC2),
	// 	int2rune(outputB2),
	// 	int2rune(outputA2),
	// 	int2rune(outputP2))

	return int2rune(outputP2)
}

func (e *Machine) Type(text string) string {
	chars := make([]rune, len(text))
	for i, c := range text {
		chars[i] = e.Press(c)
	}

	return string(chars)
}

func int2rune(i int8) rune {
	return rune(i) + 64
}

func rune2Int(c rune) int8 {
	return int8(c - 64)
}
