package fuzzer

import (
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) replacement(rech chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	isActiveWg := &sync.WaitGroup{}
	replacements := make(map[string]bool)

	for i, c := range f.Domain {
		pre := f.Domain[:i]
		suf := f.Domain[i+1:]

		for _, layout := range f.Keyboards {
			for _, r := range layout.Keys[c] {
				replacement := pre + string(r) + suf
				isActiveWg.Add(1)
				urlparser.IsActiveDomain(replacement, f.TLD, rech, isActiveWg)

				replacements[replacement] = true
			}
		}
	}

	isActiveWg.Wait()
}
