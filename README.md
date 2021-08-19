# traceur

Ensemble de fonctions pour effectuer des tracés géométriques
sur GoSpace.

## Exemple d'utilisation :

Un programme à compléter :
https://goplay.space/#7e_eWbs31sd

## Documentation :
package traceur // import "github.com/univ-tlse2/traceur"


### FUNCTIONS

`func Color(col string)`
    Color modifie la couleur du stylet (noir par défaut)

`func Down()`
    Down met le stylet en contact avec la feuille

`func East()`
    East dirige le stylet vers l'est

`func Forward(step float64)`
    Forward avance le stylet de step pas dans la direction courante

`func GetColor() string`
    GetColor renvoie la couleur de tracé courante

`func Left()`
    Left tourne le stylet de 90° vers la gauche

`func North()`
    North dirige le stylet vers le nord

`func Pivote(angle int)`
    Pivote tourne le stylet de angle degrés vers la droite

`func Right()`
    Right tourne le stylet de 90° vers la droite

`func Say(mess string)`
    Say affiche une bulle contenant le message mess

`func South()`
    South dirige le stylet vers le sud

`func Step()`
    Step avance le stylet d'un pas dans la direction courante

`func Up()`
    Up lève le stylet

`func West()`
    West dirige le stylet vers l'ouest

