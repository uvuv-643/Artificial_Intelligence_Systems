package data

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"strings"
)

func BoxPlot(dataset Dataset, output string) {
	dataset = copyDataset(dataset)
	index := 0
	for column := range dataset.Data[0] {
		p := plot.New()
		p.Title.Text = "BOX PLOT " + strings.ReplaceAll(strings.ToUpper(column), "_", " ")
		var values plotter.Values
		for _, row := range dataset.Data {
			values = append(values, row[column])
		}
		box, err := plotter.NewBoxPlot(vg.Points(30), float64(index), values)
		if err != nil {
			panic(err)
		}
		p.Add(box)
		if err := p.Save(5*vg.Inch, 5*vg.Inch, output+"/"+column+".png"); err != nil {
			panic(err)
		}
		index++
	}

}

func ScatterPlot(dataset Dataset, target string, output string) {
	fields := dataset.Fields()
	for _, field := range fields {
		p := plot.New()
		p.Title.Text = fmt.Sprintf("Scatter Plot: %s vs. %s", target, field)
		p.X.Label.Text = field
		p.Y.Label.Text = target

		pts := make(plotter.XYs, len(dataset.Data))
		for i, x := range dataset.Data {
			pts[i].X = x[field]
			pts[i].Y = x[target]
		}

		s, err := plotter.NewScatter(pts)
		if err != nil {
			panic(err)
		}
		s.GlyphStyle.Shape = draw.CircleGlyph{}
		s.GlyphStyle.Color = plotutil.Color(0)
		s.GlyphStyle.Radius = vg.Length(1)
		p.Add(s)

		if err := p.Save(8*vg.Inch, 8*vg.Inch, output+"/"+field+".png"); err != nil {
			panic(err)
		}
	}

}
