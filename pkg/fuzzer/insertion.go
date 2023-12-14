package fuzzer

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) insertion(inch chan string, wg *sync.WaitGroup) {
	log.Debug("Running insertion")
	defer wg.Done()
	isActiveWg := &sync.WaitGroup{}
	for i := 1; i < len(f.Domain)-1; i++ {
		prefix, origC, suffix := f.Domain[:i], rune(f.Domain[i]), f.Domain[i+1:]
		for _, keys := range f.Keyboards {
			for _, c := range keys.Keys[origC] {
				isActiveWg.Add(1)
				go urlparser.IsActiveDomain(prefix+string(c)+string(origC)+suffix, f.TLD, inch, isActiveWg)
				isActiveWg.Add(1)
				go urlparser.IsActiveDomain(prefix+string(origC)+string(c)+suffix, f.TLD, inch, isActiveWg)
			}
		}
	}
	isActiveWg.Wait()
}
