package fuzzer

import (
	"fmt"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) insertion(wg *sync.WaitGroup, inch chan string) {
	fmt.Println("Inside insertion----------------------------------------------")

	defer wg.Done()
	defer close(inch)

	for i := 1; i < len(f.Domain)-1; i++ {
		prefix, origC, suffix := f.Domain[:i], rune(f.Domain[i]), f.Domain[i+1:]
		for _, keys := range f.Keyboards {
			for _, c := range keys.Keys[origC] {
				validDomain1, err := urlparser.IsActiveDomain(prefix+string(c)+string(origC)+suffix, f.TLD)
				if err == nil && validDomain1 != "" {
					inch <- validDomain1
				}

				validDomain2, _ := urlparser.IsActiveDomain(prefix+string(origC)+string(c)+suffix, f.TLD)
				if err == nil && validDomain2 != "" {
					inch <- validDomain2
				}

			}
		}
	}
}
