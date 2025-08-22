package enigma

// Plugboard swaps pairs of letters before and after rotor processing.
type Plugboard struct {
	wiring map[int]int
}

// NewPlugboard creates a plugboard with the given pairs like ["AB", "CD"].
func NewPlugboard(pairs []string) *Plugboard {
	p := &Plugboard{wiring: map[int]int{}}
	for i := 0; i < 26; i++ {
		p.wiring[i] = i
	}
	for _, pair := range pairs {
		if len(pair) != 2 {
			continue
		}
		a := int(pair[0] - 'A')
		b := int(pair[1] - 'A')
		p.wiring[a] = b
		p.wiring[b] = a
	}
	return p
}

// Swap returns the swapped value according to the plugboard configuration.
func (p *Plugboard) Swap(c int) int {
	if p == nil {
		return c
	}
	return p.wiring[c]
}
