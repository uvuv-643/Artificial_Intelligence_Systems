package data

import (
	"errors"
	"math"
)

var mean = make(map[string]float64)
var squaredMean = make(map[string]float64)
var variance = make(map[string]float64)
var standardDeviation = make(map[string]float64)
var quantile = make(map[string]float64)
var _min = make(map[string]float64)
var _max = make(map[string]float64)
var corr = make(map[string]map[string]float64)

func Mean(dataset Dataset) map[string]float64 {
	mean = make(map[string]float64)
	for _, row := range dataset.Data {
		for key, value := range row {
			mean[key] += value / float64(len(dataset.Data))
		}
	}
	return mean
}

func Variance(dataset Dataset) map[string]float64 {
	mean = Mean(dataset)
	variance = make(map[string]float64)
	for key := range dataset.Data[0] {
		for _, value := range dataset.Data {
			variance[key] += (value[key] - mean[key]) * (value[key] - mean[key]) / float64(len(dataset.Data))
		}
	}
	return variance
}

func StandardDeviation(dataset Dataset) map[string]float64 {
	variance = Variance(dataset)
	for key, value := range variance {
		standardDeviation[key] = math.Sqrt(value)
	}
	return standardDeviation
}

func Quantile(dataset Dataset, level float64) (map[string]float64, error) {
	variationSeries := GetVariationSeries(dataset)
	quantileIndex := float64(len(dataset.Data)) * level
	if int(quantileIndex) < len(dataset.Data) && int(quantileIndex) >= 0 {
		if quantileIndex-float64(int(quantileIndex)) == 0 {
			if int(quantileIndex)+1 < len(dataset.Data) {
				for key := range dataset.Data[int(quantileIndex)] {
					currentElement := variationSeries[key][int(quantileIndex)]
					nextElement := variationSeries[key][int(quantileIndex)+1]
					quantile[key] = (currentElement + nextElement) / 2
				}
			}
		} else {
			for key := range dataset.Data[int(quantileIndex)] {
				quantile[key] = variationSeries[key][int(quantileIndex)+1]
			}
		}
		return quantile, nil
	}
	return quantile, errors.New("wrong quantile level")
}

func Min(dataset Dataset) map[string]float64 {
	_min = make(map[string]float64)
	for k, v := range dataset.Data[0] {
		_min[k] = v
	}
	for _, row := range dataset.Data {
		for key, value := range row {
			_min[key] = math.Min(value, _min[key])
		}
	}
	return _min
}

func Max(dataset Dataset) map[string]float64 {
	_max = make(map[string]float64)
	for k, v := range dataset.Data[0] {
		_max[k] = v
	}
	for _, row := range dataset.Data {
		for key, value := range row {
			_max[key] = math.Max(value, _max[key])
		}
	}
	return _max
}

func Corr(dataset Dataset) map[string]map[string]float64 {
	mean = Mean(dataset)
	for key1 := range dataset.Data[0] {
		corr[key1] = make(map[string]float64)
		for key2 := range dataset.Data[0] {
			var currentCov float64 = 0
			for index := range dataset.Data {
				currentCov += (dataset.Data[index][key1] - mean[key1]) * (dataset.Data[index][key2] - mean[key2])
			}
			currentCorr := currentCov / (math.Sqrt(variance[key1]*variance[key2]) * float64(len(dataset.Data)))
			corr[key1][key2] = currentCorr
		}
	}
	return corr
}
