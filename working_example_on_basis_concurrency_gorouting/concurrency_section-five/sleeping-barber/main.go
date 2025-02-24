package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// sleeping-barber problem

var seatingCapacity = 2
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			// if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("there is nothing to do, so %s takes a nap", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed - send the barber home & close thi goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s' hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("closing shop for the day.")

	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		// will block and wait until all barbers done
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	color.Green("--------------------------------------------------------------------")
	color.Green("the barbershop is now closed for the day and everyone has gone home.")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("*** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room", client)
		default:
			color.Red("the waiting room is full, so $s leaves!", client)
		}
	} else {
		color.Red("the shop is already closed, so %s leaves!", client)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Yellow("The sleeping barber problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}

	color.Green("the shop is open for the day!")

	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Gerard")
	shop.addBarber("Milton")
	shop.addBarber("Susan")
	shop.addBarber("Kelly")
	shop.addBarber("Pat")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("client %d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}

// ╰─λ go run main.go                                                                        1 (0.532s) < 12:28:56
// The sleeping barber problem
// ---------------------------
// the shop is open for the day!
// Frank goes to the waiting room to check for clients
// there is nothing to do, so Frank takes a nap
// Pat goes to the waiting room to check for clients
// there is nothing to do, so Pat takes a nap
// Gerard goes to the waiting room to check for clients
// Milton goes to the waiting room to check for clients
// there is nothing to do, so Gerard takes a nap
// Susan goes to the waiting room to check for clients
// there is nothing to do, so Susan takes a nap
// there is nothing to do, so Milton takes a nap
// Kelly goes to the waiting room to check for clients
// there is nothing to do, so Kelly takes a nap
// *** client 1 arrives!
// client 1 takes a seat in the waiting room
// client 1 wakes Frank up
// Frank is cutting client 1' hair.
// *** client 2 arrives!
// client 2 takes a seat in the waiting room
// client 2 wakes Pat up
// Pat is cutting client 2' hair.
// *** client 3 arrives!
// client 3 takes a seat in the waiting room
// client 3 wakes Gerard up
// Gerard is cutting client 3' hair.
// *** client 4 arrives!
// client 4 takes a seat in the waiting room
// client 4 wakes Susan up
// Susan is cutting client 4' hair.
// *** client 5 arrives!
// client 5 takes a seat in the waiting room
// client 5 wakes Milton up
// Milton is cutting client 5' hair.
// *** client 6 arrives!
// client 6 takes a seat in the waiting room
// client 6 wakes Kelly up
// Kelly is cutting client 6' hair.
// *** client 7 arrives!
// client 7 takes a seat in the waiting room
// *** client 8 arrives!
// client 8 takes a seat in the waiting room
// *** client 9 arrives!
// the waiting room is full, so $s leaves!
// %!(EXTRA string=client 9)Frank is finished cutting client 1's hair.
