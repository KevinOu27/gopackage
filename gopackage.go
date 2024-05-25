package trimmean

import "sort"

func TrimmedMean(data []float64, trim ...float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sort.Float64s(data)

	var lowerTrim, upperTrim float64
	if len(trim) == 1 {
		lowerTrim, upperTrim = trim[0], trim[0]
	} else if len(trim) >= 2 {
		lowerTrim, upperTrim = trim[0], trim[1]
	}
	trimLow := int(lowerTrim * float64(len(data)))
	trimHigh := int(upperTrim * float64(len(data)))

	trimmedData := data[trimLow : len(data)-trimHigh]
	sum := 0.0
	for _, v := range trimmedData {
		sum += v
	}
	return sum / float64(len(trimmedData))
}
