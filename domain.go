package domainnames

import (
	"strings"

	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
)

// DomainAndRoot checks for a valid look domain and returns a normalised version and
// the publicly registrable root for the domain.
func DomainAndRoot(domain string) (string, string, string, error) {
	domain = RemoveFQDN(domain)

	parts := strings.Split(domain, ".")
	if l := len(parts); l < 2 {
		return domain, "", "", ErrMalformedDomain
	}
	ps, icann := publicsuffix.PublicSuffix(domain)
	if !icann {
		return domain, "", "", ErrNotAnIcannDomain
	}

	i := strings.LastIndex(domain, ps)
	if i < 1 {
		return domain, "", "", ErrMalformedPublicSuffix
	}

	host := domain[0 : i-1]
	i = strings.LastIndex(host, ".")
	if i != -1 {
		host = host[i+1:]
	}
	if host == "" {
		return domain, "", "", ErrMalformedDomain
	}

	return domain, host, ps, nil
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
