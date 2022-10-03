package main

// func TestRotateOnce(t *testing.T) {
// 	p := position{c: 1, b: 1, a: 1}
// 	notchPosition := position{a: 18, b: 6, c: 23}

// 	enigma := enigma{
// 		currentPosition: &p,
// 		notch:           notchPosition,
// 	}

// 	enigma.rotate()
// 	require.Equal(t, int8(1), enigma.currentPosition.c)
// 	require.Equal(t, int8(1), enigma.currentPosition.b)
// 	require.Equal(t, int8(2), enigma.currentPosition.a)
// }

// func TestRotate(t *testing.T) {
// 	cases := []struct {
// 		rotations int8
// 		expected  position
// 	}{
// 		{rotations: 16, expected: position{c: 1, b: 1, a: 17}},
// 		{rotations: 17, expected: position{c: 1, b: 2, a: 18}},
// 		{rotations: 18, expected: position{c: 1, b: 2, a: 19}},
// 		{rotations: 25, expected: position{c: 1, b: 2, a: 26}},
// 		{rotations: 26, expected: position{c: 1, b: 2, a: 1}},
// 		{rotations: 18 + 26, expected: position{c: 1, b: 3, a: 19}},
// 	}

// 	for _, test := range cases {
// 		t.Run(fmt.Sprintf("Rotate=%d", test.rotations), func(t *testing.T) {
// 			p := position{c: 1, b: 1, a: 1}
// 			notchPosition := position{a: 18, b: 6, c: 23}

// 			enigma := enigma{
// 				currentPosition: &p,
// 				notch:           notchPosition,
// 			}

// 			for i := 0; i < int(test.rotations); i++ {
// 				enigma.rotate()
// 			}
// 			require.Equal(t, test.expected.c, enigma.currentPosition.c)
// 			require.Equal(t, test.expected.b, enigma.currentPosition.b)
// 			require.Equal(t, test.expected.a, enigma.currentPosition.a)
// 		})
// 	}
// }
