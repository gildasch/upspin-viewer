package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"upspin.io/client"
	"upspin.io/config"
	_ "upspin.io/transports"
	"upspin.io/upspin"
)

func main() {
	// client := client.New(config.New())
	conf, err := config.FromFile("/home/gildas/upspin/config.anonymous")
	if err != nil {
		fmt.Println(err)
	}
	client := client.New(conf)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.TrimPrefix(r.URL.Path, "/")
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
