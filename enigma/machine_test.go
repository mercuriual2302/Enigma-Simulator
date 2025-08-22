package enigma

import "testing"

func TestEncodeReversible(t *testing.T) {
	m := NewMachine([]string{"I", "II", "III"}, []int{0, 0, 0}, "B", nil)
	plain := "HELLOWORLD"
	cipher := m.Encode(plain)

	m2 := NewMachine([]string{"I", "II", "III"}, []int{0, 0, 0}, "B", nil)
	decoded := m2.Encode(cipher)

	if decoded != plain {
		t.Fatalf("expected %s got %s", plain, decoded)
	}
}
