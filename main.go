package main

import "github.com/Jullentuss/course/course"

func main() {
	Go := course.Course{
		"Go desde 0",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()
}
