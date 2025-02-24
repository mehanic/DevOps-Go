package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("received order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** we ran out of ingrediens for pizza #%d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** the cook quit while making pizza #%d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza order #%d is ready", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0
	// run forever or until we receive a quit notification
	// try to make pizzas

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				// close simple-channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("---------------------------------")
	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("the customer is really mad")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** error closing channel", err)
			}
		}
	}

	// print out ending message
	color.Cyan("----------------")
	color.Cyan("done for the day")

	color.Cyan("we made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("it was an awful day")
	case pizzasFailed >= 6:
		color.Red("it was not very good day")
	case pizzasFailed >= 4:
		color.Yellow("it was an ok day")
	case pizzasFailed >= 2:
		color.Yellow("it was a pretty good day")
	default:
		color.Green("it was a great day")
	}
}

// ╰─λ go run main.go                                                                        0 (0.001s) < 12:55:57
// The Pizzeria is open for business
// ---------------------------------
// received order #1
// making pizza #1. It will take 2 seconds...
// received order #2
// making pizza #2. It will take 4 seconds...
// pizza order #1 is ready
// order #1 is out for delivery
// received order #3
// making pizza #3. It will take 1 seconds...
// pizza order #2 is ready
// order #2 is out for delivery
// received order #4
// making pizza #4. It will take 1 seconds...
// pizza order #3 is ready
// order #3 is out for delivery
// received order #5
// making pizza #5. It will take 5 seconds...
// pizza order #4 is ready
// order #4 is out for delivery
// received order #6
// making pizza #6. It will take 5 seconds...
// pizza order #5 is ready
// order #5 is out for delivery
// received order #7
// making pizza #7. It will take 1 seconds...
// *** we ran out of ingrediens for pizza #6
// the customer is really mad
// received order #8
// making pizza #8. It will take 1 seconds...
// *** the cook quit while making pizza #7
// the customer is really mad
// received order #9
// making pizza #9. It will take 2 seconds...
// pizza order #8 is ready
// order #8 is out for delivery
// received order #10
// making pizza #10. It will take 4 seconds...
// pizza order #9 is ready
// order #9 is out for delivery
// pizza order #10 is ready
// order #10 is out for delivery
// Done making pizzas...
// ----------------
// done for the day
// we made 8 pizzas, but failed to make 2, with 10 attempts in total
// it was a pretty good day
