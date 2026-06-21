package util

func Round2Dec(value float64) float64 {
	return float64(int(value*100)) / 100
}
