package utils

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"math"
	"os"
	"sort"
	"strings"
)

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func PrintRow(row map[string]float64) {
	keys := make([]string, 0)
	for key := range row {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(keys)
	var values []string
	for _, key := range keys {
		values = append(values, fmt.Sprintf("%.2f", row[key]))
	}
	table.Append(values)
	table.Render()
}

func PrintlnNewLine(data ...interface{}) {
	fmt.Println()
	fmt.Println()
	fmt.Println(data...)
}

func PrintMatrixByMaps(matrix map[string]map[string]float64) {

	keys := make([]string, 0)
	keys = append(keys, "")
	for key := range matrix {
		keys = append(keys, key)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(keys)

	for _, key1 := range keys {
		var values []string
		for _, key2 := range keys {
			if key2 == "" && key1 != "" {
				values = append(values, strings.ReplaceAll(strings.ToUpper(key1), "_", " "))
			} else if key1 != "" {
				values = append(values, fmt.Sprintf("%.4f", matrix[key1][key2]))
			}
		}
		table.Append(values)
	}

	table.Render()

}

func Dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2.0) + math.Pow(y2-y1, 2.0))
}

func PrintResultComparing(initial, predicted []float64) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Initial value", "Predicted value"})
	printedSize := int(math.Min(float64(len(initial)), 10.0))
	for i := range initial[:printedSize] {
		var values []string
		values = append(values, fmt.Sprintf("%.4f", initial[i]))
		values = append(values, fmt.Sprintf("%.4f", predicted[i]))
		table.Append(values)
	}
	table.Render()
}
