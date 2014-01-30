package braintree

import (
	"encoding/base64"
	"regexp"
	"strings"
)

type InvalidSignatureError struct {
}

func (i *InvalidSignatureError) Error() string {
	return "Invalid Signature"
}

type WebhookNotificationGateway struct {
	*Braintree
}

var tagRewritter = regexp.MustCompile(`(<\/?[^_\s]+?)_([^_]+?>)`)

func (w *WebhookNotificationGateway) Parse(signature, payload string) (*WebhookNotification, error) {
	if err := w.verifySignature(signature, payload); err != nil {
		return nil, err
	}

	xmlStr, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	xmlStr = tagRewritter.ReplaceAllFunc(xmlStr, func(str []byte) []byte {
		return []byte(strings.Replace(string(str), "_", "-", -1))
	})

	return NewWebhookNotification([]byte(xmlStr))
}

func (w *WebhookNotificationGateway) verifySignature(signature, payload string) error {
	publicKey, signature := w.matchingSignaturePair(signature)
	payloadSignature := Hexdigest(w.PrivateKey, payload)

	if len(publicKey) == 0 || !SecureCompare(signature, payloadSignature) {
		return &InvalidSignatureError{}
	}

	return nil
}

func (w *WebhookNotificationGateway) matchingSignaturePair(signature string) (string, string) {
	pairs := strings.Split(signature, "&")
	const separator = "|"
	for _, pair := range pairs {
		if strings.Contains(pair, separator) {
			comps := strings.Split(pair, separator)
			if comps[0] == w.PublicKey {
				return comps[0], comps[1]
			}
		}
	}
	return "", ""
}
