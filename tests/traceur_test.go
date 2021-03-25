package traceur

import (
	"testing"

	"github.com/univ-tlse2/traceur"
)

// go test -v pour voir le Printf
func TestForward(t *testing.T) {
	traceur.Init()
	traceur.Down("black")
	traceur.Forward(1)
	if int(traceur.Y) != 1 || int(traceur.X) != 0 {
		t.Errorf("Forward failed - expected x=%d y=%d - got x=%d t=%d", 1, 0, int(traceur.X), int(traceur.Y))
	}
}
