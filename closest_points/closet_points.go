package closest_points

import "math"

type Point struct {
	X int
	Y int
}

func ClosestPoints(input []Point) (Point, Point) {
	shortestDistance := math.MaxFloat64
	shortest1, shortest2 := 0, 0

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			//calculate distance between point-i and point-j
			// we can use the pythagorean formula C^2=A^2+B^2 to calculate the distance between two points.
			heightSq := math.Pow(float64(input[j].Y-input[i].Y), 2)
			widthSq := math.Pow(float64(input[j].X-input[i].X), 2)

			// C=SquareRoot(A^2+B^2)
			hypotenuse := math.Sqrt(float64(heightSq) + float64(widthSq))
			if hypotenuse < shortestDistance {
				shortest1, shortest2 = i, j
				shortestDistance = hypotenuse
			}
		}
	}

	return input[shortest1], input[shortest2]
}
