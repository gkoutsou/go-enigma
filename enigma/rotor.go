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
	return r.mapping[character+r.currentPos-1]
}

func (r *Rotor) PassBack(character int8) int8 {
	return r.mappingBack[character+r.currentPos-1]
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
