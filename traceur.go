// Package traceur fournit un ensemble de fonctions pour effectuer simplement
// des tracés géométriques sur goplay.space (https://goplay.space/)
package traceur

import "fmt"

// Angle pour revenir au Nord (en degrés)
var angleToReturnToNorth int

// init initialise l'environnement du robot (appelée automatiquement au chargement du package)
func init() {
	fmt.Println("draw mode")
	angleToReturnToNorth = 0
}

// Forward avance le stylet de step pas dans la direction courante
func Forward(step float64) {
	fmt.Printf("forward %f\n", step)
}

// Step avance le stylet d'un pas dans la direction courante
func Step() { fmt.Println("forward 1") }

// Right tourne le stylet de 90° vers la droite
func Right() {
	fmt.Println("right")
	angleToReturnToNorth -= 90
}

// Left tourne le stylet de 90° vers la gauche
func Left() {
	fmt.Println("left")
	angleToReturnToNorth += 90
}

// Say affiche une bulle contenant le message mess
func Say(mess string) { fmt.Printf("say %s\n", mess) }

// Up lève le stylet
func Up() { fmt.Println("color off") }

// Down change la couleur du stylet et le met en contact avec la feuille
func Down(col string) { fmt.Printf("color %s\n", col) }

// Color est un synonyme de Down
func Color(col string) { fmt.Printf("color %s\n", col) }

// Pivote tourne le stylet de angle degrés vers la droite
func Pivote(angle int) {
	fmt.Printf("right %d\n", angle)
	angleToReturnToNorth -= angle
}

// North dirige le stylet vers le nord
func North() {
	if angleToReturnToNorth < 0 {
		fmt.Printf("left %d\n", -angleToReturnToNorth)
	} else {
		fmt.Printf("right %d\n", angleToReturnToNorth)
	}
	angleToReturnToNorth = 0
}

// South dirige le stylet vers le sud
func South() {
	North()
	Pivote(180)
}

// East dirige le stylet vers l'est
func East() {
	North()
	Right()
}

// West dirige le stylet vers l'ouest
func West() {
	North()
	Left()
}
