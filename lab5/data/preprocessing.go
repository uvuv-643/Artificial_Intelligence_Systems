package data

import (
	"errors"
	"golang.org/x/exp/rand"
	"lab5/utils"
	"math"
)

func copyDataset(dataset Dataset) Dataset {
	newDataset := Dataset{
		Data: make([]map[string]float64, len(dataset.Data)),
	}
	for index, row := range dataset.Data {
		newDataset.Data[index] = make(map[string]float64)
		for key, value := range row {
			newDataset.Data[index][key] = value
		}
	}
	return newDataset
}

func Drop(dataset Dataset, columns ...string) Dataset {
	dataset = copyDataset(dataset)
	for _, key := range columns {
		for _, element := range dataset.Data {
			delete(element, key)
		}
	}
	return dataset
}

func Get(dataset Dataset, column string) []float64 {
	answer := make([]float64, 0)
	for _, element := range dataset.Data {
		answer = append(answer, element[column])
	}
	return answer
}

func Filter(dataset Dataset, callback func(map[string]float64) bool) Dataset {
	newDataset := Dataset{
		Data: make([]map[string]float64, 0),
	}
	for _, row := range dataset.Data {
		if callback(row) {
			newDataset.Data = append(newDataset.Data, row)
		}
	}
	return newDataset
}

func Map(dataset Dataset, callback func(map[string]float64) map[string]float64) Dataset {
	newDataset := Dataset{
		Data: make([]map[string]float64, 0),
	}
	for _, row := range dataset.Data {
		newRow := make(map[string]float64)
		for key, value := range row {
			newRow[key] = value
		}
		newDataset.Data = append(newDataset.Data, callback(newRow))
	}
	return newDataset
}

func TrainTestSplit(dataset Dataset, testSize float64, randomState uint64) (Dataset, Dataset, error) {

	trainDataset := Dataset{
		Data: make([]map[string]float64, 0),
	}
	testDataset := Dataset{
		Data: make([]map[string]float64, 0),
	}

	if testSize <= 0 || testSize >= 1 {
		return dataset, dataset, errors.New("wrong test size param")
	}

	// делим изначальный датасет на два
	rand.Seed(randomState)
	for _, row := range dataset.Data {
		if rand.Float64() <= testSize {
			testDataset.Data = append(testDataset.Data, row)
		} else {
			trainDataset.Data = append(trainDataset.Data, row)
		}
	}

	// выравниваем размер полученных датасетов до testSize * dataset <= actualTestSize
	haveToBeRemovedFromTestCount := len(testDataset.Data) - int(testSize*float64(len(dataset.Data)))
	if haveToBeRemovedFromTestCount > 0 {
		leftElements := testDataset.Data[len(testDataset.Data)-haveToBeRemovedFromTestCount:]
		testDataset.Data = testDataset.Data[0 : len(testDataset.Data)-haveToBeRemovedFromTestCount]
		trainDataset.Data = append(trainDataset.Data, leftElements...)
	}

	haveToBeRemovedFromTrainCount := len(trainDataset.Data) - int((1-testSize)*float64(len(dataset.Data)))
	if haveToBeRemovedFromTrainCount > 0 {
		leftElements := trainDataset.Data[len(trainDataset.Data)-haveToBeRemovedFromTrainCount:]
		trainDataset.Data = trainDataset.Data[0 : len(trainDataset.Data)-haveToBeRemovedFromTrainCount]
		testDataset.Data = append(testDataset.Data, leftElements...)
	}

	return trainDataset, testDataset, nil
}

func CreateDistanceField(dataset Dataset, reachestHouses Dataset) Dataset {
	return Map(dataset, func(row map[string]float64) map[string]float64 {
		targetDistance := math.MaxFloat64
		for _, house := range reachestHouses.Data {
			currentDist := utils.Dist(house["longitude"], house["latitude"], row["longitude"], row["latitude"])
			targetDistance = math.Min(targetDistance, currentDist)
		}
		row["distance_to_reach_people"] = 0
		if targetDistance > 0.01 {
			row["distance_to_reach_people"] = math.Log10(targetDistance + 1)
		}
		return row
	})
}
