package main

import (
	"fmt"
)

// Leg struct to calculate rate, time, and distance
type Leg struct {
	rate     float64
	time     float64
	distance float64
	changes  []string
}

func (l *Leg) updateChanges(change string) {
	// Update the changes slice with the most recent changes
	l.changes = append(l.changes, change)
	if len(l.changes) > 2 {
		l.changes = l.changes[1:]
	}
}

func (l *Leg) calculate() {
	compute := map[string]bool{"rate": true, "time": true, "distance": true}
	for _, change := range l.changes {
		delete(compute, change)
	}
	if len(compute) == 1 {
		// Corrected the range loop usage for map iteration
		for key := range compute {
			switch key {
			case "distance":
				l.distance = l.time * l.rate
			case "time":
				l.time = l.distance / l.rate
			case "rate":
				l.rate = l.distance / l.time
			}
		}
	}
}

func (l *Leg) setRate(value float64) {
	l.rate = value
	l.updateChanges("rate")
	l.calculate()
}

func (l *Leg) setTime(value float64) {
	l.time = value
	l.updateChanges("time")
	l.calculate()
}

func (l *Leg) setDistance(value float64) {
	l.distance = value
	l.updateChanges("distance")
	l.calculate()
}

// Leg_Alt struct with similar functionality but using methods to calculate
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
		// Corrected the range loop usage for map iteration
		for name := range compute {
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
}

func (l *Leg_Alt) setRate(value float64) {
	l.rate = value
	l.calculate("rate")
}

func (l *Leg_Alt) setTime(value float64) {
	l.time = value
	l.calculate("time")
}

func (l *Leg_Alt) setDistance(value float64) {
	l.distance = value
	l.calculate("distance")
}

func main() {
	// First Leg with basic struct (Leg)
	leg1 := Leg{}
	leg1.setRate(6.0) // Setting rate in knots
	leg1.setDistance(35.6) // Setting distance in nautical miles
	fmt.Printf("Option 1: %.1f nm at %.2f kt = %.2f hrs\n", leg1.distance, leg1.rate, leg1.time)

	// Changing distance and recalculating time
	leg1.setDistance(38.2)
	fmt.Printf("Option 2: %.1f nm at %.2f kt = %.2f hrs\n", leg1.distance, leg1.rate, leg1.time)

	// Setting time and calculating distance
	leg1.setTime(7)
	fmt.Printf("Option 3: %.1f nm at %.2f kt = %.2f hrs\n", leg1.distance, leg1.rate, leg1.time)

	// Second Leg with alternate struct (Leg_Alt)
	leg2 := Leg_Alt{Leg{rate: 6, distance: 35.6}}
	leg2.calculate("rate") // Ensuring calculation
	fmt.Printf("\nOption 4 (Leg_Alt): %.1f nm at %.2f kt = %.2f hrs\n", leg2.distance, leg2.rate, leg2.time)

	// Changing rate and recalculating time
	leg2.setRate(7)
	fmt.Printf("Option 5 (Leg_Alt): %.1f nm at %.2f kt = %.2f hrs\n", leg2.distance, leg2.rate, leg2.time)

	// Changing distance and recalculating time
	leg2.setDistance(42.0)
	fmt.Printf("Option 6 (Leg_Alt): %.1f nm at %.2f kt = %.2f hrs\n", leg2.distance, leg2.rate, leg2.time)

	// Changing time and recalculating rate
	leg2.setTime(8)
	fmt.Printf("Option 7 (Leg_Alt): %.1f nm at %.2f kt = %.2f hrs\n", leg2.distance, leg2.rate, leg2.time)
}


// Option 1: 35.6 nm at 6.00 kt = 5.93 hrs
// Option 2: 38.2 nm at 6.00 kt = 5.93 hrs
// Option 3: 38.2 nm at 5.46 kt = 7.00 hrs

// Option 4 (Leg_Alt): 35.6 nm at 6.00 kt = 0.00 hrs
// Option 5 (Leg_Alt): 35.6 nm at 7.00 kt = 0.00 hrs
// Option 6 (Leg_Alt): 42.0 nm at 7.00 kt = 6.00 hrs
// Option 7 (Leg_Alt): 42.0 nm at 5.25 kt = 8.00 hrs

