package fuzzer

import (
	"log"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

var hyphenationwg sync.WaitGroup

func (f *Fuzzer) hyphenation(hych chan string, Domain string, hyphenationwg *sync.WaitGroup) {
	log.Println("Inside hyphenation----------------------------------------------")
	defer hyphenationwg.Done()
	isActiveWg := &sync.WaitGroup{}
	for i := 1; i < len(Domain); i++ {
		newDomain := Domain[:i] + "-" + Domain[i:]
		isActiveWg.Add(1)
		go urlparser.IsActiveDomain(newDomain, "http", hych, isActiveWg)
	}
	isActiveWg.Wait()
}
