package main

import (
	"testing"
)

func BenchmarkCheckTing1Discard(b *testing.B) {
	_, cnt, err := convert("123m 22378p 345899s")
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < 10; i++ {
		checkTing1Discard(cnt)
	}
}

func TestAnalysis(t *testing.T) {
	var raw string
	raw = "11222333789s fa fa" //
	raw = "2355789p 356778s"
	raw = "4578999m 45p 11145s"
	raw = "123345567m 34p 345s"
	raw = "123m 2378p 34599s bei"
	raw = "2334567788s 5699p"
	raw = "123m 22378p 345899s"
	raw = "123m 22378p 345899s"
	raw = "1234m 22277p 3456s"
	raw = "123m 2378p 234999s"
	raw = "45689m 1189p 22256s" // 41775557 => 7800426
	raw = "12367m 123667p 556s"
	raw = "12378m 12378p 123s"
	raw = "123m 2378p 34599s bei" // 5180198 => 416416

	// http://blog.sina.com.cn/s/blog_7f78b76f0100s0nl.html
	raw = "11379m 347p 277s zhong zhong zhong"
	raw = "334578m 11468p 235s"
	raw = "478m 33588p 457899s"
	raw = "2233688m 1234p 378s"
	raw = "1233347m 23699p 88s"

	raw = "56778m 1245s 23388p"
	raw = "23m 22456p 1156899s"
	//raw = "2379m 22399s 23479p"
	analysis(raw)
}
