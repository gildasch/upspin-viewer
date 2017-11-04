package anonymous

import (
	"upspin.io/factotum"
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

var (
	// from upspin.io/config/initconfig.go
	defaultPacking     = upspin.EEPack
	defaultKeyEndpoint = upspin.Endpoint{
		Transport: upspin.Remote,
		NetAddr:   "key.upspin.io:443",
	}
)

type config struct {
	factotum upspin.Factotum
}

func (config) UserName() upspin.UserName      { return username }
func (c *config) Factotum() upspin.Factotum   { return c.factotum }
func (config) Packing() upspin.Packing        { return defaultPacking }
func (config) KeyEndpoint() upspin.Endpoint   { return defaultKeyEndpoint }
func (config) DirEndpoint() upspin.Endpoint   { return upspin.Endpoint{} }
func (config) StoreEndpoint() upspin.Endpoint { return upspin.Endpoint{} }
func (config) Value(string) string            { return "" }

func newConfig() (*config, error) {
	factotum, err := factotum.NewFromKeys([]byte(public), []byte(secret), nil)
	if err != nil {
		return nil, err
	}

	return &config{factotum: factotum}, nil
}
