package ClientUtil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"os"
)

func CreateClientCSR() (*rsa.PrivateKey, []byte, error) {
	// Generate client's private key
	clientPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Define client's CSR template
	clientCSRTemplate := &x509.CertificateRequest{
		Subject: pkix.Name{
			Country:            []string{"US"},
			Province:           []string{"California"},
			Locality:           []string{"San Francisco"},
			Organization:       []string{"My Client Company"},
			OrganizationalUnit: []string{"IT Department"},
			CommonName:         "client.com",
		},
	}

	// Create the CSR
	clientCSRDER, err := x509.CreateCertificateRequest(rand.Reader, clientCSRTemplate, clientPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	// Save client's private key to file
	clientPrivateKeyFile, err := os.Create("client_private_key.pem")
	if err != nil {
		return nil, nil, err
	}
	defer clientPrivateKeyFile.Close()

	pem.Encode(clientPrivateKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(clientPrivateKey),
	})

	// Save client's CSR to file
	clientCSRFile, err := os.Create("client_csr.pem")
	if err != nil {
		return nil, nil, err
	}
	defer clientCSRFile.Close()

	pem.Encode(clientCSRFile, &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: clientCSRDER,
	})

	return clientPrivateKey, clientCSRDER, nil
}
