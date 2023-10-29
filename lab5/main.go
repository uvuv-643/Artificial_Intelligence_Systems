package main

import (
	"fmt"
	"golang.org/x/exp/rand"
	"lab5/config"
	"lab5/data"
	"lab5/models"
	"lab5/utils"
	"strconv"
)

func main() {

	targetField := "Wine"
	configuration := config.GetConfig()
	dataset := data.LoadData(configuration.DatasetName)

	utils.PrintlnNewLine("FIRST 10 ELEMENTS OF DATASET")
	dataset.Print()

	utils.PrintlnNewLine("MIN VALUES")
	utils.PrintRow(data.Min(dataset))

	utils.PrintlnNewLine("MAX VALUES")
	utils.PrintRow(data.Max(dataset))

	utils.PrintlnNewLine("MEAN VALUES")
	utils.PrintRow(data.Mean(dataset))

	utils.PrintlnNewLine("STANDARD DEVIATION")
	utils.PrintRow(data.StandardDeviation(dataset))

	utils.PrintlnNewLine("MEDIAN")
	median, err := data.Quantile(dataset, 0.5)
	if err != nil {
		fmt.Println("There is no median for this dataset, sorry!")
	}
	utils.PrintRow(median)

	utils.PrintlnNewLine("QUANTILE 0.25")
	quantile1, err := data.Quantile(dataset, 0.25)
	if err != nil {
		fmt.Println("There is no quantile .25 for this dataset, sorry!")
	}
	utils.PrintRow(quantile1)

	utils.PrintlnNewLine("QUANTILE 0.75")
	quantile2, err := data.Quantile(dataset, 0.75)
	if err != nil {
		fmt.Println("There is no quantile .75 for this dataset, sorry!")
	}
	utils.PrintRow(quantile2)

	testSize, _ := strconv.ParseFloat(configuration.SplitTestSize, 64)
	randomSplitState, _ := strconv.ParseUint(configuration.SplitRandomState, 10, 64)
	train, test, err := data.TrainTestSplit(dataset, testSize, randomSplitState)
	if err != nil {
		panic(err)
	}
	utils.PrintlnNewLine("TRAIN")
	train.Print()
	utils.PrintlnNewLine("TEST")
	test.Print()

	initialTestData := data.Get(test, targetField)
	test = data.Drop(test, targetField)

	k, _ := strconv.ParseInt(configuration.KnnK, 10, 32)
	func() {

		train1 := data.Drop(train)
		test1 := data.Drop(test)

		ln := data.LinearNorm{}
		ln.Fit(train1, targetField)
		train1 = ln.Apply(train1)
		test1 = ln.Apply(test1)

		knn := models.KNN{}
		knn.Fit(train1, targetField, int(k))

		utils.PrintlnNewLine("First model - all columns")
		train1.Print()

		predicted := knn.Predict(test1)

		fmt.Println()
		fmt.Println("Initial \n", initialTestData)
		fmt.Println("Predicted \n", predicted)
		fmt.Println()
		fmt.Println("Confusion Matrix (predicted - rows, expected - cols)")
		utils.PrintConfusionMatrix(initialTestData, predicted)

	}()

	func() {

		fields := make([]string, 0)
		for key := range dataset.Data[0] {
			if key != targetField {
				if rand.Float32() < 0.5 {
					fields = append(fields, key)
				}
			}
		}

		train1 := data.Drop(train, fields...)
		test1 := data.Drop(test, fields...)

		ln := data.LinearNorm{}
		ln.Fit(train1, targetField)
		train1 = ln.Apply(train1)
		test1 = ln.Apply(test1)

		knn := models.KNN{}
		knn.Fit(train1, targetField, int(k))

		utils.PrintlnNewLine("Second model - random columns")
		train1.Print()

		predicted := knn.Predict(test1)

		fmt.Println()
		fmt.Println("Initial \n", initialTestData)
		fmt.Println("Predicted \n", predicted)
		fmt.Println()
		fmt.Println("Confusion Matrix (predicted - rows, expected - cols)")
		utils.PrintConfusionMatrix(initialTestData, predicted)

	}()

}
