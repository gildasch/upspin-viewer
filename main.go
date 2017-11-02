package main

import (
	"fmt"

	"upspin.io/client"
	"upspin.io/config"
	_ "upspin.io/transports"
)

func main() {
	// client := client.New(config.New())
	conf, err := config.FromFile("/home/gildas/upspin/config.anonymous")
	if err != nil {
		fmt.Println(err)
	}
	client := client.New(conf)

	_, err = client.Lookup("gildaschbt@gmail.com/p", true)
	// ds, err := client.Lookup("augie@upspin.io", true)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("DirServer: %#v\n", ds)

	b, err := client.Get("gildaschbt@gmail.com/p/index.html")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
