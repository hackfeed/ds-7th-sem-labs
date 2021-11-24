package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hackfeed/ds-7th-sem-labs/lab_05/pkg/digsig"
)

var (
	pubKey, file, sig string
)

func init() {
	flag.StringVar(&pubKey, "pub", "pubkey.pem", "public key filename")
	flag.StringVar(&file, "file", "test.txt", "file to verify")
	flag.StringVar(&sig, "sig", "signature.sig", "signature file")
}

func main() {
	flag.Parse()
	isValid, err := digsig.Verify(pubKey, file, sig)
	if err != nil {
		log.Fatalf("Failed to sign file, error is: %s", err)
	}
	if !isValid {
		fmt.Println("Signature is corrupted")
	} else {
		fmt.Println("Signature is correct")
	}
}
