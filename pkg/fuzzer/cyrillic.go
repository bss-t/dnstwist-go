package fuzzer

import (
	"fmt"
	"strings"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) cyrillic(wg *sync.WaitGroup, cych chan string) {
	fmt.Println("Inside cyrillic----------------------------------------------")

	defer wg.Done()
	defer close(cych)

	cdomain := f.Domain
	actualDomain := f.Domain + "." + f.TLD
	// Mapping of latin to cyrillic characters
	latinToCyrillic := map[string]string{
		"a": "а", "b": "ь", "c": "с", "d": "ԁ", "e": "е", "g": "ԍ", "h": "һ",
		"i": "і", "j": "ј", "k": "к", "l": "ӏ", "m": "м", "o": "о", "p": "р",
		"q": "ԛ", "s": "ѕ", "t": "т", "v": "ѵ", "w": "ԝ", "x": "х", "y": "у",
	}

	// Replace latin characters with cyrillic characters in the domain
	for l, c := range latinToCyrillic {
		cdomain = strings.ReplaceAll(cdomain, l, c)
	}

	for i, c := range cdomain {
		if c != rune(actualDomain[i]) {
			validDomain, err := urlparser.IsActiveDomain(cdomain, f.TLD)
			if err == nil && validDomain != "" {
				cych <- validDomain
			}

			break
		}
	}
}
