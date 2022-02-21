package solution

import "math"

type SidesNum int

const (
	SidesTriangle SidesNum = 3
	SidesSquare   SidesNum = 4
	SidesCircle   SidesNum = 0
)

func CalcSquare(sideLen float64, sidesNum SidesNum) float64 {
	if sidesNum == 3 {
		return (sideLen * sideLen * math.Sqrt(3) / 4)
	} else if sidesNum == 4 {
		return math.Pow(sideLen, 2)
	} else if sidesNum == 0 {
		return math.Pi * math.Pow(sideLen, 2)
	}
	return 0
}
