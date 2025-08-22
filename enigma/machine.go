package enigma

import (
	"strings"
)

// Machine represents a complete Enigma machine with three rotors, reflector and plugboard.
type Machine struct {
	rotors    []*Rotor // ordered left to right: 0-left,2-right
	reflector string
	plugboard *Plugboard
}

// Default rotor and reflector wirings.
var rotorWiring = map[string]struct {
	wiring string
	notch  rune
}{
	"I":   {"EKMFLGDQVZNTOWYHXUSPAIBRCJ", 'Q'},
	"II":  {"AJDKSIRUXBLHWTMCQGZNPYFVOE", 'E'},
	"III": {"BDFHJLCPRTXVZNYEIWGAKMUSQO", 'V'},
	"IV":  {"ESOVPZJAYQUIRHXLNFTGKDCMWB", 'J'},
	"V":   {"VZBRGITYUPSDNHLXAWMJQOFECK", 'Z'},
}

var reflectorWiring = map[string]string{
	"B": "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	"C": "FVPJIAOYEDRZXWGCTKUQSBNMHL",
}

// NewMachine creates a new Enigma machine with given rotor names, initial positions,
// reflector and plugboard pairs.
func NewMachine(rotorNames []string, positions []int, reflector string, plugs []string) *Machine {
	if len(rotorNames) != 3 || len(positions) != 3 {
		panic("need three rotors and positions")
	}
	rotors := make([]*Rotor, 3)
	for i := 0; i < 3; i++ {
		cfg, ok := rotorWiring[rotorNames[i]]
		if !ok {
			panic("unknown rotor " + rotorNames[i])
		}
		rotors[i] = NewRotor(cfg.wiring, cfg.notch, positions[i], 0)
	}
	return &Machine{
		rotors:    rotors,
		reflector: reflectorWiring[reflector],
		plugboard: NewPlugboard(plugs),
	}
}

// stepRotors implements rotor stepping including double-step.
func (m *Machine) stepRotors() {
	if m.rotors[1].AtNotch() {
		m.rotors[0].Step()
		m.rotors[1].Step()
	}
	if m.rotors[2].AtNotch() {
		m.rotors[1].Step()
	}
	m.rotors[2].Step()
}

// Encode enciphers text using the machine's current state.
func (m *Machine) Encode(text string) string {
	var sb strings.Builder
	text = strings.ToUpper(text)
	for _, r := range text {
		if r < 'A' || r > 'Z' {
			continue
		}
		m.stepRotors()
		c := int(r - 'A')
		c = m.plugboard.Swap(c)
		// through rotors right to left
		for i := 2; i >= 0; i-- {
			c = m.rotors[i].Forward(c)
		}
		// reflector
		c = int(m.reflector[c] - 'A')
		// back through rotors left to right
		for i := 0; i < 3; i++ {
			c = m.rotors[i].Backward(c)
		}
		c = m.plugboard.Swap(c)
		sb.WriteByte(byte('A' + c))
	}
	return sb.String()
}
