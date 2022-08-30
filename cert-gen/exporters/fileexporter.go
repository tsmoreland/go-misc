package exporters

import (
	"bytes"
	"encoding/pem"
	"os"
	"path/filepath"
)

type FileExporter struct {
	certificateFile        string
	certificateKeyFile     string
	certificateRequestFile string
}

func NewFileExporter(folder string) (*FileExporter, error) {
	_, err := os.Stat(folder)
	if err != nil {
		return nil, err
	}

	certificateFile := filepath.Join(folder, "server.cer")
	certificateKeyFile := filepath.Join(folder, "server.key")
	certificateRequestFile := filepath.Join(folder, "server.csr")

	return &FileExporter{
		certificateFile:        certificateFile,
		certificateKeyFile:     certificateKeyFile,
		certificateRequestFile: certificateRequestFile,
	}, nil
}

func (e FileExporter) ExportCertificate(source []byte, buffer *bytes.Buffer) error {
	defer buffer.Reset()
	return writePemFile(e.certificateFile, buffer, "CERTIFICATE", source)
}

func (e FileExporter) ExportPrivateKey(source []byte, buffer *bytes.Buffer) error {
	defer buffer.Reset()
	return writePemFile(e.certificateKeyFile, buffer, "RSA PRIVATE KEY", source)
}

func (e FileExporter) ExportCertificateSigningRequest(source []byte, buffer *bytes.Buffer) error {
	defer buffer.Reset()
	return writePemFile(e.certificateRequestFile, buffer, "CERTIFICATE REQUEST", source)
}

func writePemFile(filename string, buffer *bytes.Buffer, pemLabel string, source []byte) error {
	if err := pem.Encode(buffer, &pem.Block{Type: pemLabel, Bytes: source}); err != nil {
		return err
	}

	return os.WriteFile(filename, buffer.Bytes(), 0644)
}
