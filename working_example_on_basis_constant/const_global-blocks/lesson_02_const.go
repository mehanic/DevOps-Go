package main

import (
	"fmt"
	"math"
)

const SingleLineString = "Single line constant"

const (
	MultiLineFirst = 1
	MultiLineSecond
	MultiLineThird = "third"
	MultiLineFourth
)

const (
	IotaUseCaseFirst = 1 << iota
	IotaUseCaseSecond
	IotaUseCaseThird
	IotaUseCaseFourth
)

const (
	IotaUseCase2First = iota*3 - 2
	IotaUseCase2Second
	IotaUseCase2Third
	IotaUseCase2Fourth
)

func main() {
	fmt.Println(SingleLineString, "\n")

	fmt.Println("MultiLineFirst ", MultiLineFirst)
	fmt.Println("MultiLineSecond ", MultiLineSecond)
	fmt.Println("MultiLineThird ", MultiLineThird)
	fmt.Println("MultiLineFourth ", MultiLineFourth, "\n")

	fmt.Println("IotaUseCaseFirst ", IotaUseCaseFirst)
	fmt.Println("IotaUseCaseSecond ", IotaUseCaseSecond)
	fmt.Println("IotaUseCaseThird ", IotaUseCaseThird)
	fmt.Println("IotaUseCaseFourth ", IotaUseCaseFourth, "\n")

	fmt.Println("IotaUseCase2First ", IotaUseCase2First)
	fmt.Println("IotaUseCase2Second ", IotaUseCase2Second)
	fmt.Println("IotaUseCase2Third ", IotaUseCase2Third)
	fmt.Println("IotaUseCase2Fourth ", IotaUseCase2Fourth, "\n")

	main1()
	fmt.Println("-------------------")
	main2()

}

func main1() {
	// const name = "foo"
	// name = "baa"
	// fmt.Println(name)

	const (
		sun = iota // 0
		mon        // 1
		tue        // 2
	)
	fmt.Println(sun, mon, tue)

}

const s string = "constant"

func main2() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

// Single line constant

// MultiLineFirst  1
// MultiLineSecond  1
// MultiLineThird  third
// MultiLineFourth  third

// IotaUseCaseFirst  1
// IotaUseCaseSecond  2
// IotaUseCaseThird  4
// IotaUseCaseFourth  8

// IotaUseCase2First  -2
// IotaUseCase2Second  1
// IotaUseCase2Third  4
// IotaUseCase2Fourth  7

// 0 1 2
// -------------------
// constant
// 6e+11
// 600000000000
// -0.28470407323754404
