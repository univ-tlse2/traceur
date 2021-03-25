package main

import (
	"github.com/univ-tlse2/traceur"
)

func main() {
	const nb_cote = 4
	const taille = 6

	traceur.PrintCoords()

	traceur.Init()
	traceur.Down("black")

	var i int
	for i = 0; i < 2; i++ {
		traceur.Forward(taille)
		traceur.Right()
		traceur.PrintCoords()
	}
	traceur.Center()
	traceur.PrintCoords()
}
