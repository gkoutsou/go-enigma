package main

import (
	"fmt"

	"github.com/gkoutsou/go-enigma/enigma"
)

func main() {
	settings := enigma.Settings{
		RingSetting:          enigma.DefaultRotorSetting,
		InitialPosition:      enigma.DefaultRotorSetting,
		PlugboardConnections: "QA ED FG BO LP CS RT UJ HN ZW",
	}

	e := enigma.Machine{
		RotorA:    &enigma.RotorIII,
		RotorB:    &enigma.RotorII,
		RotorC:    &enigma.RotorI,
		Reflector: enigma.ReflectorB,
	}

	err := e.Init(settings)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", e.Type("ABCDE"))
}
