package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hackfeed/ds-7th-sem-labs/lab_02/pkg/enigma"
)

var (
	inputFile, outputFile, checkFile string
)

func init() {
	flag.StringVar(&inputFile, "input", "go.mod", "file to encode")
	flag.StringVar(&outputFile, "output", "go.mod.decoded", "encoding result file")
	flag.StringVar(&checkFile, "check", "go.mod.check", "decoding encoded file result for clearance")
}

func main() {
	flag.Parse()

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	seed := time.Now().UnixNano()
	enigmaMachine := enigma.NewEnigma(seed)
	rotorsCopy := []*enigma.Rotor{
		enigma.NewRotor(seed),
		enigma.NewRotor(seed),
		enigma.NewRotor(seed),
	}

	encoded := enigmaMachine.Encode(data)
	err = os.WriteFile(outputFile, encoded, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	enigmaMachine.Rotors = rotorsCopy
	decoded := enigmaMachine.Encode(encoded)
	err = os.WriteFile(checkFile, decoded, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	isEqual := bytes.Equal(data, decoded)
	fmt.Printf("Input and decoded files are equal: %t\n", isEqual)
}
