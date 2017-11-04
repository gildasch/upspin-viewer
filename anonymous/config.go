package anonymous

import (
	"upspin.io/factotum"
	"upspin.io/upspin"
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
	username upspin.UserName
	factotum upspin.Factotum
}

func (c *config) UserName() upspin.UserName   { return c.username }
func (c *config) Factotum() upspin.Factotum   { return c.factotum }
func (config) Packing() upspin.Packing        { return defaultPacking }
func (config) KeyEndpoint() upspin.Endpoint   { return defaultKeyEndpoint }
func (config) DirEndpoint() upspin.Endpoint   { return upspin.Endpoint{} }
func (config) StoreEndpoint() upspin.Endpoint { return upspin.Endpoint{} }
func (config) Value(string) string            { return "" }

func newConfig(username string, public, secret []byte) (*config, error) {
	factotum, err := factotum.NewFromKeys(public, secret, nil)
	if err != nil {
		return nil, err
	}

	return &config{
		username: upspin.UserName(username),
		factotum: factotum,
	}, nil
}
