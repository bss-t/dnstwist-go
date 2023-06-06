package fuzzer

import (
	"fmt"
	"log"
	"os"
	"sync"

	tld "github.com/jpillora/go-tld"
)

var wg sync.WaitGroup

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

func (f *Fuzzer) Fuzz(host string, dictionary []string, tld_dictionary []string) []string {

	validDomains := make([]string, 0)

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
	fmt.Printf("%50s = [ %s ] [ %s ] [ %s ] [ %s ] [ %s ] [ %t ]\n",
		u, u.Subdomain, u.Domain, u.TLD, u.Port, u.Path, u.ICANN)

	f.Domain, f.Subdomain, f.TLD = u.Domain, u.Subdomain, u.TLD

	wg.Add(1)
	hoch := make(chan string, 10000)
	go f.homoglyph(&wg, hoch)

	wg.Add(1)
	bsch := make(chan string, 10000)
	go f.bitsquatting(&wg, bsch)

	wg.Add(1)
	cych := make(chan string, 10000)
	go f.cyrillic(&wg, cych)

	wg.Add(1)
	hych := make(chan string, 10000)
	go f.hyphenation(&wg, hych)

	wg.Add(1)
	inch := make(chan string, 10000)
	go f.insertion(&wg, inch)

	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	file1, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		select {
		case ho := <-hoch:
			if ho != "" {
				fmt.Println("Consuming hoch----------------------------------------------")
				validDomains = append(validDomains, ho)

				if _, err = file1.WriteString(ho + "\n"); err != nil {
					panic(err)
				}
			}

		case bs := <-bsch:
			if bs != "" {
				fmt.Println("Consuming bsch----------------------------------------------")
				validDomains = append(validDomains, bs)

				if _, err = file1.WriteString(bs + "\n"); err != nil {
					panic(err)
				}
			}

		case cy := <-cych:
			if cy != "" {
				fmt.Println("Consuming cych----------------------------------------------")
				validDomains = append(validDomains, cy)

				if _, err = file1.WriteString(cy + "\n"); err != nil {
					panic(err)
				}
			}

		case hy := <-hych:
			if hy != "" {
				fmt.Println("Consuming hych----------------------------------------------")
				validDomains = append(validDomains, hy)

				if _, err = file1.WriteString(hy + "\n"); err != nil {
					panic(err)
				}
			}
		case in := <-inch:
			if in != "" {
				fmt.Println("Consuming inch----------------------------------------------")
				validDomains = append(validDomains, in)

				if _, err = file1.WriteString(in + "\n"); err != nil {
					panic(err)
				}
			}

		}
	}
	// for domain := range hoch {
	// 	fmt.Println("Consuming hoch----------------------------------------------")
	// 	validDomains = append(validDomains, domain)
	// }
	// for domain := range bsch {
	// 	fmt.Println("Consuming bsch----------------------------------------------")

	// 	validDomains = append(validDomains, domain)
	// }

	// for domain := range cych {
	// 	fmt.Println("Consuming cych----------------------------------------------")

	// 	validDomains = append(validDomains, domain)
	// }

	// for domain := range hych {
	// 	fmt.Println("Consuming hych----------------------------------------------")

	// 	validDomains = append(validDomains, domain)
	// }

	// for domain := range inch {
	// 	fmt.Println("Consuming inch----------------------------------------------")
	// 	validDomains = append(validDomains, domain)
	// }

	wg.Wait()
	return validDomains
}
