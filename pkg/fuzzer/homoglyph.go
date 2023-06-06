package fuzzer

import (
	"fmt"
	"strings"
	"sync"

	"github.com/balasiddhartha-t/dnstwist-go/pkg/urlparser"
)

func (f *Fuzzer) homoglyph(wg *sync.WaitGroup, hoch chan string) {
	fmt.Println("Inside homoglyph----------------------------------------------")

	defer wg.Done()
	defer close(hoch)

	type void struct{}
	var member void

	glyph := f.Glyphs
	resultVals := make([]string, 0)

	result1 := mix(f.Domain, glyph)

	set := make(map[string]void)
	for _, r := range result1 {
		set[r] = member
		resultVals = append(resultVals, r)
	}

	for r := range resultVals {
		result2 := mix(resultVals[r], glyph)
		for _, r2 := range result2 {
			_, exists := set[r2]
			if !exists {
				set[r2] = member
				resultVals = append(resultVals, r2)

			}
		}
	}

	for k := range resultVals {
		validDomain, err := urlparser.IsActiveDomain(resultVals[k], f.TLD)
		if err == nil && validDomain != "" {
			hoch <- validDomain
		}

	}
}

func mix(domain string, glyph map[byte][]string) []string {

	var result []string
	for w := 1; w < len(domain); w++ {
		for i := 0; i < len(domain)-w+1; i++ {
			pre := domain[:i]
			win := domain[i : i+w]
			suf := domain[i+w:]
			for _, char := range win {
				for _, g := range glyph[byte(char)] {
					result = append(result, pre+strings.ReplaceAll(win, string(char), g)+suf)
				}
			}
		}
	}
	return result
}
