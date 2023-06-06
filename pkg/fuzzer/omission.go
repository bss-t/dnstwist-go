package fuzzer

import (
	"log"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) omission(omch chan string, wg *sync.WaitGroup) {
	log.Println("Inside omission----------------------------------------------")
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
