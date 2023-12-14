[dnstwist]()
===============================

DNS fuzzing is an automated workflow for discovering potentially malicious domains targeting your organisation. This tool works by generating a large list of permutations based on a domain name you provide and then checking if any of those permutations are in use.
Additionally, it can generate fuzzy hashes of the web pages to see if they are part of an ongoing phishing attack or brand impersonation, and much more!

This is a go version of https://github.com/elceef/dnstwist

Installation
-------------


**Git**
You can checkout the version of code.

```
$ git clone https://github.com/balasiddhartha-t/dnstwist-go.git
$ cd dnstwist-go
$ go build -o dnstwist-go .\cmd\main.go
```
***Note:*** For windows users please use ```go build -o dnstwist-go.exe .\cmd\main.go```
