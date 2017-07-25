package main

import (
	"fmt"
	"os"
	"regexp"

	"golang.org/x/net/idna"
)

var usage = `
██████╗ ██╗   ██╗███╗   ██╗██╗   ██╗ ██████╗ ██████╗ ██████╗ ███████╗██████╗ 
██╔══██╗██║   ██║████╗  ██║╚██╗ ██╔╝██╔════╝██╔═══██╗██╔══██╗██╔════╝██╔══██╗
██████╔╝██║   ██║██╔██╗ ██║ ╚████╔╝ ██║     ██║   ██║██║  ██║█████╗  ██████╔╝
██╔═══╝ ██║   ██║██║╚██╗██║  ╚██╔╝  ██║     ██║   ██║██║  ██║██╔══╝  ██╔══██╗
██║     ╚██████╔╝██║ ╚████║   ██║   ╚██████╗╚██████╔╝██████╔╝███████╗██║  ██║
╚═╝      ╚═════╝ ╚═╝  ╚═══╝   ╚═╝    ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝

Converts punycode to unicode and vice-versa.

Usage:
	punycoder <domain>
`

func main() {

	if len(os.Args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	domain := os.Args[1]
	var re = regexp.MustCompile(`^xn--`)

	switch {
	case domain == "-h", domain == "--help":
		fmt.Println(usage)
	case re.Match([]byte(domain)):
		//punycode was provided, convert to Unicode
		unidomain, _ := idna.ToUnicode(domain)
		fmt.Printf("%s\n", unidomain)
	default:
		//convert to ASCII punycode
		asciidomain, _ := idna.ToASCII(domain)
		fmt.Printf("%s\n", asciidomain)
	}

}
