package main

import (
	"flag"
	"fmt"
	"github.com/k0swe/lotw-qsl"
	"os"
)

func main() {
	user := flag.String("username", "", "QRZ.com login name")
	pw := flag.String("password", "", "QRZ.com password")
	flag.Parse()
	if *user == "" || *pw == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	adifStr, err := lotw.Query(*user, *pw, &lotw.QueryOpts{})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(adifStr)
}
