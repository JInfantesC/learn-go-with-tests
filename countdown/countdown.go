package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

//La interfaz sleeper será implementada por todos las estructuras que tengan el metodo Sleep()
type Sleeper interface {
	Sleep()
}

// El Sleeper por defecto es una estructura con un método Sleep(). Implementa Sleeper.
type DefaultSleeper struct{}

//El método ejecuta time.Sleep(), el funcionamiento real.
func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}
	s.Sleep()
	fmt.Fprint(out, finalWord)
}
