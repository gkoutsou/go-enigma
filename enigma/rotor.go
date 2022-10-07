package enigma

type Rotor struct {
	mapping     map[int8]int8
	mappingBack map[int8]int8
	mappingS    string

	currentPos  int8
	notchPos    int8
	ringSetting int8
}

func (r *Rotor) rotate() bool {
	r.currentPos++
	if r.currentPos > 26 {
		r.currentPos = 1
	}

	if r.currentPos == r.notchPos {
		return true
	}

	return false
}

func (r *Rotor) init(initialPosition, ringSetting int8) {
	r.mapping = map[int8]int8{}
	r.mappingBack = map[int8]int8{}

	m := []rune(r.mappingS)
	for i := int8(0); i < int8(26); i++ {
		r.mapping[i+1] = rune2Int(m[i])
		r.mappingBack[rune2Int(m[i])] = i + 1
	}

	r.currentPos = initialPosition
	r.ringSetting = ringSetting
}

func (r *Rotor) Pass(character int8) int8 {
	position := character + (r.currentPos - 1) - (r.ringSetting - 1) + 26
	for position > 26 {
		position -= 26
	}

	output := r.mapping[position] - (r.currentPos - 1) + (r.ringSetting - 1) - 26
	for output <= 0 {
		output += 26
	}

	return output
}

func (r *Rotor) PassBack(character int8) int8 {
	position := character + (r.currentPos - 1) - (r.ringSetting - 1) + 26
	for position > 26 {
		position -= 26
	}

	output := r.mappingBack[position] - (r.currentPos - 1) + (r.ringSetting - 1) - 26
	for output <= 0 {
		output += 26
	}

	return output
}

// RotorI contains the following mapping EKMFLGDQVZNTOWYHXUSPAIBRCJ. The notch is at position R
var RotorI = Rotor{
	mappingS: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	notchPos: 18,
}

// RotorII contains the following mapping AJDKSIRUXBLHWTMCQGZNPYFVOE. The notch is at position F
var RotorII = Rotor{
	mappingS: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	notchPos: 6,
}

// RotorIII contains the following mapping BDFHJLCPRTXVZNYEIWGAKMUSQO. The notch is at position W
var RotorIII = Rotor{
	mappingS: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	notchPos: 23,
}

// RotorIV contains the following mapping ESOVPZJAYQUIRHXLNFTGKDCMWB. The notch is at position K
var RotorIV = Rotor{
	mappingS: "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	notchPos: 11,
}

// RotorV contains the following mapping VZBRGITYUPSDNHLXAWMJQOFECK. The notch is at position A
var RotorV = Rotor{
	mappingS: "VZBRGITYUPSDNHLXAWMJQOFECK",
	notchPos: 1,
}
