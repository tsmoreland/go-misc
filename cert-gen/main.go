package main

import (
	"bytes"
	"cert-gen/certificatetemplates"
	"cert-gen/exporters"
	"crypto/x509"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Insufficient Arguments")
	}

	exporter, err := exporters.NewFileExporter(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	template, privateKey, err := certificatetemplates.NewLocalhostCertificate(4096)
	if err != nil {
		log.Fatal(err)
	}

	derBytes, privateKey, err := certificatetemplates.CreateCertificateForLocalhost(4096)
	if err != nil {
		log.Fatal(err)
	}
	csrBytes, err := certificatetemplates.CreateSigningRequestForLocalhost(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	out := &bytes.Buffer{}
	if err := exporter.ExportCertificate(derBytes, out); err != nil {
		log.Fatal(err)
	}
	out.Reset()

	if err := exporter.ExportPrivateKey(x509.MarshalPKCS1PrivateKey(privateKey), out); err != nil {
		log.Fatal(err)
	}

	if err := exporter.ExportCertificateSigningRequest(csrBytes, out); err != nil {
		log.Fatal(err)
	}
	out.Reset()

	/*

		requestTemplate := certificatetemplates.NewLocalhostSigningRequest()

		csrBytes, err := x509.CreateCertificateRequest(rand.Reader, requestTemplate, privateKey)
		if err != nil {
			log.Fatal(err)
			return
		}
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
	*/
}
