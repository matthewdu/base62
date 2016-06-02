package base62

import (
	"testing"
)

var tests = []struct {
	n int64
	s string
}{
	{0, "0"},
	{10, "a"},
	{61, "Z"},
	{62, "10"},
	{630, "aa"},
	{3781504209452600, "hjNv8tS3K"},
}

func TestEncode(t *testing.T) {
	for _, tt := range tests {
		actual := Encode(tt.n)
		if actual != tt.s {
			t.Errorf("Encode(%d): expected %s, actual %s", tt.n, tt.s, actual)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, tt := range tests {
		if actual, err := Decode(tt.s); actual != tt.n || err != nil {
			t.Error(err)
			t.Errorf("Decode(%s): expected %d, actual %d", tt.s, tt.n, actual)
		}
	}

	// Should return error on invalid token
	if _, err := Decode("A-+?"); err == nil {
		t.Errorf("Token A-+? should have caused error")
	}
}

func BenchmarkEncode(b *testing.B) {
	encodeValues := [6]int64{7, 117, 343, 2401, 823543, 300124211606973}
	for i := 0; i < b.N; i++ {
		Encode(encodeValues[i%6])
	}
}

func BenchmarkDecode(b *testing.B) {
	decodeValues := [6]string{"1jbk3jb", "JAI3j1NA", "0", "aaa", "0ija", "12345K"}
	for i := 0; i < b.N; i++ {
		Decode(decodeValues[i%6])
	}
}
