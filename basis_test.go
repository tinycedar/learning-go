package basis

import (
	"log"
	"math"
	"testing"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}

type square struct {
	length float64
}

func (s *square) area() float64 {
	return math.Pow(s.length, 2)
}

func printArea(shape shape) {
	log.Println("shape: ", shape)
	log.Println("area: ", shape.area())
}

func TestInterface(t *testing.T) {
	printArea(&rectangle{width: 5, height: 9})
	printArea(&square{length: 8})
}
