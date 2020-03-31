package utils

import (
	"testing"
)

func TestIsOnion(t *testing.T) {
	// v2 onion
	test := IsOnion("maschqwtr4lqt2pj.onion")
	if !test {
		t.Errorf("%s should have matched!", "maschqwtr4lqt2pj.onion")
	}

	test = IsOnion("subdomain.lots.of.maschqwtr4lqt2pj.onion")
	if !test {
		t.Errorf("%s should have matched!", "subdomain.lots.of.maschqwtr4lqt2pj.onion")
	}
	// v3 onion
	test := IsOnion("2pr59arwnnfopb5yu4a3p62u7x3hnszzwn56ezs48e5r2teo56ybis8w.onion")
	if !test {
		t.Errorf("%s should have matched!", "2pr59arwnnfopb5yu4a3p62u7x3hnszzwn56ezs48e5r2teo56ybis8w.onion")
	}

	test = IsOnion("subdomain.lots.of.2pr59arwnnfopb5yu4a3p62u7x3hnszzwn56ezs48e5r2teo56ybis8w.onion")
	if !test {
		t.Errorf("%s should have matched!", "subdomain.lots.of.2pr59arwnnfopb5yu4a3p62u7x3hnszzwn56ezs48e5r2teo56ybis8w.onion")
	}
	
	test = IsOnion("thisis9notavalidonion.onion")
	if test {
		t.Errorf("%s should not have matched!", "thisis9notavalidonion.onion")
	}

	test = IsOnion("maschqwtr4lqt2pj.com")
	if test {
		t.Errorf("%s should not have matched!", "maschqwtr4lqt2pj.com")
	}

	test = IsOnion("2pr59arwnnfopb5yu4a3p62u7x3hnszzwn56ezs48e5r2teo56ybis8w.com")
	if test {
		t.Errorf("%s should not have matched!", "2pr59arwnnfopb5yu4a3p62u7x3hnszzwn56ezs48e5r2teo56ybis8w.com")
	}
	
	test = IsOnion("www.thisisnotanonionitistoolong.onion")
	if test {
		t.Errorf("%s should not have matched!", "www.thisisnotanonionitistoolong.onion")
	}

}
