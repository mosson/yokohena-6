package tree

import "testing"

func TestRead(t *testing.T) {
	a, b := read("28->10")
	if a != 28 {
		t.Errorf("expected 28, actual %v", a)
	}

	if b != 10 {
		t.Errorf("expected 10, actual %v", b)
	}
}

var tests = map[string]string{
	"5->2":   "mo",
	"28->10": "au",
	"1->1":   "me",
	"40->40": "me",
	"27->27": "me",
	"7->2":   "mo",
	"40->13": "mo",
	"9->3":   "mo",
	"4->1":   "mo",
	"1->3":   "da",
	"12->35": "da",
	"3->8":   "da",
	"6->19":  "da",
	"38->40": "si",
	"9->8":   "si",
	"4->2":   "si",
	"15->16": "si",
	"40->12": "au",
	"10->4":  "au",
	"21->5":  "au",
	"8->2":   "au",
	"3->5":   "ni",
	"11->39": "ni",
	"2->13":  "ni",
	"13->32": "ni",
	"14->22": "co",
	"40->34": "co",
	"5->8":   "co",
	"12->10": "co",
	"1->27":  "-",
	"8->1":   "-",
	"12->22": "-",
	"2->40":  "-",
	"32->31": "-",
	"13->14": "-",
}

func TestMo(t *testing.T) {
	if solve("5->2") != "mo" {
		t.Errorf("expected mo, actual %v", solve("5->2"))
	}
}

func TestAu(t *testing.T) {
	if solve("28->10") != "au" {
		t.Errorf("expected au, actual %v", solve("28->10"))
	}
}

func TestMe(t *testing.T) {
	if solve("1->1") != "me" {
		t.Errorf("expected me, actual %v", solve("1->1"))
	}
}

func TestAll(t *testing.T) {
	for k, v := range tests {
		if solve(k) != v {
			t.Errorf("%v: expected %v, actual %v", k, v, solve(k))
		}
	}
}

func TestConAll(t *testing.T) {
	for k, v := range tests {
		r := conSolve(k)
		if r != v {
			t.Errorf("%v: expected %v, actual %v", k, v, r)
		}
	}
}

func BenchmarkAll(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range tests {
			solve(k)
		}
	}
}

func BenchmarkConAll(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range tests {
			conSolve(k)
		}
	}
}
