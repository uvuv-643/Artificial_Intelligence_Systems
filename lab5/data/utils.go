package data

import (
	"math"
	"slices"
)

func GetVariationSeries(dataset Dataset) map[string][]float64 {
	variationSeries := make(map[string][]float64)
	for key := range dataset.Data[0] {
		variationSeries[key] = make([]float64, 0)
	}
	for _, row := range dataset.Data {
		for key, value := range row {
			variationSeries[key] = append(variationSeries[key], value)
		}
	}
	for key := range dataset.Data[0] {
		slices.Sort(variationSeries[key])
	}
	return variationSeries
}

func R2(initial []float64, predicted []float64) float64 {
	errSum := 0.0
	initialMean := 0.0
	variance_ := 0.0
	for i, value := range initial {
		errSum += math.Pow(value-predicted[i], 2)
		initialMean += value / float64(len(initial))
	}
	for _, value := range initial {
		variance_ += math.Pow(value-initialMean, 2)
	}

	//fmt.Println(float64(len(predicted)), Variance(dataset)[target])
	return 1 - errSum/variance_
}
