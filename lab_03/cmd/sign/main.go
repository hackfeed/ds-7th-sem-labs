package main

import (
	"flag"
	"log"

	"github.com/hackfeed/ds-7th-sem-labs/lab_03/pkg/digsig"
)

var (
	privKey, file, sig string
)

func init() {
	flag.StringVar(&privKey, "priv", "privkey.pem", "private key filename")
	flag.StringVar(&file, "file", "test.txt", "file to sign")
	flag.StringVar(&sig, "sig", "signature.sig", "file to save signature")
}

func main() {
	flag.Parse()
	if err := digsig.Sign(privKey, file, sig); err != nil {
		log.Fatalf("Failed to sign file, error is: %s", err)
	}
}
