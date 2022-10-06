package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		rotations int
		c         int8
		b         int8
		a         int8
	}{
		{rotations: 16, c: 1, b: 1, a: 17},
		{rotations: 17, c: 1, b: 2, a: 18},
		{rotations: 18, c: 1, b: 2, a: 19},
		{rotations: 25, c: 1, b: 2, a: 26},
		{rotations: 26, c: 1, b: 2, a: 1},
		{rotations: 18 + 26, c: 1, b: 3, a: 19},
		{rotations: 18 + 26*5, c: 2, b: 7, a: 19},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Rotate=%d", test.rotations), func(t *testing.T) {
			e := Machine{
				RotorA: &Rotor{
					currentPos: 1,
					notchPos:   18,
				},
				RotorB: &Rotor{
					currentPos: 1,
					notchPos:   6,
				},
				RotorC: &Rotor{
					currentPos: 1,
					notchPos:   23,
				},
			}

			for i := 0; i < int(test.rotations); i++ {
				e.rotate()
			}
			require.Equal(t, test.c, e.RotorC.currentPos)
			require.Equal(t, test.b, e.RotorB.currentPos)
			require.Equal(t, test.a, e.RotorA.currentPos)
		})
	}
}

func TestRotateEnd(t *testing.T) {
	e := Machine{
		RotorA: &Rotor{
			currentPos: 26,
			notchPos:   1,
		},
		RotorB: &Rotor{
			currentPos: 26,
			notchPos:   1,
		},
		RotorC: &Rotor{
			currentPos: 26,
			notchPos:   1,
		},
	}

	e.rotate()
	require.Equal(t, int8(1), e.RotorA.currentPos)
	require.Equal(t, int8(1), e.RotorB.currentPos)
	require.Equal(t, int8(1), e.RotorC.currentPos)
}

func TestPressFromAAZ(t *testing.T) {
	cases := []struct {
		input  rune
		output rune
	}{
		{input: 'A', output: 'U'},
		{input: 'B', output: 'E'},
		{input: 'C', output: 'J'},
		{input: 'U', output: 'A'},
		{input: 'Z', output: 'H'},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Char=%c", test.input), func(t *testing.T) {
			settings := Settings{
				RingSetting:     DefaultRotorSetting,
				InitialPosition: NewRotorSetting('A', 'A', 'Z'),
			}
			enigma := Machine{
				RotorA:    &RotorIII,
				RotorB:    &RotorII,
				RotorC:    &RotorI,
				Reflector: ReflectorB,
			}

			enigma.Init(settings)

			require.Equal(t, test.output, enigma.Press(test.input))
		})
	}
}

func TestPress_FromAAA(t *testing.T) {
	input := 'A'
	output := "B"

	settings := Settings{
		RingSetting:     DefaultRotorSetting,
		InitialPosition: DefaultRotorSetting,
	}

	enigma := Machine{
		RotorA:    &RotorIII,
		RotorB:    &RotorII,
		RotorC:    &RotorI,
		Reflector: ReflectorB,
	}

	enigma.Init(settings)

	require.Equal(t, output, string(enigma.Press(input)))
}

func TestPress_WithRingSetting_FromAAA(t *testing.T) {
	input := 'A'
	output := "O"

	settings := Settings{
		RingSetting:     NewRotorSetting('A', 'A', 'B'),
		InitialPosition: NewRotorSetting('A', 'A', 'Z'),
	}

	enigma := Machine{
		RotorA:    &RotorIII,
		RotorB:    &RotorII,
		RotorC:    &RotorI,
		Reflector: ReflectorB,
	}

	enigma.Init(settings)

	require.Equal(t, output, string(enigma.Press(input)))
}
