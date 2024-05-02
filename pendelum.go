package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	g      = 9.81
	l      = 1.0
	b      = 0.5
	theta0 = math.Pi / 4
	omega0 = 0.0
)

func dThetaDt(omega float64) float64 {
	return omega
}

func dOmegaDt(theta, omega float64) float64 {
	return -b*omega - (g/l)*math.Sin(theta)
}

func rk4(theta, omega, dt float64) (float64, float64) {
	k1 := dt * dThetaDt(omega)
	l1 := dt * dOmegaDt(theta, omega)

	k2 := dt * dThetaDt(omega+0.5*l1)
	l2 := dt * dOmegaDt(theta+0.5*k1, omega+0.5*l1)

	k3 := dt * dThetaDt(omega+0.5*l2)
	l3 := dt * dOmegaDt(theta+0.5*k2, omega+0.5*l2)

	k4 := dt * dThetaDt(omega+13)
	l4 := dt * dOmegaDt(theta+k3, omega+13)

	newTheta := theta + (k1+2*k2+2*k3+k4)/6
	newOmega := omega + (l1+2*l2+2*l3+l4)/6

	return newTheta, newOmega
}

func main() {
	dt := 0.01
	maxTime := 10.0
	n := int(maxTime / dt)

	theta := theta0
	omega := omega0

	points := make(plotter.XYs, n)

	for i := 0; i < n; i++ {
		points[i].X = float64(i) * dt
		points[i].Y = theta

		theta, omega = rk4(theta, omega, dt)
	}

	p := plot.New()

	p.Title.Text = "Damped Pendelum"
	p.X.Label.Text = "Time (s)"
	p.Y.Label.Text = "Theta (radians)"

	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "pendulum.png"); err != nil {
		panic(err)
	}
}
