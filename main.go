package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"upspin.io/client"
	"upspin.io/factotum"
	_ "upspin.io/transports"
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

type Config struct {
	factotum upspin.Factotum
}

func (Config) UserName() upspin.UserName      { return username }
func (c *Config) Factotum() upspin.Factotum   { return c.factotum }
func (Config) Packing() upspin.Packing        { return defaultPacking }
func (Config) KeyEndpoint() upspin.Endpoint   { return defaultKeyEndpoint }
func (Config) DirEndpoint() upspin.Endpoint   { return upspin.Endpoint{} }
func (Config) StoreEndpoint() upspin.Endpoint { return upspin.Endpoint{} }
func (Config) Value(string) string            { return "" }

func newConfig() (*Config, error) {
	factotum, err := factotum.NewFromKeys([]byte(public), []byte(secret), nil)
	if err != nil {
		return nil, err
	}

	return &Config{factotum: factotum}, nil
}

func main() {
	conf, err := newConfig()
	if err != nil {
		fmt.Println(err)
	}
	client := client.New(conf)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.TrimPrefix(r.URL.Path, "/")

		de, err := client.Lookup(upspin.PathName(url), true)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if de.IsDir() {
			des, err := client.Glob(upspin.AllFilesGlob(upspin.PathName(url)))
			if err != nil {
				fmt.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for _, dess := range des {
				fmt.Fprintf(w, "<a href='/%s'>%s</a><br />", dess.SignedName, dess.SignedName)
			}
			return
		}

		f, err := client.Open(upspin.PathName(url))
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		_, err = io.Copy(w, f)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
