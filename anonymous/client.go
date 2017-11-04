package anonymous

import (
	"upspin.io/client"
	"upspin.io/upspin"
)

const (
	username = "anonymous@yopmail.com"
	public   = `p256
22093798818893024622336269998686442571435612562520865182046305461083477133512
17797210462092364380982608279430136085800970116434296109768820456901066959334
`
	secret = `59547353025968832991402423037620858463955029496026875163089246321869979101175`
)

// NewClient returns a read-only upspin client
//
// The client is connected with a read-only account:
// anonymous@yopmail.com
func NewClient() (upspin.Client, error) {
	return NewCustomClient(username, []byte(public), []byte(secret))
}

// NewClient returns a read-only upspin client from a custom username
// and a public & secret upspin key pair.
//
// Note: the store and directory endpoints are not set in the returned
// client so it will not be able to write anything. For more general
// upspin client, use the official constructors from upspin.io/client/
func NewCustomClient(username string, public, secret []byte) (upspin.Client, error) {
	conf, err := newConfig(username, []byte(public), []byte(secret))
	if err != nil {
		return nil, err
	}
	return client.New(conf), nil
}
