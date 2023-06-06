package main

import (
	"io"
	"log"
	"os"
	"time"

	fuzz "github.com/balasiddhartha-t/dnstwist-go/pkg/fuzzer"
	urlparser "github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func init() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(mw)
}
func main() {

	// Hardcoded values to be read from the args
	start := time.Now()
	domain := "https://google.com"
	dictionary := make([]string, 0)
	tld := make([]string, 0)

	var url urlparser.ParsedUrl
	log.Println("Welcome to DNS Twist using go")
	url = url.Parse(domain)
	f := fuzz.Fuzzer{}
	fuzzer := f.Fuzz(url.Scheme+"://"+url.Host, dictionary, tld)
	log.Println(fuzzer)
	elapsed := time.Since(start)
	log.Printf("Time taken by fuzzer is ----------------> %s", elapsed)
}
