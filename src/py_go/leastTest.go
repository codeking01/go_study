package main

/**
 * @author xiongjl
 * @since 2023/11/30  12:09
 */

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math/rand"
	"path/filepath"
	"runtime"
)

func main() {
	var a, b float64 = 0.7, 3
	points1 := plotter.XYs{}
	points2 := plotter.XYs{}

	for i := 0; i <= 10; i++ {
		points1 = append(points1, plotter.XY{
			X: float64(i),
			Y: a*float64(i) + b,
		})
		points2 = append(points2, plotter.XY{
			X: float64(i),
			Y: a*float64(i) + b + (2*rand.Float64() - 1),
		})
	}

	fa, fb := LeastSquares(points2)
	points3 := plotter.XYs{}
	for i := 0; i <= 10; i++ {
		points3 = append(points3, plotter.XY{
			X: float64(i),
			Y: fa*float64(i) + fb,
		})
	}

	plt := plot.New()
	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 10, 10

	if err := plotutil.AddLinePoints(plt,
		"line1", points1,
		"line2", points2,
		"line3", points3,
	); err != nil {
		panic(err)
	}
	// 获取源码的路径（当前源码所在的文件夹）
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get the caller's information")
	}
	// 获取源码目录
	currentSourceDir := filepath.Dir(filename)
	fileName := filepath.Join(currentSourceDir, "03-least-squares.png")
	fmt.Println(fileName)
	if err := plt.Save(5*vg.Inch, 5*vg.Inch, fileName); err != nil {
		panic(err)
	}
}

func LeastSquares(points plotter.XYs) (a, b float64) {
	var xSum, ySum float64
	for _, point := range points {
		xSum += point.X
		ySum += point.Y
	}
	xAvg, yAvg := xSum/float64(points.Len()), ySum/float64(points.Len())

	var xySum, xxSum float64
	for _, point := range points {
		xySum += (point.X - xAvg) * (point.Y - yAvg)
		xxSum += (point.X - xAvg) * (point.X - xAvg)
	}

	a = xySum / xxSum
	b = yAvg - a*xAvg
	return
}
