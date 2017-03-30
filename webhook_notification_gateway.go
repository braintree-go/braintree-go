package braintree

import (
	"encoding/base64"
	"encoding/xml"
)

type WebhookNotificationGateway struct {
	*Braintree
}

func (w *WebhookNotificationGateway) Parse(signature, payload string) (*WebhookNotification, error) {
	hmacer := newHmacer(w.Braintree.PublicKey, w.Braintree.PrivateKey)
	if verified, err := hmacer.verifySignature(signature, payload); err != nil {
		return nil, err
	} else if !verified {
		return nil, SignatureError{}
	}

	xmlNotification, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	var n WebhookNotification
	err = xml.Unmarshal(xmlNotification, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (w *WebhookNotificationGateway) Verify(challenge string) (string, error) {
	hmacer := newHmacer(w.Braintree.PublicKey, w.Braintree.PrivateKey)
	digest, err := hmacer.hmac(challenge)
	if err != nil {
		return ``, err
	}
	return hmacer.publicKey + `|` + digest, nil
}
