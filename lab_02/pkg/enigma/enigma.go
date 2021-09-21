package enigma

type Enigma struct {
	Rotors []*Rotor
	*Reflector
}

func NewEnigma(seed int64) *Enigma {
	return &Enigma{
		Rotors: []*Rotor{
			NewRotor(seed),
			NewRotor(seed),
			NewRotor(seed),
		},
		Reflector: NewReflector(),
	}
}

func (e *Enigma) Encode(data []byte) []byte {
	encoded := make([]byte, 0)

	for _, b := range data {
		encoded = append(encoded, e.encodeByte(b))
	}

	return encoded
}

func (e *Enigma) encodeByte(data byte) byte {
	for i := range e.Rotors {
		data = e.Rotors[i].GetStraight(data)
	}
	data = e.Reflector.Values[data]
	for i := range e.Rotors {
		data = e.Rotors[len(e.Rotors)-i-1].GetReverse(data)
	}

	for _, rotor := range e.Rotors {
		rotor.Rotate()
		if rotor.Rotations != 0 {
			break
		}
	}

	return data
}
