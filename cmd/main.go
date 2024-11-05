package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	fuzz "github.com/bss-t/dnstwist-go/pkg/fuzzer"
	urlparser "github.com/bss-t/dnstwist-go/pkg/urlparser"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var debug bool

func main() {
	var url urlparser.ParsedUrl

	// TODO: Hardcoded values to be read from the args
	dir := "logs"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Info(dir, " does not exist, creating")
		os.Mkdir("logs", 0755)
	} else {
		fmt.Println("The provided directory named", dir, "exists")
	}
	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	log.SetFormatter(&log.JSONFormatter{})
	debug := isDebugMode()

	// Read the value of the Domain
	pflag.String("domain", "", "Domain for which the fuzzers need to be generated")
	pflag.BoolVarP(&debug, "debug", "d", false, "Start in Debug mode")
	pflag.Parse()
	log.Info("debug: ", debug)

	if debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Bind the flag to viper
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()

	if viper.IsSet("domain") {
		start := time.Now()
		domain := viper.GetString("domain")
		dictionary := make([]string, 0)
		tld := make([]string, 0)
		f := fuzz.Fuzzer{}
		if !strings.Contains(domain, "http") {
			log.Error("Please pass the domain with http:// or https://")
			return
		}
		url = url.Parse(domain)
		f.Fuzz(url.Scheme+"://"+url.Host, dictionary, tld)
		elapsed := time.Since(start)
		log.Info("Time taken by fuzzer is ", elapsed)
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
