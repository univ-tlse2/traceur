package figures

import . "github.com/univ-tlse2/traceur"

// DrawRect dessine un rectangle de hauteur x largeur pas avec un contour de la couleur indiquée.
// Son coin inférieur gauche est à la position initiale du stylet.
// Après le tracé, le stylet est orienté vers le nord.
func DrawRect(hauteur, largeur float64, couleur string) {
	//North()
	Color(couleur)
	for cpteur := 1; cpteur <= 2; cpteur++ {
		// Trace un demi-rectangle
		Forward(hauteur)
		Right()
		Forward(largeur)
		Right()
	}
	//North()
}

// DrawSquare dessine un carré de nbPas de côté avec un contour de la couleur indiquée.
// Son coin inférieur gauche est à la position initiale du stylet.
// Après le tracé, le stylet est orienté vers le nord.
func DrawSquare(nbPas float64, couleur string) {
	DrawRect(nbPas, nbPas, couleur)
}

// DrawTriangle dessine un triangle équilatéral avec des côtés de nbPas et un contour de la couleur indiquée.
// Son coin inférieur gauche est à la position initiale du stylet.
// Après le tracé, le stylet est orienté vers le nord.
func DrawTriangle(nbPas float64, couleur string) {
	//West()
	Left()
	Color(couleur)
	for nbCotes := 1; nbCotes <= 3; nbCotes++ {
		Pivote(180 - 60)
		Forward(nbPas)
	}
	//North()
}
