package structs

type Retangle struct {
	width  float64
	height float64
}

func Perimeter(retangle Retangle) float64 {
	return 2 * (retangle.width + retangle.height)
}

func Area(retangle Retangle) float64 {
	return retangle.width * retangle.height
}
