package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

type XypSign struct {
	KeyPath string
}

type SignatureData struct {
	AccessToken string
	Timestamp   string
	Signature   string
}

func (x *XypSign) toBeSigned(accessToken, timestamp string) SignatureData {
	return SignatureData{
		AccessToken: accessToken,
		Timestamp:   timestamp,
	}
}

func (x *XypSign) Generate(accessToken, timestamp string) (SignatureData, error) {
	// Read private key file
	privateKeyBytes, err := os.ReadFile(x.KeyPath)
	if err != nil {
		return SignatureData{}, fmt.Errorf("failed to read private key: %w", err)
	}

	privateKeyString := string(privateKeyBytes)
	privateKeyString = strings.ReplaceAll(privateKeyString, "\n", "")
	privateKeyString = strings.ReplaceAll(privateKeyString, "-----BEGIN PRIVATE KEY-----", "")
	privateKeyString = strings.ReplaceAll(privateKeyString, "-----END PRIVATE KEY-----", "")

	decodedKey, err := base64.StdEncoding.DecodeString(privateKeyString)
	if err != nil {
		return SignatureData{}, fmt.Errorf("failed to decode private key: %w", err)
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(decodedKey)
	if err != nil {
		return SignatureData{}, fmt.Errorf("failed to parse private key: %w", err)
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return SignatureData{}, fmt.Errorf("private key is not RSA")
	}

	data := x.toBeSigned(accessToken, timestamp)
	message := []byte(data.AccessToken + "." + data.Timestamp)

	// Create hash
	hash := sha256.Sum256(message)

	// Sign the hash
	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		return SignatureData{}, fmt.Errorf("failed to create signature: %w", err)
	}

	// Encode the signature
	data.Signature = base64.StdEncoding.EncodeToString(signature)

	return data, nil
}
