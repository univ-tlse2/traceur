# traceur/figures

Ensemble de fonctions pour effectuer des tracés de figures
sur GoSpace.

## Exemple d'utilisation :

Un programme à compléter :
https://goplay.space/#Y5_G3L32h_g

## Documentation :
package figures // import "github.com/univ-tlse2/traceur/figures"

### FUNCTIONS

`func DrawRect(hauteur, largeur float64, couleur string)`
    DrawRect dessine un rectangle de _hauteur_ x _largeur_ pas avec un contour de la
    _couleur_ indiquée. Le dessin commence à la position du stylet, dans la
direction courante, en commençant par la hauteur et dans le sens horaire. À la fin du tracé, la direction
du stylet est celle de la largeur finale.

`func DrawSquare(nbPas float64, couleur string)`
    DrawSquare dessine un carré de _nbPas_ de côté avec un contour de la _couleur_
    indiquée. Le dessin commence à la position du stylet, dans la
direction courante. Le tracé s'effectue dans le sens
horaire. À la fin du tracé, la direction
du stylet est celle du dernier côté

`func DrawTriangle(nbPas float64, couleur string)`
    DrawTriangle dessine un triangle équilatéral avec des côtés de _nbPas_ et un
    contour de la _couleur_ indiquée. Le dessin commence à la position du stylet, dans la
direction courante, en commençant par le premier côté. Le tracé s'effectue dans le sens 
horaire. À la fin du tracé, la direction
du stylet est celle de la base.

