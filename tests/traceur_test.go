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

func TestCenterNorth(t *testing.T) {
	const nb_cote = 2
	const taille = 6

	traceur.Init()
	traceur.Down("black")

	traceur.North()

	var i int
	for i = 0; i < nb_cote; i++ {
		traceur.Forward(taille)
		traceur.Right()
	}
	traceur.Center()
	traceur.PrintCoords()

	if int(traceur.Y) != 0 || int(traceur.X) != 0 {
		t.Errorf("Forward failed - expected x=%d y=%d - got x=%d t=%d", 1, 0, int(traceur.X), int(traceur.Y))
	}
}

func TestCenterSouth(t *testing.T) {
	const nb_cote = 2
	const taille = 6

	traceur.Init()
	traceur.Down("black")

	traceur.South()

	var i int
	for i = 0; i < nb_cote; i++ {
		traceur.Forward(taille)
		traceur.Right()
	}
	traceur.Center()

	traceur.PrintCoords()

	if int(traceur.Y) != 0 || int(traceur.X) != 0 {
		t.Errorf("Forward failed - expected x=%d y=%d - got x=%d t=%d", 1, 0, int(traceur.X), int(traceur.Y))
	}
}

func TestCenter15(t *testing.T) {
	const nb_cote = 2
	const taille = 6

	traceur.Init()
	traceur.Down("black")

	traceur.Pivote(15)

	var i int
	for i = 0; i < nb_cote; i++ {
		traceur.Forward(taille)
		traceur.Right()
	}
	traceur.Center()

	traceur.PrintCoords()

	if int(traceur.Y) != 0 || int(traceur.X) != 0 {
		t.Errorf("Forward failed - expected x=%d y=%d - got x=%d t=%d", 1, 0, int(traceur.X), int(traceur.Y))
	}
}

func TestCenterDummy(t *testing.T) {
	const nb_cote = 2
	const taille = 6

	traceur.Init()
	traceur.Down("black")

	traceur.Center()

	traceur.PrintCoords()

	if int(traceur.Y) != 0 || int(traceur.X) != 0 {
		t.Errorf("Forward failed - expected x=%d y=%d - got x=%d t=%d", 1, 0, int(traceur.X), int(traceur.Y))
	}
}
