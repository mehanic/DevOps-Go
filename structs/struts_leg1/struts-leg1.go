package main

import (
	"fmt"
)

type Leg struct {
	rate     float64
	time     float64
	distance float64
	changes  []string
}

func (l *Leg) updateChanges(change string) {
	// Update the changes slice with the most recent changes
	if len(l.changes) >= 2 {
		l.changes = l.changes[1:]
	}
	l.changes = append(l.changes, change)
}

func (l *Leg) calculate() {
	compute := map[string]bool{"rate": true, "time": true, "distance": true}
	for _, ch := range l.changes {
		delete(compute, ch)
	}
	if len(compute) == 1 {
		change := ""
		for ch := range compute {
			change = ch
		}
		switch change {
		case "distance":
			l.distance = l.time * l.rate
		case "time":
			l.time = l.distance / l.rate
		case "rate":
			l.rate = l.distance / l.time
		}
	}
}

func (l *Leg) SetRate(value float64) {
	l.rate = value
	l.updateChanges("rate")
	l.calculate()
}

func (l *Leg) SetTime(value float64) {
	l.time = value
	l.updateChanges("time")
	l.calculate()
}

func (l *Leg) SetDistance(value float64) {
	l.distance = value
	l.updateChanges("distance")
	l.calculate()
}

// Leg_Alt struct inherits from Leg but uses methods for calculation
type Leg_Alt struct {
	Leg
}

func (l *Leg_Alt) calcDistance() {
	l.distance = l.time * l.rate
}

func (l *Leg_Alt) calcTime() {
	l.time = l.distance / l.rate
}

func (l *Leg_Alt) calcRate() {
	l.rate = l.distance / l.time
}

func (l *Leg_Alt) calculate(change string) {
	l.updateChanges(change)
	compute := map[string]bool{"rate": true, "time": true, "distance": true}
	for _, ch := range l.changes {
		delete(compute, ch)
	}
	if len(compute) == 1 {
		name := ""
		for ch := range compute {
			name = ch
		}
		switch name {
		case "distance":
			l.calcDistance()
		case "time":
			l.calcTime()
		case "rate":
			l.calcRate()
		}
	}
}

func (l *Leg_Alt) SetRate(value float64) {
	l.rate = value
	l.calculate("rate")
}

func (l *Leg_Alt) SetTime(value float64) {
	l.time = value
	l.calculate("time")
}

func (l *Leg_Alt) SetDistance(value float64) {
	l.distance = value
	l.calculate("distance")
}

func main() {
	leg1 := &Leg{}
	leg1.SetRate(6.0)
	leg1.SetDistance(35.6)
	fmt.Printf("option 1 %.1fnm at %.2fkt = %.2fhr\n", leg1.distance, leg1.rate, leg1.time)
	
	leg1.SetDistance(38.2)
	fmt.Printf("option 2 %.1fnm at %.2fkt = %.2fhr\n", leg1.distance, leg1.rate, leg1.time)
	
	leg1.SetTime(7)
	fmt.Printf("option 3 %.1fnm at %.2fkt = %.2fhr\n", leg1.distance, leg1.rate, leg1.time)
	
	leg2 := &Leg_Alt{Leg{rate: 6, distance: 35.6}}
	fmt.Printf("%.1f %.1f %.1f\n", leg2.rate, leg2.time, leg2.distance)
	fmt.Printf("option 1 %.1fnm at %.2fkt = %.2fhr\n", leg2.distance, leg2.rate, leg2.time)
}

// option 1 35.6nm at 6.00kt = 5.93hr
// option 2 38.2nm at 6.00kt = 5.93hr
// option 3 38.2nm at 5.46kt = 7.00hr
// 6.0 0.0 35.6
// option 1 35.6nm at 6.00kt = 0.00hr
