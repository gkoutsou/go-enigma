package main

import (
	"fmt"

	"github.com/gkoutsou/go-enigma/enigma"
)

// rotors
// I: EKMFLGDQVZNTOWYHXUSPAIBRCJ notch: R
// II: AJDKSIRUXBLHWTMCQGZNPYFVOE notch: F
// III: BDFHJLCPRTXVZNYEIWGAKMUSQO notch: W
// IV: ESOVPZJAYQUIRHXLNFTGKDCMWB notch: K
// V: VZBRGITYUPSDNHLXAWMJQOFECK notch: A

//reflectors
// A: EJMZALYXVBWFCRQUONTSPIKHGD
// B: YRUHQSLDPXNGOKMIEBFZCWVJAT
// C: FVPJIAOYEDRZXWGCTKUQSBNMHL

func main() {
	// rotorI := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"

	// enigma := enigma.Machine{
	// 	RotorA: &enigma.RotorI,
	// 	RotorB: &enigma.RotorII,
	// 	RotorC: &enigma.RotorIII,
	// }

	// enigma.Init(1, 1, 1)

	// A => H
	// B => I
	// C => L

	enigma := enigma.Machine{
		RotorA:    &enigma.RotorIII,
		RotorB:    &enigma.RotorII,
		RotorC:    &enigma.RotorI,
		Reflector: enigma.ReflectorB,
	}

	enigma.Init(1, 1, 1)

	fmt.Printf("%s", enigma.Type("ABCDE"))
}
