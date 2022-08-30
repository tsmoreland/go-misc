package exporters

import (
	"bytes"
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

	buffer.Reset()
	return nil
}

func (e FileExporter) ExportPrivateKey(source []byte, buffer *bytes.Buffer) error {
	buffer.Reset()
	return nil
}

func (e FileExporter) ExportCertificateSigningRequest(source []byte, buffer *bytes.Buffer) error {

	buffer.Reset()
	return nil
}
