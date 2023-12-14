package main

import (
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	fuzz "github.com/balasiddhartha-t/dnstwist-go/pkg/fuzzer"
	urlparser "github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	var url urlparser.ParsedUrl

	// TODO: Hardcoded values to be read from the args
	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	log.SetFormatter(&log.JSONFormatter{})
	debugMode := isDebugMode()
	if debugMode {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Read the value of the Domain
	pflag.String("domain", "", "Domain for which the fuzzers need to be generated")
	pflag.Parse()

	// Bind the flag to viper
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()

	if viper.IsSet("domain") {
		start := time.Now()
		domain := viper.GetString("domain")
		dictionary := make([]string, 0)
		tld := make([]string, 0)
		f := fuzz.Fuzzer{}
		url = url.Parse(domain)
		f.Fuzz(url.Scheme+"://"+url.Host, dictionary, tld)
		elapsed := time.Since(start)
		log.Info("Time taken by fuzzer is ---------------- ", elapsed)
	} else {
		log.Error("Please pass domain as a flag")
	}

}

func isDebugMode() bool {
	// Check if the DEBUG environment variable is set
	debugEnv := os.Getenv("DEBUG")
	if debugEnv != "" && debugEnv != "0" {
		return true
	}

	return false
}
