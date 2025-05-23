package main

import (
	"fmt"
	"sync"
	"time"
)

// dining philosophers

// Philosopher is a struct which stores information about a philosopher.
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers list of all philosophers.
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define some vars
var hunger = 3 // how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func main() {
	// print out a welcome message
	fmt.Println("dining philosophers problem")
	fmt.Println("---------------------------")
	fmt.Println("the table is empty")

	time.Sleep(sleepTime)

	// start the meal
	dine()

	// print out finished message
	fmt.Println("the table is empty")
}

func dine() {
	//eatTime = 0 * time.Second
	//sleepTime = 0 * time.Second
	//thinkTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal.
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the ph at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)

	seated.Done()

	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table.")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}

// ╰─λ go run main.go                                                                        0 (0.001s) < 12:53:53
// dining philosophers problem
// ---------------------------
// the table is empty
// Socrates is seated at the table.
// Locke is seated at the table.
// Aristotle is seated at the table.
// Plato is seated at the table.
// Pascal is seated at the table.
// 	Pascal takes the left fork.
// 	Pascal takes the right fork.
// 	Pascal has both forks and is eating.
// 	Plato takes the right fork.
// 	Plato takes the left fork.
// 	Plato has both forks and is eating.
// 	Aristotle takes the left fork.
// 	Plato is thinking.
// 	Pascal is thinking.
// 	Pascal put down the forks.
// 	Pascal takes the left fork.
// 	Pascal takes the right fork.
// 	Pascal has both forks and is eating.
// 	Plato put down the forks.
// 	Plato takes the right fork.
// 	Plato takes the left fork.
// 	Plato has both forks and is eating.
// 	Plato is thinking.
// 	Pascal is thinking.
// 	Socrates takes the left fork.
// 	Plato put down the forks.
// 	Locke takes the left fork.
// 	Locke takes the right fork.
// 	Locke has both forks and is eating.
// 	Aristotle takes the right fork.
// 	Aristotle has both forks and is eating.
// 	Pascal put down the forks.
// 	Aristotle is thinking.
// 	Locke is thinking.
// 	Locke put down the forks.
// 	Locke takes the left fork.
// 	Locke takes the right fork.
// 	Locke has both forks and is eating.
// 	Aristotle put down the forks.
// 	Pascal takes the left fork.
// 	Socrates takes the right fork.
// 	Socrates has both forks and is eating.
// 	Socrates is thinking.
// 	Locke is thinking.
// 	Locke put down the forks.
// 	Locke takes the left fork.
// 	Locke takes the right fork.
// 	Socrates put down the forks.
// 	Plato takes the right fork.
// 	Locke has both forks and is eating.
// 	Aristotle takes the left fork.
// 	Locke is thinking.
// 	Pascal takes the right fork.
// 	Pascal has both forks and is eating.
// 	Locke put down the forks.
// Locke is satisfied.
// Locke left the table.
// 	Plato takes the left fork.
// 	Plato has both forks and is eating.
// 	Plato is thinking.
// 	Pascal is thinking.
// 	Pascal put down the forks.
// Pascal is satisfied.
// Pascal left the table.
// 	Aristotle takes the right fork.
// 	Aristotle has both forks and is eating.
// 	Plato put down the forks.
// Plato is satisfied.
// Plato left the table.
// 	Socrates takes the left fork.
// 	Aristotle is thinking.
// 	Aristotle put down the forks.
// 	Socrates takes the right fork.
// 	Socrates has both forks and is eating.
// 	Socrates is thinking.
// 	Socrates put down the forks.
// 	Socrates takes the left fork.
// 	Socrates takes the right fork.
// 	Socrates has both forks and is eating.
// 	Socrates is thinking.
// 	Aristotle takes the left fork.
// 	Aristotle takes the right fork.
// 	Aristotle has both forks and is eating.
// 	Socrates put down the forks.
// Socrates is satisfied.
// Socrates left the table.
// 	Aristotle is thinking.
// 	Aristotle put down the forks.
// Aristotle is satisfied.
// Aristotle left the table.
// the table is empty
