package enigma

type Rotor struct {
	mappingS    string
	mapping     map[int8]int8
	mappingBack map[int8]int8

	currentPos int8
	notchPos   int8
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

func (r *Rotor) init(pos int8) {
	r.mapping = map[int8]int8{}
	r.mappingBack = map[int8]int8{}

	m := []rune(r.mappingS)
	for i := int8(0); i < int8(26); i++ {
		r.mapping[i+1] = rune2Int(m[i])
		r.mappingBack[rune2Int(m[i])] = i + 1
	}

	r.currentPos = pos
}

func (r *Rotor) Pass(character int8) int8 {
	position := character + (r.currentPos - 1)
	if position > 26 {
		position -= 26
	}

	output := r.mapping[position] - (r.currentPos - 1)
	if output <= 0 {
		output += 26
	}

	return output
}

func (r *Rotor) PassBack(character int8) int8 {
	position := character + (r.currentPos - 1)
	if position > 26 {
		position -= 26
	}

	output := r.mappingBack[position] - (r.currentPos - 1)
	if output <= 0 {
		output += 26
	}

	return output
}

// I: EKMFLGDQVZNTOWYHXUSPAIBRCJ notch: R
var RotorI = Rotor{
	mappingS: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	notchPos: 18,
}

// II: AJDKSIRUXBLHWTMCQGZNPYFVOE notch: F
var RotorII = Rotor{
	mappingS: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	notchPos: 6,
}

// III: BDFHJLCPRTXVZNYEIWGAKMUSQO notch: W
var RotorIII = Rotor{
	mappingS: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	notchPos: 23,
}

// TODO:
// IV: ESOVPZJAYQUIRHXLNFTGKDCMWB notch: K
// V: VZBRGITYUPSDNHLXAWMJQOFECK notch: A
