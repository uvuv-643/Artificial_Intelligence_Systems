package models

import (
	"gonum.org/v1/gonum/mat"
	"lab4/data"
	"math"
)

type LinearRegression struct {
	Title  string
	result []float64
	fields []string
	target string
	Coef   map[string]float64
}

//func (lr *LinearRegression) Fit(dataset data.Dataset, target string) {
//
//	lr.target = target
//	lr.fields = make([]string, 0)
//	lr.Coef = make(map[string]float64)
//
//	for key := range dataset.Data[0] {
//		if key != target {
//			lr.fields = append(lr.fields, key)
//		}
//	}
//
//	n := len(dataset.Data)
//	m := len(lr.fields)
//
//	lr.Coef_ = make([]float64, m+1)
//
//	y := make([]float64, 0)
//	for _, row := range dataset.Data {
//		y = append(y, row[target])
//	}
//
//	x := make([][]float64, 0)
//	for index := range dataset.Data {
//		x = append(x, make([]float64, 0))
//		for _, field := range lr.fields {
//			x[index] = append(x[index], dataset.Data[index][field])
//		}
//	}
//
//	fmt.Println(x[0:10])
//
//	xT := make([][]float64, m+1)
//	for i := range xT {
//		xT[i] = make([]float64, n)
//	}
//
//	for i := 0; i < n; i++ {
//		for j := 0; j < m; j++ {
//			xT[j][i] = x[i][j]
//		}
//		xT[m][i] = 1.0
//	}
//
//	xTxInv := make([][]float64, m+1)
//	for i := range xTxInv {
//		xTxInv[i] = make([]float64, m+1)
//	}
//
//	xTy := make([]float64, m+1)
//
//	mean := data.Mean(dataset)[target]
//	for i := 0; i <= m; i++ {
//		for j := 0; j <= m; j++ {
//			for k := 0; k < n; k++ {
//				xTxInv[i][j] += xT[i][k] * xT[j][k]
//			}
//		}
//		for k := 0; k < n; k++ {
//			xTy[i] += xT[i][k] * (y[k] - mean)
//		}
//	}
//
//	for i := 0; i <= m; i++ {
//		for j := 0; j <= m; j++ {
//			lr.Coef_[i] += xTxInv[i][j] * xTy[j]
//		}
//	}
//
//	fmt.Println(lr.Coef_)
//
//}

func (lr *LinearRegression) Fit(dataset data.Dataset, target string) {
	lr.target = target
	lr.fields = make([]string, 0)
	lr.Coef = make(map[string]float64)

	for key := range dataset.Data[0] {
		if key != target {
			lr.fields = append(lr.fields, key)
		}
	}

	xMatrixData := make([]float64, 0)
	for index := range dataset.Data {
		for _, field := range lr.fields {
			xMatrixData = append(xMatrixData, dataset.Data[index][field])
		}
	}

	xTransMatrixData := make([]float64, len(dataset.Data)*len(lr.fields))
	for i := range dataset.Data {
		for j := range lr.fields {
			if i < j {
				xTransMatrixData[i*len(lr.fields)+j] = xTransMatrixData[j*len(dataset.Data)+i]
			}
		}
	}

	yMatrixData := make([]float64, 0)
	for _, row := range dataset.Data {
		yMatrixData = append(yMatrixData, row[target])
	}

	xMatrix := mat.NewDense(len(dataset.Data), len(lr.fields), xMatrixData)
	//xTransMatrix := mat.NewDense(len(lr.fields), len(dataset.Data), xMatrixData)
	yMatrix := mat.NewDense(len(dataset.Data), 1, yMatrixData)

	qr := new(mat.QR)
	qr.Factorize(xMatrix)
	q := new(mat.Dense)
	r := new(mat.Dense)
	qr.QTo(q)
	qr.RTo(r)

	var xTxInvxTy mat.Dense

	xTxInvxTy.Mul(q.T(), yMatrix)

	c := make([]float64, len(lr.fields))
	for i := len(lr.fields) - 1; i >= 0; i-- {
		c[i] = xTxInvxTy.At(i, 0)
		for j := i + 1; j < len(lr.fields); j++ {
			c[i] -= c[j] * r.At(i, j)
		}
		c[i] /= r.At(i, i)
	}

	for i, key := range lr.fields {
		lr.Coef[key] = c[i]
	}

}

func (lr *LinearRegression) Predict(dataset data.Dataset) []float64 {
	answers := make([]float64, len(dataset.Data))
	for index, row := range dataset.Data {
		answer := 0.0
		for _, field := range lr.fields {
			answer += lr.Coef[field] * row[field]
		}
		answer = math.Max(0, math.Min(500000.0, answer))
		answers[index] = answer
	}
	return answers
}
