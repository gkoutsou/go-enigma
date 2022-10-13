# go-egigma

A golang implementation of the Enigma Machine. 

Can be customised with the 5 different settings: 
1. Ring order - the order of the 3 rotors
1. Initial position - the starting position of the rotors
1. Ring setting - the positioning of the wiring compared to the rotors themselves
1. Reflector - the reflector used to feed the output of the last rotor as an input of the same rotor 
1. Plugboard Setting - the optional extra letter substitutions. 

This library implements 5 Rotors (I-V) and 3 Reflectors (A-C), as described in [Wikipedia](https://en.wikipedia.org/wiki/Enigma_rotor_details).


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
	if err != nil {
        panic("initialisation error")
    }

	output := e.Type("TEXTTOENCODE")