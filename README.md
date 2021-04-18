# traceur

Ensemble de fonctions pour effectuer des tracés géométriques
sur GoSpace.

## Exemple d'utilisation :

Un programme à compléter : https://goplay.space/#SIxGSHoHBIW

## Documentation :
package traceur // import "github.com/univ-tlse2/traceur"


### FUNCTIONS

`func Color(col string)`
    Color est un synonyme de Down

`func Down(col string)`
    Down change la couleur du stylet et le met en contact avec la feuille

`func East()`
    East dirige le stylet vers l'est

`func Forward(step float64)`
    Forward avance le stylet de _step_ pas dans la direction courante

`func Init()`
    Init initialise l'environnement du robot

`func Left()`
    Left tourne le stylet de 90° vers la gauche

`func North()`
    North dirige le stylet vers le nord

`func Pivote(angle int)`
    Pivote tourne le stylet de _angle_ degrés vers la droite

`func Right()`
    Right tourne le stylet de 90° vers la droite

`func Say(mess string)`
    Say affiche une bulle contenant le message _mess_

`func South()`
    South dirige le stylet vers le sud

`func Step()`
    Step avance le stylet d'un pas dans la direction courante

`func Up()`
    Up lève le stylet

`func West()`
    West dirige le stylet vers l'ouest

