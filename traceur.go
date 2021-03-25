package traceur

import (
	"fmt"
	"math"
)

var (
	// direction en radiant
	Direction_rad float64 = 0
	X             float64 = 0
	Y             float64 = 0
)

// getDirection récupère la direction dans l'unité demandé "deg" ou "rad"
func GetDirection(unite string) float64 {
	var res float64
	if unite == "rad" {
		res = Direction_rad
	} else {
		res = Direction_rad * 180 / math.Pi
	}
	return res
}

// toDeg converti les radian en degré
func toDeg(rad float64) int {
	deg := rad * 180 / math.Pi
	return int(deg) % 360
}

// toDeg converti les degré en radiant
func toRad(deg float64) float64 {
	return deg * math.Pi / 180
}

// angle_to_center trouve l'angle permettant de s'orienter vers le centre
// précondition : être orienté vers le nord
// TODO : integrer la précondition dans la fonction
func angle_to_center(x float64, y float64) float64 {
	var angle_to_center_rad float64
	// tan() = x/y
	if y == 0 {
		if X > 0 {
			angle_to_center_rad = -math.Pi / 2
		} else {
			angle_to_center_rad = math.Pi / 2
		}
	} else {
		if y > 0 {
			angle_to_center_rad = -Direction_rad + math.Pi + math.Atan(x/y)
		} else {
			angle_to_center_rad = -Direction_rad + math.Atan(x/y)
		}
	}
	return angle_to_center_rad
}

// Init initialise l'environnement du robot
func Init() {
	fmt.Println("draw mode")
	X = 0
	Y = 0
	Direction_rad = 0
}

// Forward avance le stylet de step pas dans la direction courante
func Forward(step float64) {
	fmt.Printf("forward %f\n", step)
	X += step * math.Sin(GetDirection("rad"))
	Y += step * math.Cos(GetDirection("rad"))
}

// Step avance le stylet d'un pas dans la direction courante
func Step() { fmt.Println("forward 1") }

// Right tourne le stylet de 90° vers la droite
func Right() {
	fmt.Println("right")
	Direction_rad += math.Pi / 2
}

// Left tourne le stylet de 90° vers la gauche
func Left() {
	fmt.Println("left")
	Direction_rad -= math.Pi / 2
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
	Direction_rad += toRad(float64(angle))
}

// North aller au nord
func North() {
	Pivote(-int(toDeg(Direction_rad)))
}

// South aller au sud
func South() {
	Pivote(180 - int(toDeg(Direction_rad)))
}

// East aller à l'est
func East() {
	North()
	Right()
}

// West aller à l'ouest
func West() {
	North()
	Left()
}

// Center retourne au centre du dessin
func Center() {
	if int(X) != 0 || int(Y) != 0 {
		var (
			angle_to_center_rad float64
			distance_to_center  float64
		)
		// a partir des coordonnées trouver la distance vers le centre
		distance_to_center = math.Sqrt(X*X + Y*Y)

		// tourner se réoriente vers le nord
		North()
		// a partir des coordonnées, trouver l'angle vers le centre
		angle_to_center_rad = angle_to_center(X, Y)
		Pivote(int(toDeg(angle_to_center_rad)))
		Forward(distance_to_center)
	}
	North()
}

// PrintCoords affiche les coordonnées du traceur
func PrintCoords() {
	fmt.Printf("  ###### Coordonnées ########\n")
	fmt.Printf("  ## Direction : %d°\n", toDeg(Direction_rad))
	fmt.Printf("  ## X : %f\n", X)
	fmt.Printf("  ## Y : %f\n", Y)
	fmt.Printf("  ###########################\n")
}
