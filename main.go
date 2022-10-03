package main

import "fmt"

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

type enigma struct {
	currentPosition *position
	notch           position
}

type position struct {
	c int8 // left most
	b int8
	a int8 // right most
}

func (e *enigma) rotate() {
	p := e.currentPosition
	notch := e.notch

	p.a++
	if p.a > 26 {
		p.a = 1
	}

	if p.a == notch.a {
		p.b++
		if p.b > 26 {
			p.b = 1
		}
	}

	if p.a == notch.a && p.b == notch.b {
		p.c++
		if p.c > 26 {
			p.c = 1
		}
	}
}

func main() {
	rotorI := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	// rotorSetting

	currentPosition := position{
		c: 1,
		b: 1,
		a: 1,
	}

	notchPosition := position{
		a: 18,
		b: 6,
		c: 23,
	}

	enigma := enigma{
		currentPosition: &currentPosition,
		notch:           notchPosition,
	}

	fmt.Println(rotorI)
	fmt.Println(enigma)
}
