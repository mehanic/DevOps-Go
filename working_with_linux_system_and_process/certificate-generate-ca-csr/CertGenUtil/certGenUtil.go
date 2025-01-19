package CertGenUtil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func SignCSR(caPrivateKey *rsa.PrivateKey, caCert *x509.Certificate, csrDER []byte) error {
	// Parse the client's CSR
	clientCSR, err := x509.ParseCertificateRequest(csrDER)
	if err != nil {
		return err
	}

	// Define client's certificate template
	clientCertTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2024),
		Subject:      clientCSR.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0), // 1 year validity
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	// Sign the CSR with the CA's private key to create the client certificate
	clientCertDER, err := x509.CreateCertificate(rand.Reader, clientCertTemplate, caCert, clientCSR.PublicKey, caPrivateKey)
	if err != nil {
		return err
	}

	// Save client's certificate to file
	clientCertFile, err := os.Create("client_certificate.pem")
	if err != nil {
		return err
	}
	defer clientCertFile.Close()

	pem.Encode(clientCertFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCertDER,
	})

	return nil
}
