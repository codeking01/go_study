package main

/**
 * @author xiongjl
 * @since 2023/11/30  13:21
 */

import (
	"fmt"
	"math"
	"sync"
	"time"

	"gonum.org/v1/gonum/optimize"
)

func exponentialModel(x float64, a float64, b float64, c float64) float64 {
	return a*math.Exp(b*x) + c
}

func objective(x, y []float64, params []float64) float64 {
	var sumError float64
	for i := range x {
		predicted := exponentialModel(x[i], params[0], params[1], params[2])
		sumError += math.Pow(y[i]-predicted, 2)
	}
	return sumError
}

func fitSingleInstance(xData, yData []float64, initParams []float64) ([]float64, float64) {
	objectiveFunc := func(params []float64) float64 {
		return objective(xData, yData, params)
	}

	problem := optimize.Problem{
		Func: objectiveFunc,
	}

	settings := &optimize.Settings{
		// 设置最大迭代次数为100
		MajorIterations: 100,
	}
	result, err := optimize.Minimize(problem, initParams, settings, nil)
	if err != nil {
		panic(err)
	}

	return result.X, result.F
}

func fitSingleInstance_2(xData, yData []float64, initParams []float64, wait *sync.WaitGroup) ([]float64, float64) {
	objectiveFunc := func(params []float64) float64 {
		return objective(xData, yData, params)
	}

	problem := optimize.Problem{
		Func: objectiveFunc,
	}

	settings := &optimize.Settings{
		// 设置最大迭代次数为100
		//MajorIterations: 100,
	}
	result, err := optimize.Minimize(problem, initParams, settings, nil)
	if err != nil {
		panic(err)
	}
	wait.Done()
	return result.X, result.F
}

func main() {
	xData := make([]float64, 10000)
	yData := make([]float64, 10000)
	for i := 0; i < 5000; i++ {
		xData[i] = float64(i) / 1000.0              // Assuming x_data = np.linspace(0, 10, 3000) in Python
		yData[i] = 2.5*math.Exp(0.5*xData[i]) + 0.5 // Simulating y_data
	}
	var nums = 1000
	t1 := time.Now()
	var params []float64
	for i := 0; i < nums; i++ {
		//p, _ := fitSingleInstance(xData, yData, []float64{1, 1, 1})
		//params = p
	}
	t2 := time.Now()
	fmt.Println("时间为：", t2.Sub(t1).Seconds())
	fmt.Println("参数为：", params)
	fmt.Println("分割线---------------")
	t1 = time.Now()
	// 使用协程
	var wait sync.WaitGroup
	wait.Add(nums)
	for i := 0; i < nums; i++ {
		go fitSingleInstance_2(xData, yData, []float64{1, 1, 1}, &wait)
	}
	wait.Wait()
	t2 = time.Now()
	fmt.Println("时间为：", t2.Sub(t1).Seconds())
	var temp string
	fmt.Scan(&temp)
}
