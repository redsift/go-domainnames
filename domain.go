package domainnames

import (
	"strings"

	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
)

// VerifyDomainFormatAndRoot checks for a valid look domain and returns
// the publicly registrable root for the domain. Returns empty string if
// domain is already publicly registrable
func VerifyDomainFormatAndRoot(domain string) (string, error) {
	domain = RemoveFQDN(domain)

	parts := strings.Split(domain, ".")
	if l := len(parts); l < 2 {
		return "", ErrMalformedDomain
	}
	ps, icann := publicsuffix.PublicSuffix(domain)
	if !icann {
		return "", ErrNotAnIcannDomain
	}

	i := strings.LastIndex(domain, ps)
	if i < 1 {
		return "", ErrMalformedPublicSuffix
	}

	host := domain[0 : i-1]
	i = strings.LastIndex(host, ".")
	if i != -1 {
		host = host[i+1:]
	}
	if host == "" {
		return "", ErrMalformedDomain
	}

	if composed := host + "." + ps; composed != domain {
		return host + "." + ps, nil
	}

	return "", nil
}

// NormalizeFQDNAndPuny makes a domain fully qualified, trims and lowercases it
// and converts it to ascii if IDNA
func NormalizeFQDNAndPuny(domain string) string {
	domain = strings.ToLower(strings.TrimSpace(domain))

	if len(domain) == 0 {
		return ""
	}

	var err error
	punycode := ""
	if punycode, err = idna.ToASCII(domain); err != nil {
		punycode = domain
	}
	if punycode[len(punycode)-1] != '.' {
		return punycode + "."
	}

	return punycode
}

// RemoveFQDN makes a domain lower case and removes the FQDN dot if present
func RemoveFQDN(domain string) string {
	domain = strings.ToLower(strings.TrimSpace(domain))

	l := len(domain)

	if len(domain) == 0 {
		return ""
	}

	if domain[l-1:] == "." {
		return domain[0 : l-1]
	}

	return domain
}
