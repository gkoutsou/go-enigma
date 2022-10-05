package main

import (
	"fmt"

	"github.com/gkoutsou/go-enigma/enigma"
)

func main() {

	settings := enigma.Settings{
		RingSetting:          enigma.NewRotorSetting(1, 1, 1),
		InitialPosition:      enigma.NewRotorSetting(1, 1, 1),
		PlugboardConnections: "QA ED FG BO LP CS RT UJ HN ZW",
	}

	enigma := enigma.Machine{
		RotorA:    &enigma.RotorIII,
		RotorB:    &enigma.RotorII,
		RotorC:    &enigma.RotorI,
		Reflector: enigma.ReflectorB,
	}

	enigma.Init(settings)

	fmt.Printf("%s", enigma.Type("ABCDE"))
}
