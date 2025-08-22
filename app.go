package main

import (
	"context"
	"fmt"

	"github.com/mercuriual2302/Enigma-Simulator/enigma"
)

// App struct
type App struct {
	ctx     context.Context
	machine *enigma.Machine
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{machine: enigma.NewMachine([]string{"I", "II", "III"}, []int{0, 0, 0}, "B", nil)}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Encrypt enciphers the supplied text using the Enigma machine.
func (a *App) Encrypt(text string) string {
	return a.machine.Encode(text)
}

// Configure resets the machine with new settings.
func (a *App) Configure(rotors []string, positions []int, reflector string, plugs []string) {
	a.machine = enigma.NewMachine(rotors, positions, reflector, plugs)
}
