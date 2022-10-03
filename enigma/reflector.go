package enigma

type Reflector struct {
	mappingS string
	mapping  map[int8]int8
}

func (r *Reflector) init() {
	r.mapping = map[int8]int8{}

	m := []rune(r.mappingS)
	for i := int8(0); i < int8(26); i++ {
		r.mapping[i+1] = rune2Int(m[i])
	}
}

func (r *Reflector) Pass(character int8) int8 {
	return r.mapping[character]
}

// A: EJMZALYXVBWFCRQUONTSPIKHGD
var ReflectorA = Reflector{
	mappingS: "EJMZALYXVBWFCRQUONTSPIKHGD",
}

// B: YRUHQSLDPXNGOKMIEBFZCWVJAT
var ReflectorB = Reflector{
	mappingS: "YRUHQSLDPXNGOKMIEBFZCWVJAT",
}

// C: FVPJIAOYEDRZXWGCTKUQSBNMHL
var ReflectorC = Reflector{
	mappingS: "FVPJIAOYEDRZXWGCTKUQSBNMHL",
}
