package traceur

import (
	"fmt"
	"math"
)

var (
	// direction en radiant
	direction_rad float64 = 0
	x             float64 = 0
	y             float64 = 0
)

// getDirection récupère la direction dans l'unité demandé "deg" ou "rad"
func getDirection(unite string) float64 {
	var res float64
	if unite == "rad" {
		res = direction_rad
	} else {
		res = direction_rad * 180 / math.Pi
	}
	return res
}

// toDeg converti les radian en degré
func toDeg(rad float64) float64 {
	return rad * 180 / math.Pi
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
		if x > 0 {
			angle_to_center_rad = -math.Pi / 2
		} else {
			angle_to_center_rad = math.Pi / 2
		}
	} else {
		if y > 0 {
			angle_to_center_rad = -direction_rad + math.Pi + math.Atan(x/y)
		} else {
			angle_to_center_rad = -direction_rad + math.Atan(x/y)
		}
	}
	return angle_to_center_rad
}

// Init initialise l'environnement du robot
func Init() { fmt.Println("draw mode") }

// Forward avance le stylet de step pas dans la direction courante
func Forward(step float64) {
	fmt.Printf("forward %f\n", step)
	x += step * math.Sin(getDirection("rad"))
	y += step * math.Cos(getDirection("rad"))
}

// Step avance le stylet d'un pas dans la direction courante
func Step() { fmt.Println("forward 1") }

// Right tourne le stylet de 90° vers la droite
func Right() {
	fmt.Println("right")
	direction_rad += math.Pi / 2
}

// Left tourne le stylet de 90° vers la gauche
func Left() {
	fmt.Println("left")
	direction_rad -= math.Pi / 2
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
	direction_rad += toRad(float64(angle))
}

// North aller au nord
func North() {
	Pivote(-int(toDeg(direction_rad)))
}

// South aller au sud
func South() {
	Pivote(180 - int(toDeg(direction_rad)))
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
	var angle_to_center_rad float64
	var distance_to_center float64
	// a partir des coordonnées trouver la distance vers le centre
	distance_to_center = math.Sqrt(x*x + y*y)

	// tourner se réoriente vers le nord
	North()
	// a partir des coordonnées, trouver l'angle vers le centre
	angle_to_center_rad = angle_to_center(x, y)
	Pivote(int(toDeg(angle_to_center_rad)))
	Forward(distance_to_center)
	North()
}

// GetCoords renvoi les coordonnées du traceur
func GetCoords() map[string]float64 {
	coords := map[string]float64{
		"x":         x,
		"y":         y,
		"direction": toDeg(direction_rad),
	}
	return coords
}

// PrintCoords affiche les coordonnées du traceur
func PrintCoords() {
	fmt.Println("----")
	fmt.Printf("Direction : %f°\n", toDeg(direction_rad))
	fmt.Printf("X : %f\n", x)
	fmt.Printf("Y : %f\n", y)
	fmt.Println("----")
}
