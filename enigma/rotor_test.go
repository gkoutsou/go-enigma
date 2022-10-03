package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateOnce(t *testing.T) {
	r := Rotor{currentPos: 1, notchPos: 5}

	rotateNext := r.rotate()
	require.False(t, rotateNext)
	require.Equal(t, int8(2), r.currentPos)
}

func TestRotateWithNotch(t *testing.T) {
	r := Rotor{currentPos: 1, notchPos: 2}

	rotateNext := r.rotate()
	require.True(t, rotateNext)
	require.Equal(t, int8(2), r.currentPos)
}

func TestRotateTwice(t *testing.T) {
	r := Rotor{currentPos: 1, notchPos: 3}

	rotateNext := r.rotate()
	require.False(t, rotateNext)
	require.Equal(t, int8(2), r.currentPos)

	rotateNext = r.rotate()
	require.True(t, rotateNext)
	require.Equal(t, int8(3), r.currentPos)
}

func TestRotate26(t *testing.T) {
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

func TestPassA_WhenA(t *testing.T) {
	r := RotorIII
	r.init(2)

	input := rune2Int('A')
	output := r.Pass(input)
	require.Equal(t, "C", string(int2rune(output)))

}
