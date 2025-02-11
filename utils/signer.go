package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// XypSign is a struct that holds the path to the private key file.
type XypSign struct {
	KeyPath string
}

func (x *XypSign) __GetPrivKey() *rsa.PrivateKey {
	keyData, _ := ioutil.ReadFile(x.KeyPath)
	block, _ := pem.Decode(keyData)
	privKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return privKey
}

func (x *XypSign) __toBeSigned(accessToken, timestamp string) map[string]string {
	return map[string]string{
		"accessToken": accessToken,
		"timeStamp":   timestamp,
	}
}

func (x *XypSign) __buildParam(toBeSigned map[string]string) string {
	fmt.Println(toBeSigned["accessToken"] + "." + toBeSigned["timeStamp"])
	return toBeSigned["accessToken"] + "." + toBeSigned["timeStamp"]
}

func (x *XypSign) Sign(accessToken, timestamp string) (map[string]string, string) {
	toBeSigned := x.__toBeSigned(accessToken, timestamp)
	param := x.__buildParam(toBeSigned)

	hasher := sha256.New()
	hasher.Write([]byte(param))
	digest := hasher.Sum(nil)

	pkey := x.__GetPrivKey()
	signer, _ := rsa.SignPKCS1v15(rand.Reader, pkey, crypto.SHA256, digest)
	signature := base64.StdEncoding.EncodeToString(signer)

	return toBeSigned, signature
}
