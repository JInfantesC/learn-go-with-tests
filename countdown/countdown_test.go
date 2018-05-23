package main

import (
	"bytes"
	"testing"
)

//SpySleeper es una estructura que implementa la interfaz Sleeper.
type SpySleeper struct {
	Calls int
}

//El método Sleep() de Spy simula time.Sleep(). De este modo el test funciona rápidamente sin esperas.
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
