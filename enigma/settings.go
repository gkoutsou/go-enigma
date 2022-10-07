package enigma

type RotorSetting struct {
	a, b, c int8
}

var DefaultRotorSetting = RotorSetting{
	a: 1,
	b: 1,
	c: 1,
}

func NewRotorSetting(c, b, a rune) RotorSetting {
	return RotorSetting{
		a: rune2Int(a),
		b: rune2Int(b),
		c: rune2Int(c),
	}
}

type Settings struct {
	PlugboardConnections string
	RingSetting          RotorSetting
	InitialPosition      RotorSetting
}
