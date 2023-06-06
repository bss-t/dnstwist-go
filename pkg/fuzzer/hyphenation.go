package fuzzer

import (
	"fmt"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) hyphenation(wg *sync.WaitGroup, hych chan string) {
	fmt.Println("Inside hyphenation----------------------------------------------")

	defer wg.Done()
	defer close(hych)

	for i := 1; i < len(f.Domain); i++ {
		validDomain, err := urlparser.IsActiveDomain(f.Domain[:i]+"-"+f.Domain[i:], f.TLD)
		if err == nil && validDomain != "" {
			hych <- validDomain
		}

	}
}
