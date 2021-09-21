package enigma

import (
	"math"
	"math/rand"
)

type Rotor struct {
	Values    []byte
	Rotations int
}

func NewRotor(seed int64) *Rotor {
	return &Rotor{
		Values:    fillRotor(seed),
		Rotations: 0,
	}
}

func fillRotor(seed int64) []byte {
	values := make([]byte, math.MaxUint8+1)
	for i := 0; i <= math.MaxUint8; i++ {
		values[i] = byte(i)
	}

	rand.Seed(seed)
	rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })

	return values
}

func (r *Rotor) Rotate() {
	r.Values = append([]byte{r.Values[math.MaxUint8]}, r.Values[1:]...)
	r.Rotations++
	if r.Rotations == len(r.Values) {
		r.Rotations = 0
	}
}

func (r *Rotor) GetStraight(b byte) byte {
	return r.Values[b]
}

func (r *Rotor) GetReverse(b byte) byte {
	for i, v := range r.Values {
		if v == b {
			return byte(i)
		}
	}
	return 0
}
