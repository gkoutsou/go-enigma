package enigma

type Reflector struct {
	mapping  map[int8]int8
	mappingS string
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

// ReflectorA contains the following mapping EJMZALYXVBWFCRQUONTSPIKHGD
var ReflectorA = Reflector{
	mappingS: "EJMZALYXVBWFCRQUONTSPIKHGD",
}

// ReflectorB contains the following mapping YRUHQSLDPXNGOKMIEBFZCWVJAT
var ReflectorB = Reflector{
	mappingS: "YRUHQSLDPXNGOKMIEBFZCWVJAT",
}

// ReflectorC contains the following mapping FVPJIAOYEDRZXWGCTKUQSBNMHL
var ReflectorC = Reflector{
	mappingS: "FVPJIAOYEDRZXWGCTKUQSBNMHL",
}
