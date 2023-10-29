package data

import (
	"encoding/csv"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"lab5/utils"
	"log"
	"os"
	"sort"
	"strconv"
)

type Dataset struct {
	Data []map[string]float64
}

const LimitOfElementsToPrint = 10
const (
	DatasetPrintAllRows = iota
)

func (dataset Dataset) Fields() []string {
	fields := make([]string, 0)
	for key := range dataset.Data[0] {
		fields = append(fields, key)
	}
	return fields
}

func (dataset Dataset) Print(params ...int) {

	fmt.Println("NUMBER OF ROWS: " + string(strconv.Itoa(len(dataset.Data))))

	arrayToPrint := dataset.Data
	if !utils.Contains(params, DatasetPrintAllRows) && len(arrayToPrint) > LimitOfElementsToPrint {
		arrayToPrint = arrayToPrint[:LimitOfElementsToPrint]
	}

	keys := make([]string, 0)
	for key, _ := range arrayToPrint[0] {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(keys)

	for _, row := range arrayToPrint {
		var values []string
		for _, key := range keys {
			values = append(values, fmt.Sprintf("%.5f", row[key]))
		}
		table.Append(values)
	}

	table.Render()

}

// Read CSV file and return it as matrix.
// First row indicates the column titles.
func readCsvFile(filePath string) [][]string {

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Unable to close input file "+filePath, err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records

}

func LoadData(datasetName string) Dataset {

	dataset := make([]map[string]float64, 0)
	records := readCsvFile("assets/" + datasetName)
	columnTitles := records[0]
	for i := 1; i < len(records); i++ {
		currentRow := make(map[string]float64)
		dataset = append(dataset, currentRow)
		for j, title := range columnTitles {
			float, err := strconv.ParseFloat(records[i][j], 32)
			if err != nil {
				log.Fatal("Incorrect dataset provided, cannot parse integer value", err)
			}
			currentRow[title] = float64(float)
		}
	}

	return Dataset{
		Data: dataset,
	}

}
