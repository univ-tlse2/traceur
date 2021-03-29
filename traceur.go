package traceur

import "fmt"

// Angle pour revenir au Nord (en degrés)
var AngleToReturnToNorth int

// Init initialise l'environnement du robot
func Init() {
	fmt.Println("draw mode")
	AngleToReturnToNorth = 0
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
	AngleToReturnToNorth -= 90
}

// Left tourne le stylet de 90° vers la gauche
func Left() {
	fmt.Println("left")
	AngleToReturnToNorth += 90
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
	AngleToReturnToNorth -= angle
}

// North dirige le stylet vers le nord
func North() {
	if AngleToReturnToNorth < 0 {
		fmt.Printf("left %d\n", -AngleToReturnToNorth)
	} else {
		fmt.Printf("right %d\n", AngleToReturnToNorth)
	}
	AngleToReturnToNorth = 0
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
