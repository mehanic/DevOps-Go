package main

import "fmt"

func main() {
	rates := [...]float64{
		25.5,  // ethereum
		120.5, // wanchain
	}

	// uses magic values - not good
	fmt.Printf("1 BTC is %g ETH\n", rates[0])
	fmt.Printf("1 BTC is %g WAN\n", rates[1])
	fmt.Println("----------")
	main1()
	fmt.Println("----------")
	main2()
}

func main1() {
	const (
		ETH = 9 - iota
		WAN
		ICX
		// you can add more cryptocurrencies here
		// watch out the -1 index though!
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
		ICX: 20,
		// you can add more cryptocurrencies here
	}

	// uses well-defined names (ETH, WAN, ...) - good
	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])
	fmt.Printf("1 BTC is %g ICX\n", rates[ICX])

	fmt.Printf("%#v\n", rates)
}

func main2() {
	type (
		// integer int

		bookcase [5]int
		cabinet  [5]int
		//          ^- try changing this to: integer
		//             but first: uncomment the integer type above
	)

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Print("Are they equal? ")

	if cabinet(blue) == red {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red : %#v\n", bookcase(red))
}

// 1 BTC is 25.5 ETH
// 1 BTC is 120.5 WAN
// ----------
// 1 BTC is 25.5 ETH
// 1 BTC is 120.5 WAN
// 1 BTC is 20 ICX
// [10]float64{0, 0, 0, 0, 0, 0, 0, 20, 120.5, 25.5}
// ----------
// Are they equal? ✅
// blue: main.bookcase{6, 9, 3, 2, 1}
// red : main.bookcase{6, 9, 3, 2, 1}
