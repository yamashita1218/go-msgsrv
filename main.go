package main

import (
	"flag"
	"fmt"
	"net/http"
)

const defaultPort = "80"

type arguments struct {
	port string
}

func main() {
	args := fetchArgs()
	if err := http.ListenAndServe(":"+args.port, nil); err != nil {
		fmt.Println(err)
	}
}

func fetchArgs() *arguments {
	args := new(arguments)
	flag.StringVar(&args.port, "p", defaultPort, "port for http listen.")
	flag.StringVar(&args.port, "port", defaultPort, "port for http listen.")
	flag.Parse()
	return args
}
