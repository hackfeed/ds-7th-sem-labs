package main

import (
	"flag"
	"log"
	"os"

	"github.com/hackfeed/ds-7th-sem-labs/lab_03/pkg/des"
)

var (
	file, output, key    string
	toEncrypt, toDecrypt bool
)

func init() {
	flag.BoolVar(&toEncrypt, "e", false, "encrypt file")
	flag.BoolVar(&toDecrypt, "d", false, "decrypt file")
	flag.StringVar(&file, "f", "Makefile", "file to perform operation on")
	flag.StringVar(&output, "o", "desed", "file to store result")
	flag.StringVar(&key, "k", "password", "encryption/decryption key")
}

func main() {
	flag.Parse()

	if toEncrypt && toDecrypt {
		log.Fatalln("Can't encrypt and decrypt file at the same time")
	} else {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Can't open file, error is: %s", err)
		}

		key = des.CompleteKey(key)
		if toEncrypt {
			encrypted := des.Encrypt(data, des.GenerateKeys(key))
			err := os.WriteFile(output, encrypted, 0644)
			if err != nil {
				log.Fatalf("Can't write encrypted data, error is: %s", err)
			}
		} else {
			decrypted := des.Decrypt(data, des.GenerateKeys(key))
			err := os.WriteFile(output, decrypted, 0644)
			if err != nil {
				log.Fatalf("Can't write decrypted data, error is: %s", err)
			}
		}
	}
}
