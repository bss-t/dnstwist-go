package fuzzer

import (
	"log"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) bitsquatting(bsch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("Inside BitSquatting----------------------------------------------")
	bits := []byte(f.Domain)
	isActiveWg := &sync.WaitGroup{}
	for i := 0; i < len(bits)*8; i++ {
		// Flip a single bit
		bits[i/8] ^= 1 << uint(7-i%8)

		// Construct the new domain name
		newDomain := string(bits)
		isActiveWg.Add(1)
		go urlparser.IsActiveDomain(newDomain, f.TLD, bsch, isActiveWg)

		// Flip the bit back to its original value
		bits[i/8] ^= 1 << uint(7-i%8)
	}
	isActiveWg.Wait()
}
