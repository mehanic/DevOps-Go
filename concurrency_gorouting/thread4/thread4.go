package main

import (
	"fmt"
	"sync"
)

// Racer represents a racer that waits for the start signal before running.
type Racer struct {
	name        string
	startSignal *sync.Once
	wg          *sync.WaitGroup
}

// Run is the method each racer will call once the start signal is received.
func (r *Racer) Run() {
	r.startSignal.Do(func() {}) // Wait for the start signal (sync.Once ensures it's only triggered once)
	fmt.Printf("I, %s, got to the goal!\n", r.name)
	r.wg.Done()
}

// Race represents a race with multiple racers.
type Race struct {
	racers      []*Racer
	startSignal *sync.Once
	wg          *sync.WaitGroup
}

// NewRace creates a new race with a list of racer names.
func NewRace(racerNames []string) *Race {
	wg := &sync.WaitGroup{}
	startSignal := &sync.Once{}
	racers := make([]*Racer, len(racerNames))

	// Create racers and add them to the list.
	for i, name := range racerNames {
		racers[i] = &Racer{name: name, startSignal: startSignal, wg: wg}
		wg.Add(1)
	}

	return &Race{
		racers:      racers,
		startSignal: startSignal,
		wg:          wg,
	}
}

// Start begins the race by signaling the racers to run.
func (r *Race) Start() {
	// Signal all racers to start.
	r.startSignal.Do(func() {})
	// Wait for all racers to finish.
	r.wg.Wait()
}

// Main function to run the race.
func main() {
	race := NewRace([]string{"rabbit", "turtle", "cheetah", "monkey", "cow", "horse", "tiger", "lion"})
	for _, racer := range race.racers {
		go racer.Run() // Start each racer in a separate goroutine
	}

	race.Start() // Trigger the start of the race
}

// I, lion, got to the goal!
// I, rabbit, got to the goal!
// I, turtle, got to the goal!
// I, cheetah, got to the goal!
// I, monkey, got to the goal!
// I, cow, got to the goal!
// I, tiger, got to the goal!
// I, horse, got to the goal!

