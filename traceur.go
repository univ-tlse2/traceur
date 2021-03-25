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
