package test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gkoutsou/go-enigma/enigma"
	"github.com/stretchr/testify/require"
)

func TestEncode_WithDefaultSettings(t *testing.T) {
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
			settings := enigma.Settings{
				RingSetting:     enigma.DefaultRotorSetting,
				InitialPosition: enigma.DefaultRotorSetting,
			}

			e := enigma.Machine{
				RotorA:    &enigma.RotorIII,
				RotorB:    &enigma.RotorII,
				RotorC:    &enigma.RotorI,
				Reflector: enigma.ReflectorB,
			}

			err := e.Init(settings)
			require.NoError(t, err)

			require.Equal(t, test.output, e.Type(test.input))
		})
	}
}

func TestEncode_WithPlugboard(t *testing.T) {
	input := "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"
	expected := "VAUFLPVWMQIVFWNPCGPGVPIMKUWZREEDTTQ"

	settings := enigma.Settings{
		RingSetting:          enigma.DefaultRotorSetting,
		InitialPosition:      enigma.DefaultRotorSetting,
		PlugboardConnections: "QA ED FG BO LP CS RT UJ HN ZW",
	}
	e := enigma.Machine{
		RotorA:    &enigma.RotorIII,
		RotorB:    &enigma.RotorII,
		RotorC:    &enigma.RotorI,
		Reflector: enigma.ReflectorB,
	}

	err := e.Init(settings)
	require.NoError(t, err)

	require.Equal(t, expected, e.Type(input))
}

func TestEncode_WithRingSetting(t *testing.T) {
	input := "TRYMENOW"
	expected := "BIMFPHIL"

	settings := enigma.Settings{
		RingSetting:     enigma.NewRotorSetting('J', 'L', 'N'),
		InitialPosition: enigma.NewRotorSetting('Q', 'O', 'G'),
	}
	e := enigma.Machine{
		RotorA:    &enigma.RotorIII,
		RotorB:    &enigma.RotorII,
		RotorC:    &enigma.RotorI,
		Reflector: enigma.ReflectorB,
	}

	err := e.Init(settings)
	require.NoError(t, err)

	require.Equal(t, expected, e.Type(input))
}

func TestEncode_LongInput(t *testing.T) {
	input := randomString(1024)

	settings := enigma.Settings{
		RingSetting:          enigma.NewRotorSetting('B', 'C', 'D'),
		InitialPosition:      enigma.NewRotorSetting('E', 'F', 'G'),
		PlugboardConnections: "QA ED FG BO LP CS RT UJ HN ZW",
	}

	e := enigma.Machine{
		RotorA:    &enigma.RotorIII,
		RotorB:    &enigma.RotorII,
		RotorC:    &enigma.RotorI,
		Reflector: enigma.ReflectorB,
	}

	err := e.Init(settings)
	require.NoError(t, err)

	output := e.Type(input)
	require.Len(t, output, 1024)
	require.NotEqual(t, input, output)

	// Pass output again to a initialised machine
	// The final output should be the same as the original message

	err = e.Init(settings)
	require.NoError(t, err)

	finalOutput := e.Type(output)
	require.Len(t, finalOutput, 1024)
	require.Equal(t, input, finalOutput)
}

func randomString(n int) string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
