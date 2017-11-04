package anonymous

import (
	"upspin.io/client"
	"upspin.io/upspin"
)

func NewClient() (upspin.Client, error) {
	conf, err := newConfig()
	if err != nil {
		return nil, err
	}
	return client.New(conf), nil
}
