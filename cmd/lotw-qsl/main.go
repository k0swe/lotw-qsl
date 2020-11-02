package main

import (
	"flag"
	"fmt"
	"github.com/antihax/optional"
	"github.com/k0swe/lotw-qsl"
	"os"
)

func main() {
	user := flag.String("username", "", "QRZ.com login name")
	pw := flag.String("password", "", "QRZ.com password")
	since := flag.String("since", "", "Beginning date to query for; YYYYMMDD format")
	flag.Parse()
	if *user == "" || *pw == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	opts := &lotw.QueryOpts{}
	if *since != "" {
		opts.QsoStartdate = optional.NewString(*since)
	}
	adifStr, err := lotw.Query(*user, *pw, opts)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(adifStr)
}
