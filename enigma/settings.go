package enigma

type RotorSetting struct {
	a, b, c int8
}

var DefaultRotorSetting = RotorSetting{
	a: 1,
	b: 1,
	c: 1,
}

func NewRotorSetting(c, b, a int8) RotorSetting {
	return RotorSetting{
		a: a,
		b: b,
		c: c,
	}
}

type Settings struct {
	RingSetting          RotorSetting
	InitialPosition      RotorSetting
	PlugboardConnections string
}
