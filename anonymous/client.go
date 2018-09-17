package anonymous

import (
	"upspin.io/client"
	"upspin.io/upspin"
)

const (
	username = "anonymous@yopmail.com"
	public   = `p256
83864008841660804129850532384760972787428847923376860875699308590865504585405
31845544589173428159178438667181835186482258772868935578446998229155015180983
`
	secret = `55195194792553803807358164051594792146922824178594810928464674606099011268533`
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
