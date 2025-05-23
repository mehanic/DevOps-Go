package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err    error
		height float64
		weight float64
	)
	fmt.Println("Tính chỉ số BMI của bạn dựa trên chiều cao và cân nặng bạn nhập vào!")
	reader := bufio.NewReader(os.Stdin)

	//Lấy chiều cao từ giá trị nhập vào
	fmt.Print("Chiều cao của bạn: ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	height, err = strconv.ParseFloat(str, 64)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Lấy cân nặng từ giá trị nhập vào
	fmt.Print("Cân nặng của bạn: ")
	str1, _ := reader.ReadString('\n')
	str1 = strings.TrimSuffix(str1, "\n")
	weight, err = strconv.ParseFloat(str1, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Tính chỉ số BMI, và in ra chỉ số trên màn hình
	bmi := CalculatorBMI(height, weight)
	fmt.Println("Chỉ số BMI của bạn = ", bmi)
	switch {
	case bmi < 18.5:
		fmt.Println("Chỉ số BMI của bạn theo WHO là: Dưới chuẩn")
	case bmi > 18.5 && bmi < 24.9:
		fmt.Println("Chỉ số BMI của bạn theo WHO là: Bình thường")
	case bmi == 25:
		fmt.Println("Chỉ số BMI của bạn theo WHO là: Thừa cân")
	case bmi > 25 && bmi < 30:
		fmt.Println("Chỉ số BMI của bạn theo WHO là: Béo phì")
	case bmi > 30 && bmi < 35:
		fmt.Println("Chỉ số BMI của bạn theo WHO là: Béo phì độ I")
	case bmi > 35 && bmi < 40:
		fmt.Println("Chỉ số BMI của bạn theo WHO là: Béo phì độ II")
	case bmi > 40:
		fmt.Println("Chỉ số BMI của bạn theo WHO là: Béo phì độ III")
	}
}

// Hàm được sử dụng trong hàm main ở tr
func CalculatorBMI(height float64, weight float64) (index float64) {
	return weight / (height * height)
}
