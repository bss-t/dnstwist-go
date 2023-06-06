package urlparser

import (
	"log"
	"net"
	"net/url"
	"regexp"
	"sync"

	"golang.org/x/net/idna"
)

type ParsedUrl struct {
	Scheme   string
	Host     string
	Path     string
	Query    string
	Fragment string
}

var fqdnRegex = regexp.MustCompile(`^(https?://)?[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?)*\.[a-zA-Z]{2,}$`)

func (p ParsedUrl) Parse(urlString string) ParsedUrl {
	valid_domain := ValidateDomain(urlString)
	if !valid_domain {
		panic("Invalid Domain")
	}
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	p.Scheme = parsedUrl.Scheme
	p.Host = parsedUrl.Host
	p.Path = parsedUrl.Path
	p.Query = parsedUrl.RawQuery
	p.Fragment = parsedUrl.Fragment
	return p
}

func ValidateDomain(domain string) bool {
	// If the length is  greater than 253 return an error
	if len(domain) < 4 || len(domain) > 253 {
		return false
	}
	// Regex for  Fully Qualified Domain Names (FQDNs)
	matched := fqdnRegex.MatchString(domain)
	if matched {
		// decode the domain
		_, err := idna.ToUnicode(domain)
		if err != nil {
			log.Println("Error encoding domain:", err)
			return false
		}
		return true
	} else {
		log.Println("Match not found.")
		return false
	}
}

func IsActiveDomain(domain string, tld string, activeDomainChannel chan string, wg *sync.WaitGroup) (string, error) {
	defer wg.Done()
	if domain != "" {
		domain = domain + "." + tld
	}
	// Check if the domain name exists
	_, err := net.LookupHost(domain)
	if err == nil {
		activeDomainChannel <- domain
	}
	return "", nil
}
