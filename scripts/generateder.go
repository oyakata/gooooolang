package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func main() {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	b, err := x509.MarshalPKIXPublicKey(&key.PublicKey)

	if err != nil {
		log.Fatal(err)
	}

	priB := x509.MarshalPKCS1PrivateKey(key)

	pem.Encode(os.Stdout, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	})
	pem.Encode(os.Stdout, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: priB,
	})

	fp, err := os.OpenFile("seckeyenc/sample.der", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	fp.Write(b)

	fp2, err := os.OpenFile("seckeyenc/sample_priv.der", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fp2.Close()
	fp2.Write(priB)
}
