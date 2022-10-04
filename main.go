package main

import (
	"fmt"

	"github.com/gkoutsou/go-enigma/enigma"
)

func main() {

	enigma := enigma.Machine{
		RotorA:    &enigma.RotorIII,
		RotorB:    &enigma.RotorII,
		RotorC:    &enigma.RotorI,
		Reflector: enigma.ReflectorB,
	}

	enigma.Init(1, 1, 1, "QA ED FG BO LP CS RT UJ HN ZW")

	fmt.Printf("%s", enigma.Type("ABCDE"))
}
