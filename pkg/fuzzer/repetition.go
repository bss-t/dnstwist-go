package fuzzer

import (
	"log"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) repetition(rech chan string, wg *sync.WaitGroup) {
	log.Println("Inside repetition----------------------------------------------")
	defer wg.Done()
	isActiveWg := &sync.WaitGroup{}
	repetitions := make(map[string]bool)
	for i, c := range f.Domain {
		repetition := f.Domain[:i] + string(c) + f.Domain[i:]
		isActiveWg.Add(1)
		go urlparser.IsActiveDomain(repetition, f.TLD, rech, isActiveWg)
		repetitions[repetition] = true
	}

	isActiveWg.Wait()
}
