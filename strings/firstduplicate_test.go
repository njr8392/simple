package main

import "testing"

func TestFirstDup(t * testing.T){
	m := map[rune]rune{
		FirstDup("adsffuiola"): 'f',
		FirstDup("strings"): 's',
		FirstDup(""): ' ',
	}
	for got, want := range m{
		if got != want{
			t.Errorf("got %q, want %q", got, want)
		}
	}
}
