package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// generateData creates synthetic linear data with noise
func generateData(n int) ([]float64, []float64) {
	rand.Seed(time.Now().UnixNano())
	X := make([]float64, n)
	Y := make([]float64, n)

	trueSlope := 5.0
	trueIntercept := 1.0

	for i := 0; i < n; i++ {
		x := rand.Float64()*6 - 3 // Range: -3 to 3
		noise := rand.NormFloat64() * 7
		y := trueSlope*x + trueIntercept + noise
		X[i] = x
		Y[i] = y
	}
	return X, Y
}

// gradientDescent performs linear regression using gradient descent
func gradientDescent(X, Y []float64, theta float64, alpha float64, iterations int) float64 {
	m := float64(len(X))

	for i := 0; i < iterations; i++ {
		var lossSum float64
		var gradientSum float64

		for j := 0; j < len(X); j++ {
			h := theta*X[j] + 1 // intercept is fixed at 1
			loss := h - Y[j]
			lossSum += loss * loss
			gradientSum += loss * X[j]
		}

		cost := lossSum / (2 * m)
		fmt.Printf("After Iteration %d | Cost: %f\n", i, cost)

		gradient := gradientSum / m
		theta = theta - alpha*gradient
	}

	return theta
}

// createPlot plots the points and fitted line
func createPlot(X, Y []float64, theta float64) {
	p := plot.New()
	p.Title.Text = "Linear Regression with Gradient Descent"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Scatter data
	pts := make(plotter.XYs, len(X))
	for i := range X {
		pts[i].X = X[i]
		pts[i].Y = Y[i]
	}

	scatter, _ := plotter.NewScatter(pts)
	p.Add(scatter)

	// Fitted line
	line := plotter.NewFunction(func(x float64) float64 {
		return theta*x + 1
	})
	line.Width = vg.Points(2)
	p.Add(line)

	p.Save(6*vg.Inch, 4*vg.Inch, "regression.png")
}

func main() {
	X, Y := generateData(60)
	initialTheta := rand.Float64()
	alpha := 0.4
	iterations := 30

	finalTheta := gradientDescent(X, Y, initialTheta, alpha, iterations)
	createPlot(X, Y, finalTheta)

	fmt.Printf("Final slope (theta): %f\n", finalTheta)
	fmt.Println("Plot saved to regression.png")
}
