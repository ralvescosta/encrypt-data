package crypto

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
)

var pemDecode = pem.Decode
var parsePKCS1PubKey = x509.ParsePKCS1PublicKey
var jsonMarshal = json.Marshal
var jweEncrypt = jwe.Encrypt

func Encrypt(pubKey string, data map[string]interface{}) ([]byte, error) {
	log.Printf("[Crypto::Encrypt] encrypting...")

	block, _ := pemDecode([]byte(pubKey))
	pKey, err := parsePKCS1PubKey(block.Bytes)
	if err != nil {
		log.Println("[Crypto::Encrypt] error whiling parse public key")
		return nil, err
	}

	dataToByte, err := jsonMarshal(data)
	if err != nil {
		log.Printf("[Crypto::Encrypt] Marshal Error: %s\n", err.Error())
		return nil, errors.New("JSON parse error")
	}

	encrypted, err := jweEncrypt(dataToByte, jwa.RSA_OAEP_256, pKey, jwa.A256CBC_HS512, jwa.NoCompress)
	if err != nil {
		log.Printf("[Crypto::Encrypt] JWE Encrypt Error: %s\n", err.Error())
		return nil, errors.New("encrypt error")
	}

	log.Println("[Crypto::Encrypt] encrypted")

	return encrypted, nil
}
