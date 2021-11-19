package digsig

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateKeys(privKey, pubKey string) error {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	privateKeyPEM, err := os.Create(privKey)
	if err != nil {
		return err
	}
	defer privateKeyPEM.Close()
	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return err
	}
	if err := pem.Encode(privateKeyPEM, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}); err != nil {
		return err
	}

	publicKeyPEM, err := os.Create(pubKey)
	if err != nil {
		return err
	}
	defer publicKeyPEM.Close()
	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	if err := pem.Encode(publicKeyPEM, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}); err != nil {
		return err
	}

	return nil
}

func Sign(privKey, file, sig string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	key, err := os.ReadFile(privKey)
	if err != nil {
		return err
	}

	pem, _ := pem.Decode(key)
	privKeyParsed, err := x509.ParsePKCS8PrivateKey(pem.Bytes)
	if err != nil {
		return err
	}

	return os.WriteFile(sig, ed25519.Sign(privKeyParsed.(ed25519.PrivateKey), data), 0644)
}

func Verify(pubKey, file, sig string) (bool, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return false, err
	}
	key, err := os.ReadFile(pubKey)
	if err != nil {
		return false, err
	}
	signature, err := os.ReadFile(sig)
	if err != nil {
		return false, err
	}

	pem, _ := pem.Decode(key)
	pubKeyParsed, err := x509.ParsePKIXPublicKey(pem.Bytes)
	if err != nil {
		return false, err
	}

	return ed25519.Verify(pubKeyParsed.(ed25519.PublicKey), data, signature), nil
}
