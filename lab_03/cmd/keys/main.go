package main

import (
	"flag"
	"log"

	"github.com/hackfeed/ds-7th-sem-labs/lab_03/pkg/digsig"
)

var (
	privKey, pubKey string
)

func init() {
	flag.StringVar(&privKey, "priv", "privkey.pem", "private key filename")
	flag.StringVar(&pubKey, "pub", "pubkey.pem", "public key filename")
}

func main() {
	flag.Parse()
	if err := digsig.GenerateKeys(privKey, pubKey); err != nil {
		log.Fatalf("Failed to generate keys, error is: %s", err)
	}
}
