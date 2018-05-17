package main

import "testing"

func TestHello(t *testing.T) {
	esCorrecto := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to one person", func(t *testing.T) {
		got := Hello("YOU", "")
		want := "Hello, YOU"
		esCorrecto(t, got, want)
	})

	t.Run("saying hello to the world when no person supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		esCorrecto(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juancho", spanish)
		want := "Hola, Juancho"
		esCorrecto(t, got, want)
	})
    t.Run("in french", func(t *testing.T) {
        got := Hello("ADÈLE", french)
        want := "Bonjour, ADÈLE"
        esCorrecto(t, got, want)
    })
}
