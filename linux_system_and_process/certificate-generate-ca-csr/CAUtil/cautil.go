package CAUtil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func CreateCACertificate() (*rsa.PrivateKey, *x509.Certificate, error) {
	// Generate CA's private key
	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Define CA's certificate template
	caCertTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2024),
		Subject: pkix.Name{
			Country:            []string{"US"},
			Province:           []string{"California"},
			Locality:           []string{"San Francisco"},
			Organization:       []string{"My CA Company"},
			OrganizationalUnit: []string{"Certificate Authority"},
			CommonName:         "My CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // 1 year validity
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		IsCA:                  true,
		BasicConstraintsValid: true,
	}

	// Create CA's self-signed certificate
	caCertDER, err := x509.CreateCertificate(rand.Reader, caCertTemplate, caCertTemplate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	// Save CA's private key to file
	caPrivateKeyFile, err := os.Create("ca_private_key.pem")
	if err != nil {
		return nil, nil, err
	}
	defer caPrivateKeyFile.Close()

	pem.Encode(caPrivateKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey),
	})

	// Save CA's self-signed certificate to file
	caCertFile, err := os.Create("ca_certificate.pem")
	if err != nil {
		return nil, nil, err
	}
	defer caCertFile.Close()

	pem.Encode(caCertFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCertDER,
	})

	return caPrivateKey, caCertTemplate, nil
}
