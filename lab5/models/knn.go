package models

import (
	"lab5/data"
	"math"
	"math/rand"
	"sort"
)

type KNN struct {
	dataset data.Dataset
	target  string
	fields  []string
	k       int
}

type KNNDist struct {
	dist  float64
	class float64
}

func (knn *KNN) Fit(dataset data.Dataset, target string, k int) {
	knn.target = target
	knn.dataset = dataset
	knn.k = k
	for key := range dataset.Data[0] {
		if key != target {
			knn.fields = append(knn.fields, key)
		}
	}

}

func (knn *KNN) Predict(dataset data.Dataset) []float64 {
	var answers = make([]float64, 0)
	for _, current := range dataset.Data {
		var distances = make([]KNNDist, 0)
		for _, neighbour := range knn.dataset.Data {
			currDist := 0.0
			for _, key := range knn.fields {
				currDist += math.Pow(current[key]-neighbour[key], 2)
			}
			currDist = math.Sqrt(currDist)
			distances = append(distances, KNNDist{
				currDist,
				neighbour[knn.target],
			})
		}
		sort.Slice(distances, func(i, j int) bool {
			return distances[i].dist < distances[j].dist
		})

		var classDistribution = make(map[float64]float64)
		maxCnt := 0.0
		for _, value := range distances[:knn.k] {
			classDistribution[value.class] += 1
			maxCnt = math.Max(maxCnt, classDistribution[value.class])
		}
		var maxCntClasses = make([]float64, 0)
		for key, value := range classDistribution {
			if value == maxCnt {
				maxCntClasses = append(maxCntClasses, key)
			}
		}
		answers = append(answers, maxCntClasses[rand.Intn(len(maxCntClasses))])
	}
	return answers
}
