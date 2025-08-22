package enigma

import "strings"

// Rotor represents an Enigma rotor.
type Rotor struct {
	wiring   string
	notch    rune
	position int
	ring     int
}

// NewRotor creates a rotor with wiring, notch, position and ring setting.
func NewRotor(wiring string, notch rune, position, ring int) *Rotor {
	return &Rotor{wiring: wiring, notch: notch, position: position % 26, ring: ring % 26}
}

// Step rotates the rotor by one position. It returns true if the rotor hits its notch.
func (r *Rotor) Step() bool {
	r.position = (r.position + 1) % 26
	return r.position == int(r.notch-'A')
}

// AtNotch reports if the rotor is currently at its notch position.
func (r *Rotor) AtNotch() bool {
	return r.position == int(r.notch-'A')
}

// Forward enciphers a value from right to left through the rotor.
func (r *Rotor) Forward(c int) int {
	offset := (c + r.position - r.ring + 26) % 26
	out := int(r.wiring[offset] - 'A')
	return (out - r.position + r.ring + 26) % 26
}

// Backward enciphers a value from left to right through the rotor.
func (r *Rotor) Backward(c int) int {
	offset := (c + r.position - r.ring + 26) % 26
	out := strings.IndexRune(r.wiring, rune('A'+offset))
	return (out - r.position + r.ring + 26) % 26
}
