package enigma

import (
	"errors"
	"strings"
)

type Plugboard struct {
	mapping map[int8]int8
}

func (r *Plugboard) init(s string) error {
	r.mapping = map[int8]int8{}

	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}

	pairs := strings.Split(s, " ")
	for _, pair := range pairs {
		if len(pair) != 2 {
			return errors.New("should have pairs of characters as input")
		}

		a := rune2Int(rune(pair[0]))
		b := rune2Int(rune(pair[1]))
		r.mapping[a] = b
		r.mapping[b] = a
	}

	return nil
}

func (r *Plugboard) Pass(character int8) int8 {
	if c, ok := r.mapping[character]; ok {
		return c
	}

	return character
}
