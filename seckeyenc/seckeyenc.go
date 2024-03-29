package seckeyenc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
)

var (
	pubKey      *rsa.PublicKey
	priKey      *rsa.PrivateKey
	dummyReader = &DummyReader{}
)

type DummyReader struct{}

func (d *DummyReader) Read(p []byte) (int, error) {
	return 1, nil
}

func init() {
	b, err := ioutil.ReadFile("seckeyenc/sample.der")
	if err != nil {
		log.Fatal(err)
	}
	p, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		log.Fatal("x509.ParsePKIXPublicKey: ", err)
	}
	pubKey = p.(*rsa.PublicKey)

	b, err = ioutil.ReadFile("seckeyenc/sample_priv.der")
	if err != nil {
		log.Fatal(err)
	}
	p, err = x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		log.Fatal("x509.ParsePKCS1PrivateKey: ", err)
	}
	priKey = p.(*rsa.PrivateKey)
}

func Encrypt(msg []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, msg)
}

func EncryptBase64URLString(msg []byte) (string, error) {
	b, err := Encrypt(msg)
	if err != nil {
		return "", err
	}

	return url.QueryEscape(base64.StdEncoding.EncodeToString(b)), nil
}

func Decrypt(msg []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, priKey, msg)
}

func DecryptBase64URLString(s string) ([]byte, error) {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return nil, err
	}

	msg, err := base64.StdEncoding.DecodeString(decoded)
	if err != nil {
		return nil, err
	}

	return Decrypt(msg)
}

func Hello(value string) {
	msg := []byte(value)
	b, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey, msg)
	enc := url.QueryEscape(base64.StdEncoding.EncodeToString(b))

	b, _ = rsa.EncryptPKCS1v15(dummyReader, pubKey, msg)
	enc2 := url.QueryEscape(base64.StdEncoding.EncodeToString(b))

	fmt.Println("rand.Reader:", enc)
	fmt.Println("dummyReader:", enc2)

	esc, _ := url.QueryUnescape(enc)
	msg, _ = base64.StdEncoding.DecodeString(esc)

	dec, _ := rsa.DecryptPKCS1v15(rand.Reader, priKey, msg)
	dec2, _ := rsa.DecryptPKCS1v15(dummyReader, priKey, msg)

	fmt.Println("rand.Reader: decode:", string(dec))
	fmt.Println("dummyReader: decode:", string(dec2))
}
