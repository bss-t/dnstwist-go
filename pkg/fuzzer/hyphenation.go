package fuzzer

import (
	"sync"

	"github.com/bss-t/dnstwist-go/pkg/urlparser"
	log "github.com/sirupsen/logrus"
)

var hyphenationwg sync.WaitGroup

func (f *Fuzzer) hyphenation(hych chan string, Domain string, hyphenationwg *sync.WaitGroup) {
	log.Debug("Running hyphenation")
	defer hyphenationwg.Done()
	isActiveWg := &sync.WaitGroup{}
	for i := 1; i < len(Domain); i++ {
		newDomain := Domain[:i] + "-" + Domain[i:]
		isActiveWg.Add(1)
		go urlparser.IsActiveDomain(newDomain, "http", hych, isActiveWg)
	}
	isActiveWg.Wait()
}
