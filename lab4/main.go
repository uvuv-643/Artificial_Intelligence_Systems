package main

import (
	"fmt"
	"lab4/config"
	"lab4/data"
	"lab4/models"
	"lab4/utils"
	"strconv"
)

func main() {

	targetField := "median_house_value"
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

	utils.PrintlnNewLine("CORRELATION MATRIX")
	utils.PrintMatrixByMaps(data.Corr(dataset))

	// visualization
	go func() {
		data.BoxPlot(dataset, "target/boxplot")
	}()

	// test_train_split
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

	func() {

		reachestHouses := data.Filter(train, func(row map[string]float64) bool {
			return row[targetField] == 500000.0
		})
		train1 := data.CreateDistanceField(train, reachestHouses)
		test1 := data.CreateDistanceField(test, reachestHouses)

		go func() {
			data.ScatterPlot(train1, targetField, "target/scatter")
		}()

		train1 = data.Drop(train1, "longitude", "latitude")
		test1 = data.Drop(test1, "longitude", "latitude")

		ln := data.LinearNorm{}
		ln.Fit(train1, targetField)
		train1 = ln.Apply(train1)
		test1 = ln.Apply(test1)

		lr := models.LinearRegression{}
		lr.Fit(train1, targetField)

		utils.PrintlnNewLine("First model - new feature + linear normalization + drop columns")
		train1.Print()
		predicted := lr.Predict(test1)
		fmt.Println()
		utils.PrintResultComparing(initialTestData, predicted)
		fmt.Println("R2 for first model", data.R2(initialTestData, predicted))

	}()

	func() {

		train2 := data.Drop(train, "longitude", "latitude")
		test2 := data.Drop(test, "longitude", "latitude")

		ln := data.LinearNorm{}
		ln.Fit(train2, targetField)
		train2 = ln.Apply(train2)
		test2 = ln.Apply(test2)

		lr := models.LinearRegression{}
		lr.Fit(train2, targetField)

		utils.PrintlnNewLine("Second model - linear normalization + drop columns")
		train2.Print()
		predicted := lr.Predict(test2)
		fmt.Println()
		utils.PrintResultComparing(initialTestData, predicted)
		fmt.Println("R2 for second model", data.R2(initialTestData, predicted))

	}()

	func() {

		lr := models.LinearRegression{}
		lr.Fit(train, targetField)

		utils.PrintlnNewLine("Third model - default params")
		train.Print()
		predicted := lr.Predict(test)
		fmt.Println()
		utils.PrintResultComparing(initialTestData, predicted)
		fmt.Println("R2 for third model", data.R2(initialTestData, predicted))

	}()

	//<-time.After(time.Second * 15)

	//firstModel := <-bufferedChan
	//secondModel := <-bufferedChan
	//thirdModel := <-bufferedChan

}
