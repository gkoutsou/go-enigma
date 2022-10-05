package enigma

import (
	"fmt"
	"math/rand"
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
				RingSetting:     NewRotorSetting(1, 1, 1),
				InitialPosition: NewRotorSetting(26, 1, 1),
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

func TestPressFromAAA(t *testing.T) {
	cases := []struct {
		input  rune
		output rune
	}{
		{input: 'A', output: 'B'},
		{input: 'A', output: 'B'},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%c->%c", test.input, test.output), func(t *testing.T) {
			settings := Settings{
				RingSetting:     NewRotorSetting(1, 1, 1),
				InitialPosition: NewRotorSetting(1, 1, 1),
			}

			enigma := Machine{
				RotorA:    &RotorIII,
				RotorB:    &RotorII,
				RotorC:    &RotorI,
				Reflector: ReflectorB,
			}

			enigma.Init(settings)

			require.Equal(t, string(test.output), string(enigma.Press(test.input)))
		})
	}
}

func TestTypeFromAAA(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{input: "ABCDE", output: "BJELR"},
		{input: "THEQUICKBROWNFOX", output: "OPCILLAZFXLQTDNL"},
		{input: "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG", output: "OPCILLAZFXLQTDNLGGLEKDIZOKQKGXIEZKD"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s->%s", test.input, test.output), func(t *testing.T) {
			settings := Settings{
				RingSetting:     NewRotorSetting(1, 1, 1),
				InitialPosition: NewRotorSetting(1, 1, 1),
			}

			enigma := Machine{
				RotorA:    &RotorIII,
				RotorB:    &RotorII,
				RotorC:    &RotorI,
				Reflector: ReflectorB,
			}

			enigma.Init(settings)

			require.Equal(t, test.output, enigma.Type(test.input))
		})
	}
}

func TestTypeWithPlugboard(t *testing.T) {
	input := "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"
	expected := "VAUFLPVWMQIVFWNPCGPGVPIMKUWZREEDTTQ"

	settings := Settings{
		RingSetting:          NewRotorSetting(1, 1, 1),
		InitialPosition:      NewRotorSetting(1, 1, 1),
		PlugboardConnections: "QA ED FG BO LP CS RT UJ HN ZW",
	}
	enigma := Machine{
		RotorA:    &RotorIII,
		RotorB:    &RotorII,
		RotorC:    &RotorI,
		Reflector: ReflectorB,
	}

	enigma.Init(settings)

	require.Equal(t, expected, enigma.Type(input))
}

func TestLongInput(t *testing.T) {
	input := RandomString(1024)

	settings := Settings{
		RingSetting:          NewRotorSetting(2, 3, 4),
		InitialPosition:      NewRotorSetting(5, 6, 7),
		PlugboardConnections: "QA ED FG BO LP CS RT UJ HN ZW",
	}

	enigma := Machine{
		RotorA:    &RotorIII,
		RotorB:    &RotorII,
		RotorC:    &RotorI,
		Reflector: ReflectorB,
	}

	enigma.Init(settings)
	output := enigma.Type(input)
	require.Len(t, output, 1024)
	require.NotEqual(t, input, output)

	// Pass output again to a initialised machine
	// The final output should be the same as the original message

	enigma.Init(settings)
	finalOutput := enigma.Type(output)
	require.Len(t, finalOutput, 1024)
	require.Equal(t, input, finalOutput)
}

func RandomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
