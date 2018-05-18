package utils

import "errors"

var (
	// Something is wrong with the domain
	ErrMalformedDomain = errors.New("malformed-domain")

	// Is an known suffix but not Icann
	ErrNotAnIcannDomain = errors.New("non-icann-domain")

	// Is a known public suffix but not currect
	ErrMalformedPublicSuffix = errors.New("malformed-public-suffix")
)
