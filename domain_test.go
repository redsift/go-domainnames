package domainnames

import (
	"fmt"
	"testing"
)

func TestVerifyDomainRoots(t *testing.T) {
	test := []struct {
		d, r, p string // test that we find the registration and the public suffix
	}{
		{"a.com.foo.dyndns.com", "dyndns", "com"},
		{" bluecȯat.com.com.foo.Dyndns.com ", "dyndns", "com"},
		{" a.com.foo.Dyndns.com ", "dyndns", "com"},
		{"abc.dyi.xyz.redsift.co.uk", "redsift", "co.uk"},
		{"Xyz.redsift.co.uk", "redsift", "co.uk"},
		{"redsift.co.uk", "redsift", "co.uk"},
		{"Flowmobile.co.uk.", "flowmobile", "co.uk"},
		{"redsift.com", "redsift", "com"},
		{"redsift.pizza", "redsift", "pizza"},
		{"bluecȯat.com", "bluecȯat", "com"},
		{"something.pl", "something", "pl"},
		{"fql.bluecȯat.com", "bluecȯat", "com"},
	}

	for i, ts := range test {
		t.Run(fmt.Sprintf("%d-%s", i, ts.d), func(t *testing.T) {
			if _, r, ps, err := DomainAndRoot(ts.d); err != nil {
				t.Error(err)
			} else if r != ts.r {
				t.Error(r, "!=", ts.r)
			} else if ps != ts.p {
				t.Error(ps, "!=", ts.p)
			}
		})
	}
}

func TestVerifyNonIccan(t *testing.T) {
	test := []struct {
		d, r, p string // test that we find the registration and the public suffix
	}{
		{"mck.krakow.pl", "mck", "krakow.pl"},
		{"vmnxbironsp01.gsnet.corp", "gsnet", "corp"},
	}

	for i, ts := range test {
		t.Run(fmt.Sprintf("%d-%s", i, ts.d), func(t *testing.T) {
			if _, r, ps, err := DomainAndRoot(ts.d); err != ErrNotAnIcannDomain {
				t.Error(err)
			} else if r != ts.r {
				t.Error(r, "!=", ts.r)
			} else if ps != ts.p {
				t.Error(ps, "!=", ts.p)
			}
		})
	}
}

func TestVerifyDomainFormatErrs(t *testing.T) {
	for i, ts := range []string{"co.uk", ".co.uk", "foo.dyndns.org", "broken.notatoplevelg", "broken"} {
		t.Run(fmt.Sprintf("%d-%s", i, ts), func(t *testing.T) {
			if _, a, _, err := DomainAndRoot(ts); err == nil {
				t.Error("did not fail", a)
			} else {
				t.Log(err)
			}
		})
	}
}

func TestNormalizeFQDN(t *testing.T) {
	test := []struct {
		d, e string
	}{
		{"redsift.io", "redsift.io."},
		{"blue.redsift.io", "blue.redsift.io."},
		{"redsift.io.", "redsift.io."},
		{" Redsift.io. ", "redsift.io."},
		{" blue.Redsift.io. ", "blue.redsift.io."},
		{" Redsift.i. ", "redsift.i."},
		{" bluecȯat.com. ", "xn--bluecat-x2c.com."},
		{" bluecȯat.com ", "xn--bluecat-x2c.com."},
		{" xn--bluecat-x2c.com ", "xn--bluecat-x2c.com."},
	}

	for i, ts := range test {
		t.Run(fmt.Sprintf("%d-%s", i, ts.d), func(t *testing.T) {
			if r := NormalizeFQDNAndPuny(ts.d); r != ts.e {
				t.Error(r, "!=", ts.e)
			}
		})
	}
}

func TestRemoveFQDN(t *testing.T) {
	test := []struct {
		d, e string
	}{
		{"redsift.io", "redsift.io"},
		{"blue.redsift.io", "blue.redsift.io"},
		{"redsift.io.", "redsift.io"},
		{" Redsift.io. ", "redsift.io"},
		{" blue.Redsift.io. ", "blue.redsift.io"},
		{" Redsift.i. ", "redsift.i"},
		{" bluecȯat.com. ", "bluecȯat.com"},
		{" bluecȯat.com ", "bluecȯat.com"},
	}

	for i, ts := range test {
		t.Run(fmt.Sprintf("%d-%s", i, ts.d), func(t *testing.T) {
			if r := RemoveFQDN(ts.d); r != ts.e {
				t.Error(r, "!=", ts.e)
			}
		})
	}
}
