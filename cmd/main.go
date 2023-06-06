package main

import (
	"fmt"
	"time"

	fuzz "github.com/balasiddhartha-t/dnstwist-go/pkg/fuzzer"
	urlparser "github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func main() {
	// Hardcoded values to be read from the args
	start := time.Now()
	domain := "https://google.com"
	dictionary := make([]string, 0)
	tld := make([]string, 0)

	var url urlparser.ParsedUrl
	fmt.Println("Welcome to DNS Twist using go")
	url = url.Parse(domain)
	f := fuzz.Fuzzer{}
	fuzzer := f.Fuzz(url.Scheme+"://"+url.Host, dictionary, tld)
	fmt.Println(fuzzer)
	elapsed := time.Since(start)
	fmt.Printf("Time taken by fuzzer is ----------------> %s", elapsed)
}
