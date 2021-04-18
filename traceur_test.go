package traceur

// Tests des fonctions North(), South(), East() et West()
import "testing"

func TestInit(t *testing.T) {
	// Vérifie que le traceur est dirigé vers le nord
	Init()
	got := angleToReturnToNorth
	if got != 0 {
		t.Errorf("Got : %d, Expected: %d", got, 0)
	}
}

func TestNorthWithCardinalPoints(t *testing.T) {
	South()
	East()
	West()
	North()
	got := angleToReturnToNorth
	if got != 0 {
		t.Errorf("Got : %d, Expected: %d", got, 0)
	}
}

func TestNorthWithPivoteAndRight(t *testing.T) {
	Pivote(30)
	Right()
	Pivote(-10)
	North()
	got := angleToReturnToNorth
	if got != 0 {
		t.Errorf("Got : %d, Expected: %d", got, 0)
	}
}
