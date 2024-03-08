package Sensor
import (
	"math/rand"
)

func SensorData() map[string]int {
	data := map[string]int{
		"freezer1":  - rand.Intn(30),
		"freezer2": rand.Intn(30),
		"Geladeira1": - rand.Intn((2)),
		"Geladeira2":  rand.Intn(2),
	}
	return data
}