package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' in Spanish", func(t *testing.T) {
		got := Hello("Serhio", "Spanish")
		want := "Hola, Serhio"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' in French", func(t *testing.T) {
		got := Hello("Serhio", "French")
		want := "Bonjur, Serhio"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
