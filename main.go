package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

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
		start := time.Now()
		url := strings.TrimPrefix(r.URL.Path, "/")
		fmt.Printf("url extract: %s\n", time.Since(start))
		f, err := client.Open(upspin.PathName(url))
		fmt.Printf("client.Open: %s\n", time.Since(start))
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		_, err = io.Copy(w, f)
		fmt.Printf("io.Copy: %s\n", time.Since(start))
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("return: %s\n", time.Since(start))
	})

	fmt.Println("Listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
