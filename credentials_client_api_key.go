package braintree

import "encoding/base64"

type clientApiKey struct {
	env          Environment
	clientId     string
	clientSecret string
}

func newClientAPIKey(env Environment, clientId, clientSecret string) credentials {
	return clientApiKey{
		env:          env,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (k clientApiKey) Environment() Environment {
	return k.env
}

func (k clientApiKey) MerchantID() string {
	return ""
}

func (k clientApiKey) AuthorizationHeader() string {
	auth := k.clientId + ":" + k.clientSecret
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
