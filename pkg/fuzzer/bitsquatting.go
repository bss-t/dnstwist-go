package fuzzer

import (
	"fmt"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) bitsquatting(wg *sync.WaitGroup, bsch chan string) {
	fmt.Println("Inside BitSquatting----------------------------------------------")
	defer wg.Done()
	defer close(bsch)

	bits := []byte(f.Domain)

	for i := 0; i < len(bits)*8; i++ {
		// Flip a single bit
		bits[i/8] ^= 1 << uint(7-i%8)

		// Construct the new domain name
		newDomain := string(bits)
		validDomain, err := urlparser.IsActiveDomain(newDomain, f.TLD)
		if err == nil && validDomain != "" {
			bsch <- validDomain
		}

		// Flip the bit back to its original value
		bits[i/8] ^= 1 << uint(7-i%8)
	}
}
