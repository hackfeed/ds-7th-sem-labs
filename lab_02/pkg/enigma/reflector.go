package enigma

import (
	"math"
)

type Reflector struct {
	Values []byte
}

func NewReflector() *Reflector {
	return &Reflector{
		Values: fillReflector(),
	}
}

func fillReflector() []byte {
	values := make([]byte, math.MaxUint8+1)
	for i := 0; i <= math.MaxUint8; i++ {
		values[i] = byte(math.MaxUint8 - i)
	}
	return values
}
