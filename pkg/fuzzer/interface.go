package fuzzer

import "sync"

type IFuzzer interface {
	Fuzz(host string, dictionary []string, tld_dictionary []string)
	homoglyph(hoch chan string, wg *sync.WaitGroup)
	bitsquatting(bsch chan string, wg *sync.WaitGroup)
	cyrillic(cych chan string, wg *sync.WaitGroup)
	hyphenation(hych chan string, Domain string, hyphenationwg *sync.WaitGroup)
	insertion(inch chan string, wg *sync.WaitGroup)
	omission(omch chan string, wg *sync.WaitGroup)
	repetition(rech chan string, wg *sync.WaitGroup)
	replacement(rech chan string, wg *sync.WaitGroup)
	subdomain(rech chan string, wg *sync.WaitGroup)
}
