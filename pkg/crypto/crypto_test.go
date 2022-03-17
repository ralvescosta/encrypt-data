package crypto

import (
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"testing"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
	"github.com/stretchr/testify/assert"
)

func Test_Encrypt(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		mock(nil, nil, nil)

		result, err := Encrypt("pubKey", make(map[string]interface{}))

		assert.NoError(t, err)
		assert.Equal(t, []byte("encrypted"), result)
	})

	t.Run("should return error if some error occur in parsePKCS1PubKey", func(t *testing.T) {
		mock(errors.New("some error"), nil, nil)

		result, err := Encrypt("pubKey", make(map[string]interface{}))

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error if some error occur in jsonMarshal", func(t *testing.T) {
		mock(nil, errors.New("some error"), nil)

		result, err := Encrypt("pubKey", make(map[string]interface{}))

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error if some error occur in jweEncrypt", func(t *testing.T) {
		mock(nil, nil, errors.New("some error"))

		result, err := Encrypt("pubKey", make(map[string]interface{}))

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func mock(parsePKCS1PubKeyError, jsonMarshalError, jweEncryptError error) {
	pemDecode = func(data []byte) (p *pem.Block, rest []byte) {
		return &pem.Block{}, []byte(nil)
	}
	parsePKCS1PubKey = func(der []byte) (*rsa.PublicKey, error) {
		return &rsa.PublicKey{}, parsePKCS1PubKeyError
	}
	jsonMarshal = func(v any) ([]byte, error) {
		return []byte(""), jsonMarshalError
	}
	jweEncrypt = func(
		payload []byte,
		keyalg jwa.KeyEncryptionAlgorithm,
		key interface{},
		contentalg jwa.ContentEncryptionAlgorithm,
		compressalg jwa.CompressionAlgorithm,
		options ...jwe.EncryptOption,
	) ([]byte, error) {
		return []byte("encrypted"), jweEncryptError
	}

}
