package fuzzer

import (
	"os"
	"sync"

	tld "github.com/jpillora/go-tld"
	log "github.com/sirupsen/logrus"
)

type Fuzzer struct {
	Domain    string
	Subdomain string
	TLD       string
	Glyphs    map[byte][]string
	Keyboards []Keyboards
}

type Keyboards struct {
	Layout string
	Keys   map[rune]string
}

func (f *Fuzzer) Fuzz(host string, dictionary []string, tld_dictionary []string) {

	validDomains := make([]string, 0)
	done := make(chan struct{})
	var fuzzWg sync.WaitGroup
	qwerty := Keyboards{
		Layout: "qwerty",
		Keys: map[rune]string{
			'1': "2q", '2': "3wq1", '3': "4ew2", '4': "5re3", '5': "6tr4", '6': "7yt5", '7': "8uy6", '8': "9iu7", '9': "0oi8", '0': "po9", 'q': "12wa", 'w': "3esaq2", 'e': "4rdsw3", 'r': "5tfde4", 't': "6ygfr5", 'y': "7uhgt6", 'u': "8ijhy7", 'i': "9okju8", 'o': "0plki9", 'p': "lo0", 'a': "qwsz", 's': "edxzaw", 'd': "rfcxse", 'f': "tgvcdr", 'g': "yhbvft", 'h': "ujnbgy", 'j': "ikmnhu", 'k': "olmji", 'l': "kop", 'z': "asx", 'x': "zsdc", 'c': "xdfv", 'v': "cfgb", 'b': "vghn", 'n': "bhjm", 'm': "njk",
		},
	}
	qwertz := Keyboards{
		Layout: "qwertz",
		Keys: map[rune]string{
			'1': "2q", '2': "3wq1", '3': "4ew2", '4': "5re3", '5': "6tr4", '6': "7zt5", '7': "8uz6", '8': "9iu7", '9': "0oi8", '0': "po9", 'q': "12wa", 'w': "3esaq2", 'e': "4rdsw3", 'r': "5tfde4", 't': "6zgfr5", 'z': "7uhgt6", 'u': "8ijhz7", 'i': "9okju8", 'o': "0plki9", 'p': "lo0", 'a': "qwsy", 's': "edxyaw", 'd': "rfcxse", 'f': "tgvcdr", 'g': "zhbvft", 'h': "ujnbgz", 'j': "ikmnhu", 'k': "olmji", 'l': "kop", 'y': "asx", 'x': "ysdc", 'c': "xdfv", 'v': "cfgb", 'b': "vghn", 'n': "bhjm", 'm': "njk",
		},
	}

	azerty := Keyboards{
		Layout: "azerty",
		Keys: map[rune]string{
			'1': "2a", '2': "3za1", '3': "4ez2", '4': "5re3", '5': "6tr4", '6': "7yt5", '7': "8uy6", '8': "9iu7", '9': "0oi8", '0': "po9", 'a': "2zq1", 'z': "3esqa2", 'e': "4rdsz3", 'r': "5tfde4", 't': "6ygfr5", 'y': "7uhgt6", 'u': "8ijhy7", 'i': "9okju8", 'o': "0plki9", 'p': "lo0m", 'q': "zswa", 's': "edxwqz", 'd': "rfcxse", 'f': "tgvcdr", 'g': "yhbvft", 'h': "ujnbgy", 'j': "iknhu", 'k': "olji", 'l': "kopm", 'm': "lp", 'w': "sxq", 'x': "wsdc", 'c': "xdfv", 'v': "cfgb", 'b': "vghn", 'n': "bhj",
		},
	}

	f.Keyboards = []Keyboards{qwerty, qwertz, azerty}

	f.Glyphs = map[byte][]string{
		'0': {"o"},
		'1': {"l", "i"},
		'2': {"ƻ"},
		'3': {"8"},
		'5': {"ƽ"},
		'6': {"9"},
		'8': {"3"},
		'9': {"6"},
		'a': {"à", "á", "à", "â", "ã", "ä", "å", "ɑ", "ạ", "ǎ", "ă", "ȧ", "ą", "ə"},
		'b': {"d", "lb", "ʙ", "ɓ", "ḃ", "ḅ", "ḇ", "ƅ"},
		'c': {"e", "ƈ", "ċ", "ć", "ç", "č", "ĉ", "ᴄ"},
		'd': {"b", "cl", "dl", "ɗ", "đ", "ď", "ɖ", "ḑ", "ḋ", "ḍ", "ḏ", "ḓ"},
		'e': {"c", "é", "è", "ê", "ë", "ē", "ĕ", "ě", "ė", "ẹ", "ę", "ȩ", "ɇ", "ḛ"},
		'f': {"ƒ", "ḟ"},
		'g': {"q", "ɢ", "ɡ", "ġ", "ğ", "ǵ", "ģ", "ĝ", "ǧ", "ǥ"},
		'h': {"lh", "ĥ", "ȟ", "ħ", "ɦ", "ḧ", "ḩ", "ⱨ", "ḣ", "ḥ", "ḫ", "ẖ"},
		'i': {"1", "l", "í", "ì", "ï", "ı", "ɩ", "ǐ", "ĭ", "ỉ", "ị", "ɨ", "ȋ", "ī", "ɪ"},
		'j': {"ʝ", "ǰ", "ɉ", "ĵ"},
		'k': {"lk", "ik", "lc", "ḳ", "ḵ", "ⱪ", "ķ", "ᴋ"},
		'l': {"1", "i", "ɫ", "ł"},
		'm': {"n", "nn", "rn", "rr", "ṁ", "ṃ", "ᴍ", "ɱ", "ḿ"},
		'n': {"m", "r", "ń", "ṅ", "ṇ", "ṉ", "ñ", "ņ", "ǹ", "ň", "ꞑ"},
		'o': {"0", "ȯ", "ọ", "ỏ", "ơ", "ó", "ö", "ᴏ"},
		'p': {"ƿ", "ƥ", "ṕ", "ṗ"},
		'q': {"g", "ʠ"},
		'r': {"ʀ", "ɼ", "ɽ", "ŕ", "ŗ", "ř", "ɍ", "ɾ", "ȓ", "ȑ", "ṙ", "ṛ", "ṟ"},
		's': {"ʂ", "ś", "ṣ", "ṡ", "ș", "ŝ", "š", "ꜱ"},
		't': {"ţ", "ŧ", "ṫ", "ṭ", "ț", "ƫ"},
		'u': {"ᴜ", "ǔ", "ŭ", "ü", "ʉ", "ù", "ú", "û", "ũ", "ū", "ų", "ư", "ů", "ű", "ȕ", "ȗ", "ụ"},
		'v': {"ṿ", "ⱱ", "ᶌ", "ṽ", "ⱴ", "ᴠ"},
		'w': {"vv", "ŵ", "ẁ", "ẃ", "ẅ", "ⱳ", "ẇ", "ẉ", "ẘ", "ᴡ"},
		'x': {"ẋ", "ẍ"},
		'y': {"ʏ", "ý", "ÿ", "ŷ", "ƴ", "ȳ", "ɏ", "ỿ", "ẏ", "ỵ"},
		'z': {"ʐ", "ż", "ź", "ᴢ", "ƶ", "ẓ", "ẕ", "ⱬ"},
	}

	u, _ := tld.Parse(host)
	log.Info(u, " ", u.Subdomain, " ", u.Domain, " ", u.TLD, " ", u.Port, " ", u.Path, " ", u.ICANN)

	f.Domain, f.Subdomain, f.TLD = u.Domain, u.Subdomain, u.TLD

	activeDomainChannel := make(chan string, 10000)
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	file1, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fuzzWg.Add(1)
	go f.bitsquatting(activeDomainChannel, &fuzzWg)

	fuzzWg.Add(1)
	go f.cyrillic(activeDomainChannel, &fuzzWg)

	fuzzWg.Add(1)
	go f.homoglyph(activeDomainChannel, &fuzzWg)

	fuzzWg.Add(1)
	go f.hyphenation(activeDomainChannel, f.Domain, &fuzzWg)

	fuzzWg.Add(1)
	go f.insertion(activeDomainChannel, &fuzzWg)

	fuzzWg.Add(1)
	go f.omission(activeDomainChannel, &fuzzWg)

	fuzzWg.Add(1)
	go f.repetition(activeDomainChannel, &fuzzWg)

	fuzzWg.Add(1)
	go f.replacement(activeDomainChannel, &fuzzWg)

	fuzzWg.Add(1)
	go f.subdomain(activeDomainChannel, &fuzzWg)

	fuzzWg.Add(1)
	go f.addition(activeDomainChannel, &fuzzWg)

	go func() {
		defer func() {
			close(done) // Close the done channel to signal completion
		}()
		for {
			select {
			case validDomain, ok := <-activeDomainChannel:
				if !ok {
					return
				}
				if _, err = file1.WriteString(validDomain + "\n"); err != nil {
					panic(err)
				}
				validDomains = append(validDomains, validDomain)
			case <-done:
				return
			}
		}
	}()

	fuzzWg.Wait()
	close(activeDomainChannel)
	<-done

	log.Debug(validDomains)
}
