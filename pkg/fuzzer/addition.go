package fuzzer

import (
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) addition(rech chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	isActiveWg := &sync.WaitGroup{}
	results := make(map[string]bool)

	for i := 48; i < 58; i++ {
		result1 := f.Domain + string(rune(i))
		isActiveWg.Add(1)
		urlparser.IsActiveDomain(result1, f.TLD, rech, isActiveWg)
		results[result1] = true
		result2 := string(rune(i)) + f.Domain
		isActiveWg.Add(1)
		urlparser.IsActiveDomain(result2, f.TLD, rech, isActiveWg)
		results[result2] = true
	}
	for i := 97; i < 123; i++ {
		result1 := f.Domain + string(rune(i))
		isActiveWg.Add(1)
		urlparser.IsActiveDomain(result1, f.TLD, rech, isActiveWg)
		results[result1] = true
		result2 := string(rune(i)) + f.Domain
		isActiveWg.Add(1)
		urlparser.IsActiveDomain(result2, f.TLD, rech, isActiveWg)
		results[result2] = true
	}
	isActiveWg.Wait()
}
