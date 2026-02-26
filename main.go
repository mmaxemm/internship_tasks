package main

import (
	"fmt"
	"flag"

	"github.com/mmaxemm/internship_tasks/interfaces"
)
func main() {
	shapeNamePtr := flag.String("shape", "no", "")
	widthPtr := flag.Int("width", 0, "is skipped if -shape=circle")
	heightPtr := flag.Int("height", 0, "is skipped if -shape=circle")
	radiusPtr := flag.Int("radius", 0, "is skipped if -shape=rectangle")
	flag.Parse()
	if *shapeNamePtr != "circle" && *shapeNamePtr != "rectangle" {
		fmt.Println("-shape flag must be either circle or rectangle")
		return
	}
	
	var a interfaces.Shape
	switch *shapeNamePtr {
	case "circle":
		a = interfaces.Circle{Radius: float64(*radiusPtr)}
	case "rectangle":
		a = interfaces.Rectangle{Width: float64(*widthPtr), Height: float64(*heightPtr)}
	}
	fmt.Printf("Area is %f", a.Area())
}
