package fuzzer

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/bss-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) omission(omch chan string, wg *sync.WaitGroup) {
	log.Debug("Running omission")
	defer wg.Done()
	isActiveWg := &sync.WaitGroup{}

	omissions := make(map[string]bool)
	for i := 0; i < len(f.Domain); i++ {
		omission := f.Domain[:i] + f.Domain[i+1:]
		isActiveWg.Add(1)
		go urlparser.IsActiveDomain(omission, f.TLD, omch, isActiveWg)
		omissions[omission] = true
	}

	isActiveWg.Wait()
}
