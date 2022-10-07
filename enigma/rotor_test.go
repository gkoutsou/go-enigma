package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateOnce(t *testing.T) {
	r := Rotor{currentPos: 1, notchPos: 5}

	rotateNext := r.rotate()
	require.False(t, rotateNext)
	require.Equal(t, int8(2), r.currentPos)
}

func TestRotate_WithNotch(t *testing.T) {
	r := Rotor{currentPos: 1, notchPos: 2}

	rotateNext := r.rotate()
	require.True(t, rotateNext)
	require.Equal(t, int8(2), r.currentPos)
}

func TestRotate_Twice(t *testing.T) {
	r := Rotor{currentPos: 1, notchPos: 3}

	rotateNext := r.rotate()
	require.False(t, rotateNext)
	require.Equal(t, int8(2), r.currentPos)

	rotateNext = r.rotate()
	require.True(t, rotateNext)
	require.Equal(t, int8(3), r.currentPos)
}

func TestRotate_26Times(t *testing.T) {
	r := Rotor{currentPos: 1, notchPos: 1}

	for i := 2; i < 27; i++ {
		rotateNext := r.rotate()
		require.Equal(t, int8(i), r.currentPos)
		require.False(t, rotateNext)
	}

	rotateNext := r.rotate()
	require.Equal(t, int8(1), r.currentPos)
	require.True(t, rotateNext)
}

func TestPass_AWhenA(t *testing.T) {
	r := RotorIII
	r.init(2, 1)

	input := rune2Int('A')
	output := r.Pass(input)
	require.Equal(t, "C", string(int2rune(output)))
}

func TestPass_WithRingSetting(t *testing.T) {
	cases := []struct {
		output      string
		input       rune
		initPos     int8
		ringSetting int8
	}{
		{input: 'A', output: "P", initPos: int8(1), ringSetting: int8(2)},
		{input: 'A', output: "B", initPos: int8(2), ringSetting: int8(2)},
		{input: 'T', output: "T", initPos: int8(8), ringSetting: int8(14)},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("(%d,%d) - %c", test.initPos, test.ringSetting, test.input), func(t *testing.T) {
			r := RotorIII
			r.init(test.initPos, test.ringSetting)

			input := rune2Int(test.input)
			output := r.Pass(input)
			require.Equal(t, test.output, string(int2rune(output)))
		})
	}
}
