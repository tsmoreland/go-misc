package certificatetemplates

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"time"
)

var (
	localhostSubject = &pkix.Name{CommonName: "localhost"}
)

const (
	defaultValidFor = time.Hour * 24 * 365
)

func NewLocalhostCertificateWithDuration(bitSize int, validFor time.Duration) (*x509.Certificate, *rsa.PrivateKey, error) {
	if bitSize < 1024 {
		return nil, nil, fmt.Errorf("%d cannot be less than 1024", bitSize)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, nil, err
	}

	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               *localhostSubject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(validFor),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageTimeStamping},
		BasicConstraintsValid: true,
	}

	return template, privateKey, nil
}

func NewLocalhostCertificate(bitSize int) (*x509.Certificate, *rsa.PrivateKey, error) {
	return NewLocalhostCertificateWithDuration(bitSize, defaultValidFor)
}

func CreateCertificateForLocalhost(bitSize int) ([]byte, *rsa.PrivateKey, error) {
	return CreateCertificateForLocalhostWithDuration(bitSize, defaultValidFor)
}

func CreateCertificateForLocalhostWithDuration(bitSize int, validFor time.Duration) ([]byte, *rsa.PrivateKey, error) {
	template, privateKey, err := NewLocalhostCertificateWithDuration(bitSize, validFor)
	if err != nil {
		return nil, nil, err
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}
	return derBytes, privateKey, err
}

func CreateSigningRequestForLocalhost(privateKey *rsa.PrivateKey) ([]byte, error) {
	requestTemplate := &x509.CertificateRequest{
		Subject:            *localhostSubject,
		SignatureAlgorithm: x509.SHA512WithRSA,
	}
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, requestTemplate, privateKey)
	if err != nil {
		return nil, err
	}
	return csrBytes, nil
}
