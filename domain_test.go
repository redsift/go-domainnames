package utils

import (
	"fmt"
	"testing"
)

func TestVerifyDomainRoots(t *testing.T) {
	test := []struct {
		d, r string
	}{
		{"a.com.foo.dyndns.com", "dyndns.com"},
		{" bluecȯat.com.com.foo.Dyndns.com ", "dyndns.com"},
		{" a.com.foo.Dyndns.com ", "dyndns.com"},
		{"abc.dyi.xyz.redsift.co.uk", "redsift.co.uk"},
		{"Xyz.redsift.co.uk", "redsift.co.uk"},
		{"redsift.co.uk", ""},
		{"Flowmobile.co.uk.", ""},
		{"redsift.com", ""},
		{"redsift.pizza", ""},
		{"bluecȯat.com", ""},
		{"fql.bluecȯat.com", "bluecȯat.com"},
	}

	for i, ts := range test {
		t.Run(fmt.Sprintf("%d-%s", i, ts.d), func(t *testing.T) {
			if r, err := VerifyDomainFormatAndRoot(ts.d); err != nil {
				t.Error(err)
			} else if r != ts.r {
				t.Error(r, "!=", ts.r)
			}
		})
	}
}

func TestVerifyDomainFormatErrs(t *testing.T) {
	for i, ts := range []string{"co.uk", ".co.uk", "foo.dyndns.org", "broken.notatoplevelg", "broken"} {
		t.Run(fmt.Sprintf("%d-%s", i, ts), func(t *testing.T) {
			if a, err := VerifyDomainFormatAndRoot(ts); err == nil {
				t.Error("did not fail", a)
			} else if a != "" {
				t.Error("root was not empty")
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
