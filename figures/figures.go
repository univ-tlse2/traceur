// Package figures ajoute des fonctions de tracés de rectangles, carrés et triangles
// au package traceur
package figures

import . "github.com/univ-tlse2/traceur"

// DrawRect dessine un rectangle de hauteur x largeur pas avec un contour de la couleur indiquée.
// Son coin inférieur gauche est à la position initiale du stylet.
func DrawRect(hauteur, largeur float64, couleur string) {
	Color(couleur)
	for cpteur := 1; cpteur <= 2; cpteur++ {
		// Trace un demi-rectangle
		Forward(hauteur)
		Right()
		Forward(largeur)
		Right()
	}
}

// DrawSquare dessine un carré de nbPas de côté avec un contour de la couleur indiquée.
// Son coin inférieur gauche est à la position initiale du stylet.
func DrawSquare(nbPas float64, couleur string) {
	DrawRect(nbPas, nbPas, couleur)
}

// DrawTriangle dessine un triangle équilatéral avec des côtés de nbPas et un contour de la couleur indiquée.
// Son coin inférieur gauche est à la position initiale du stylet.
func DrawTriangle(nbPas float64, couleur string) {
	Left()
	Color(couleur)
	for nbCotes := 1; nbCotes <= 3; nbCotes++ {
		Pivote(180 - 60)
		Forward(nbPas)
	}
}
