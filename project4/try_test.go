package main

import "testing"

func Test_Repeat(t *testing.T) {
	t.Helper()
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but get %q", expected, repeated)
	}
}

func BenchmarkRepeatabc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
