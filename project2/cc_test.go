// test_hello.go
package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		got := sayHello("Chris", "english")
		want := "Hello Chris"
		assertMessage(t, got, want)
	})
	t.Run("say hello world when empty string profided,with spanish",
		func(t *testing.T) {
			got := sayHello("", "spanish")
			want := "Hola world"
			assertMessage(t, got, want)
		})
	t.Run("say hello to Tom, with french",
		func(t *testing.T) {
			got := sayHello("Tom", "french")
			want := "Bonjour Tom"
			assertMessage(t, got, want)
		})
}
