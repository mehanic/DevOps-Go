package main

import "fmt"

func main() {
	//цикл -> итерация(шаг,step)
	//i++  -> i = i + 1 (i-> счетчик)
	//i<5 -> условие работы цикла
	//for счетчик; условие; шаг { действия}
	for i := 1; i < 5; i++ {
		//действие1
		fmt.Println("start step")
		//действие2
		fmt.Println("gooo")
		//действие3
		fmt.Println("end step")
		fmt.Println("")
	}
}

// start step
// gooo
// end step

// start step
// gooo
// end step

// start step
// gooo
// end step

// start step
// gooo
// end step
