package fuzzer

import (
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) subdomain(rech chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	isActiveWg := &sync.WaitGroup{}
	subdomains := make(map[string]bool)

	for i := range f.Domain {
		if (i != 0) && (f.Domain[i] != '-' || f.Domain[i] != '.') && (f.Domain[i-1] != '-' || f.Domain[i-1] != '.') {
			subdomain := f.Domain[:i] + string('.') + f.Domain[i:]
			isActiveWg.Add(1)
			urlparser.IsActiveDomain(subdomain, f.TLD, rech, isActiveWg)
			subdomains[subdomain] = true
		}
	}
	isActiveWg.Wait()

}
