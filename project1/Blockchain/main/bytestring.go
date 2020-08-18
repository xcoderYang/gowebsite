package main

import (
	"bytes"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

var (
	user_cert []byte = []byte("-----BEGIN CERTIFICATE-----\nMIICjjCCAjWgAwIBAgIUHyimB+Aq8gRew9hth+/Fwwj14MswCgYIKoZIzj0EAwIw\nczELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNh\nbiBGcmFuY2lzY28xGTAXBgNVBAoTEG9yZzIuZXhhbXBsZS5jb20xHDAaBgNVBAMT\nE2NhLm9yZzIuZXhhbXBsZS5jb20wHhcNMTgxMDI5MDE1NzAwWhcNMTkxMDI5MDIw\nMjAwWjBCMTAwDQYDVQQLEwZjbGllbnQwCwYDVQQLEwRvcmcxMBIGA1UECxMLZGVw\nYXJ0bWVudDExDjAMBgNVBAMTBXVzZXIyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcD\nQgAEm1tTOAvKaGIJX0L9pAghEpqiT5rsYm3ecXvD37b1kHg38gaAURdMQChNk+aV\niv306DlJ3YL3jmUE5rg3CyznT6OB1zCB1DAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0T\nAQH/BAIwADAdBgNVHQ4EFgQUoGHlhAdH35xZzmiZmWtEvmw22oswKwYDVR0jBCQw\nIoAgq4oobt8b2/AWM8hELsJ5+T+e+RRmkKg0J+umla+7EDwwaAYIKgMEBQYHCAEE\nXHsiYXR0cnMiOnsiaGYuQWZmaWxpYXRpb24iOiJvcmcxLmRlcGFydG1lbnQxIiwi\naGYuRW5yb2xsbWVudElEIjoidXNlcjIiLCJoZi5UeXBlIjoiY2xpZW50In19MAoG\nCCqGSM49BAMCA0cAMEQCIHC2lVAO0WKHjOjQi9g8x6Hj7UKAmpxINuKxeV+Ph1av\nAiAa8NLD6v8FgnhC8wQsR6LGL3jlScBOBjQ5Q3oFuNM7nQ==\n-----END CERTIFICATE-----\n")
)

func main() {
	err := ioutil.WriteFile("user.cert", user_cert, 0644)
	if err != nil {
		fmt.Println("Failed to write user_cert to file")
	}

	display("user.cert")

	return
}

func display(user_cert string) {
	json_file, err := os.Open(user_cert)
	if err != nil {
		fmt.Println("Wrong")
	}
	defer json_file.Close()

	bts, _ := ioutil.ReadAll(json_file)

	cert_begin := bytes.IndexAny(bts, "-----BEGIN")
	if cert_begin == -1 {
		fmt.Println("No cert is found")
	}
	cert_txt := bts[cert_begin:]
	bl, _ := pem.Decode(cert_txt)

	if bl == nil {
		fmt.Println("Not able to decode the PEM structure")
		return
	}

	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		fmt.Println("Failed to parse certificate")
		return
	}

	// address
	hfc := sha256.New()
	hfc.Write(cert.RawSubjectPublicKeyInfo)
	digest := hfc.Sum(nil)
	fmt.Println(reflect.TypeOf(digest))
	hfc.Write(digest)
	digest = hfc.Sum(nil)
	dgst_64 := base64.StdEncoding.EncodeToString(digest)
	fmt.Println(reflect.TypeOf(dgst_64))
	fmt.Println("stakeholder address: " + dgst_64)

	// public key
	pub, err := x509.ParsePKIXPublicKey(cert.RawSubjectPublicKeyInfo)
	if err != nil {
		fmt.Println("failed to parse DER encoded public key: " + err.Error())
		return
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		fmt.Println("pub is of type RSA:", pub)
	case *dsa.PublicKey:
		fmt.Println("pub is of type DSA:", pub)
	case *ecdsa.PublicKey:
		fmt.Println("pub is of type ECDSA:", pub)
	default:
		fmt.Println("unknown type of public key")
		return
	}

	// user name
	uname := cert.Subject.CommonName
	fmt.Println("Common name is: " + uname)
	fmt.Println("Certification is done")
}

func imain() {
	str := "Hello, world"
	bts := []byte(str)
	str2 := string(bts[:])
	fmt.Println(str == str2)

	return
}
