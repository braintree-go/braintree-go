package braintree

import (
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strings"
)

type SignatureError struct {
	message string
}

func (i SignatureError) Error() string {
	if i.message == "" {
		return "Invalid Signature"
	}
	return i.message
}

func newHmacer(bt *Braintree) hmacer {
	return hmacer{bt}
}

type hmacer struct {
	*Braintree
}

func (h hmacer) verifySignature(signature, payload string) (bool, error) {
	signature, err := h.parseSignature(signature)
	if err != nil {
		return false, err
	}
	expectedSignature, err := h.hmac(payload)
	if err != nil {
		return false, err
	}
	return hmac.Equal([]byte(expectedSignature), []byte(signature)), nil
}

func (h hmacer) getMatchingSignature(signaturePairs string) (string, string) {
	pairs := strings.Split(signaturePairs, "&")
	for _, pair := range pairs {
		split := strings.Split(pair, "|")
		if len(split) == 2 && split[0] == h.Braintree.PublicKey {
			return split[0], split[1]
		}
	}
	return "", ""
}

func (h hmacer) parseSignature(signatureKeyPairs string) (string, error) {
	if !strings.Contains(signatureKeyPairs, "|") {
		return "", SignatureError{"Signature-key pair does not contain |"}
	}
	publicKey, signature := h.getMatchingSignature(signatureKeyPairs)
	if publicKey != h.Braintree.PublicKey {
		return "", SignatureError{"Signature-key pair contains the wrong public key!"}
	}
	return signature, nil
}

func (h hmacer) hmac(payload string) (string, error) {
	s := sha1.New()
	_, err := io.WriteString(s, h.PrivateKey)
	if err != nil {
		return "", errors.New("Could not write private key to SHA1")
	}
	mac := hmac.New(sha1.New, s.Sum(nil))
	_, err = mac.Write([]byte(payload))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", mac.Sum(nil)), nil
}
