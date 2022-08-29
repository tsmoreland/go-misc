package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
		return

	}

	subject := &pkix.Name{CommonName: "localhost"}
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               *subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	requestTemplate := &x509.CertificateRequest{
		Subject:            *subject,
		SignatureAlgorithm: x509.SHA512WithRSA,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, requestTemplate, privateKey)
	if err != nil {
		log.Fatal(err)
		return
	}
	out := &bytes.Buffer{}
	pem.Encode(out, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	fmt.Println(out.String())
	out.Reset()

	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatal(err)
		return
	}
	pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	fmt.Println(out.String())
	out.Reset()

	pem.Encode(out, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	fmt.Println(out.String())
	out.Reset()
}
