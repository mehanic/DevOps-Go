// Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.

// Example
// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
// The returned format must be correct in order to complete this challenge.

// Don't forget the space after the closing parentheses!

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func CreatePhoneNumberBase(n [10]uint) string { //uint8       the set of all unsigned  8-bit integers (0 to 255)

	// Convert slices to string and format as (xxx) xxx-xxxx
	f := fmt.Sprintf("%d%d%d", n[0], n[1], n[2])
	s := fmt.Sprintf("%d%d%d", n[3], n[4], n[5])
	t := fmt.Sprintf("%d%d%d%d", n[6], n[7], n[8], n[9])

	// Return the formatted phone number
	return fmt.Sprintf("(%s) %s-%s", f, s, t)

}

func main4() {
	n := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(CreatePhoneNumberBase(n))

}

//example1

func CreatePhoneNumber1(n [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
}

//

func ArrayToString(numbers interface{}) string {
	return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

func CreatePhoneNumber2(numbers [10]uint) string {
	str := ArrayToString(numbers)
	return fmt.Sprintf("(%s) %s-%s", str[0:3], str[3:6], str[6:10])
}

//example

func CreatePhoneNumber3(numbers [10]uint) string {
	tmp := make([]interface{}, len(numbers))
	for i, val := range numbers {
		tmp[i] = val
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", tmp...)
}

func CreatePhoneNumber4(numbers [10]uint) string {
	var test string = strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
	return fmt.Sprintf("(%s) %s-%s", test[0:3], test[3:6], test[6:10])
}

//

func str(numbers []uint) string {
	s := ""
	for _, n := range numbers {
		s += fmt.Sprintf("%v", n)
	}
	return s
}

func CreatePhoneNumber5(numbers [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", str(numbers[0:3]), str(numbers[3:6]), str(numbers[6:10]))
}

//

func CreatePhoneNumber6(numbers [10]uint) string {
	stringNumbers := make([]interface{}, len(numbers))
	for idx, num := range numbers {
		stringNumbers[idx] = num
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", stringNumbers...)
}

//

func CreatePhoneNumber7(numbers [10]uint) (result string) {
	for _, value := range numbers {
		result += fmt.Sprintf("%d", value)
	}

	return fmt.Sprintf("(%s) %s-%s", result[:3], result[3:6], result[6:])
}

func CreatePhoneNumber8(numbers [10]uint) string {
	conv := func(numList []uint) string {
		result := ""
		for _, val := range numList {
			result += strconv.Itoa(int(val))
		}
		return result
	}
	return fmt.Sprintf("(%s) %s-%s", conv(numbers[:3]), conv(numbers[3:6]), conv(numbers[6:]))
}

func CreatePhoneNumber9(numbers [10]uint) string {
	reg, _ := regexp.Compile("[^0-9]+")
	processedString := reg.ReplaceAllString(fmt.Sprint(numbers), "")
	return fmt.Sprintf("(%s) %s-%s", processedString[0:3], processedString[3:6], processedString[6:10])
}

func CreatePhoneNumber10(numbers [10]uint) string {
	var format string = "(xxx) xxx-xxxx"

	for _, v := range numbers {
		format = strings.Replace(format, "x", strconv.Itoa(int(v)), 1)
	}
	return format
}

func CreatePhoneNumber11(numbers [10]uint) string {
	s, _ := json.Marshal(numbers)

	str := strings.ReplaceAll(strings.Trim(string(s), "[]"), ",", "")

	return "(" + str[:3] + ") " + str[3:6] + "-" + str[6:]

}

func CreatePhoneNumber12(numbers [10]uint) string {
	return fmt.Sprintf(
		"(%d%d%d) %d%d%d-%d%d%d%d",
		numbers[0],
		numbers[1],
		numbers[2],
		numbers[3],
		numbers[4],
		numbers[5],
		numbers[6],
		numbers[7],
		numbers[8],
		numbers[9],
	)
}

func CreatePhoneNumber13(numbers [10]uint) string {
	numStr := ""

	for _, num := range numbers {
		numStr += strconv.Itoa(int(num))
	}

	return fmt.Sprintf("(%s) %s-%s", numStr[0:3], numStr[3:6], numStr[6:10])
}

//

func CreatePhoneNumber14(numbers [10]uint) string {
	bytes := make([]byte, len(numbers))
	for i, n := range numbers {
		bytes[i] = byte(48 + n)
	}
	return fmt.Sprintf("(%s) %s-%s", bytes[:3], bytes[3:6], bytes[6:])
}

//

func CreatePhoneNumber15(numbers [10]uint) string {
	var str string
	for _, n := range numbers {
		str += string('0' + n)
	}
	return fmt.Sprintf("(%v) %v-%v", str[:3], str[3:6], str[6:])
}

//

func CreatePhoneNumber16(numbers [10]uint) string {
	var phone string = ""

	for i, _ := range numbers {

		u := strconv.FormatUint(uint64(numbers[i]), 10)

		switch i {
		case 0:
			phone += "("
		case 3:
			phone += ") "
		case 6:
			phone += "-"
		}
		phone += u

	}

	return phone
}

//

func CreatePhoneNumber17(numbers [10]uint) string {
	for _, i := range numbers {
		if i > 9 {
			return ""
		}
	}

	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

//

func CreatePhoneNumber18(numbers [10]uint) string {

	bNums := make([]byte, 10)
	for i, n := range numbers {
		if n < 0 || n > 9 {
			return ""
		}

		bNums[i] = byte(n) + '0'
	}

	return "(" + string(bNums[0:3]) + ") " + string(bNums[3:6]) + "-" + string(bNums[6:])
}

//

func CreatePhoneNumber19(numbers [10]uint) string {
	e := make([]any, 10)
	for i, n := range numbers {
		e[i] = n
	}

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", e...)
}

//

func CreatePhoneNumber20(numbers [10]uint) string {
	var a []string

	for i := range numbers {
		j := (strconv.Itoa(int(numbers[i])))

		a = append(a, j)
	}
	phone := "(" + a[0] + a[1] + a[2] + ") " + a[3] + a[4] + a[5] + "-" + a[6] + a[7] + a[8] + a[9]

	return phone
}

//

func CreatePhoneNumber21(numbers [10]uint) string {
	in := make([]interface{}, len(numbers))
	for k, v := range numbers {
		in[k] = v
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", in...)
}

//

func CreatePhoneNumber22(numbers [10]uint) string {
	first := join(numbers[:3])
	second := join(numbers[3:6])
	third := join(numbers[6:])

	return fmt.Sprintf("(%s) %s-%s", first, second, third)
}

func join(nums []uint) string {
	var result string

	for _, i := range nums {
		result += strconv.Itoa(int(i))
	}

	return result
}

//

func printArray6(slice []uint) string {
	var res string
	for _, value := range slice {
		res += strconv.Itoa(int(value))
	}
	return res
}

func CreatePhoneNumber24(numbers [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", printArray6(numbers[0:3]), printArray6(numbers[3:6]), printArray6(numbers[6:]))
}

//

func CreatePhoneNumber25(numbers [10]uint) string {
	if len(numbers) < 10 {
		return ""
	}
	strnumbers := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ""), "[]")
	prefix, num := fmt.Sprintf("(%s)", strnumbers[:3]), fmt.Sprintf("%s-%s", strnumbers[3:6], strnumbers[6:])

	return fmt.Sprintf("%s %s", prefix, num)
}

//

func CreatePhoneNumber26(numbers [10]uint) string {

	var b strings.Builder
	b.Grow(14)

	for i, n := range numbers {
		switch i {
		case 0:
			fmt.Fprintf(&b, "%s%d", `(`, n)
		case 3:
			fmt.Fprintf(&b, "%s%d", `) `, n)
		case 6:
			fmt.Fprintf(&b, "%s%d", `-`, n)
		default:
			fmt.Fprintf(&b, "%d", n)
		}
	}
	return b.String()
}

//

func CreatePhoneNumber27(numbers [10]uint) string {
	var str string

	for i, number := range numbers {
		switch {
		case i == 0:
			str += fmt.Sprint("(", number)
		case i == 2:
			str += fmt.Sprint(number, ") ")
		case i == 5:
			str += fmt.Sprint(number, "-")
		default:
			str += fmt.Sprint(number)
		}
	}

	return str
}

func CreatePhoneNumber28(numbers [10]uint) string {

	var number string

	n0 := strconv.Itoa(int(numbers[0]))
	n1 := strconv.Itoa(int(numbers[1]))
	n2 := strconv.Itoa(int(numbers[2]))
	n3 := strconv.Itoa(int(numbers[3]))
	n4 := strconv.Itoa(int(numbers[4]))
	n5 := strconv.Itoa(int(numbers[5]))
	n6 := strconv.Itoa(int(numbers[6]))
	n7 := strconv.Itoa(int(numbers[7]))
	n8 := strconv.Itoa(int(numbers[8]))
	n9 := strconv.Itoa(int(numbers[9]))

	number += "(" + n0 + n1 + n2 + ")" + " " + n3 + n4 + n5 + "-" + n6 + n7 + n8 + n9

	return number
}

//

func CreatePhoneNumber29(numbers [10]uint) string {
	var phone string
	phone += "("
	phone += strconv.Itoa(int(numbers[0])) + strconv.Itoa(int(numbers[1])) + strconv.Itoa(int(numbers[2]))
	phone += ")"
	phone += " "
	phone += strconv.Itoa(int(numbers[3])) + strconv.Itoa(int(numbers[4])) + strconv.Itoa(int(numbers[5]))
	phone += "-"
	phone += strconv.Itoa(int(numbers[6])) + strconv.Itoa(int(numbers[7])) + strconv.Itoa(int(numbers[8])) + strconv.Itoa(int(numbers[9]))
	return phone
}

//

func CreatePhoneNumber30(numbers [10]uint) string {
	phone := "("
	for _, v := range numbers {
		if len(phone) == 4 {
			phone += ") "
		}
		if len(phone) == 9 {
			phone += "-"
		}
		phone += strconv.Itoa(int(v))
	}
	return phone
}

//

func CreatePhoneNumber31(nu [10]uint) string {
	return "(" + fmt.Sprint(nu[0]) + fmt.Sprint(nu[1]) + fmt.Sprint(nu[2]) + ") " + fmt.Sprint(nu[3]) + fmt.Sprint(nu[4]) + fmt.Sprint(nu[5]) + "-" + fmt.Sprint(nu[6]) + fmt.Sprint(nu[7]) + fmt.Sprint(nu[8]) + fmt.Sprint(nu[9])
}

//

func CreatePhoneNumber32(numbers [10]uint) string {
	phoneNumber := ""
	var (
		k int
		v uint
	)
	for k, v = range numbers {
		if k == 0 {
			phoneNumber += "("
		}
		if k == 3 {
			phoneNumber += ") "
		}
		if k == 6 {
			phoneNumber += "-"
		}
		phoneNumber = phoneNumber + strconv.Itoa(int(v))
	}

	return phoneNumber
}

//

func CreatePhoneNumber33(numbers [10]uint) string {
	b := []byte("(XXX) XXX-XXXX")
	i := 0
	for j := range b {
		if b[j] == 'X' {
			b[j] = byte(numbers[i] + '0')
			i++
		}
	}
	return string(b)
}

//

func CreatePhoneNumber34(numbers [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[:3][0], numbers[:3][1], numbers[:3][2], numbers[3:6][0], numbers[3:6][1], numbers[3:6][2], numbers[6:][0], numbers[6:][1], numbers[6:][2], numbers[6:][3])
}

//

func CreatePhoneNumber35(numbers [10]uint) string {
	str := "(" + strconv.Itoa(int(numbers[0])) + strconv.Itoa(int(numbers[1])) + strconv.Itoa(int(numbers[2])) + ") " + strconv.Itoa(int(numbers[3])) + strconv.Itoa(int(numbers[4])) + strconv.Itoa(int(numbers[5])) + "-" + strconv.Itoa(int(numbers[6])) + strconv.Itoa(int(numbers[7])) + strconv.Itoa(int(numbers[8])) + strconv.Itoa(int(numbers[9]))
	return str
}

//

func CreatePhoneNumber36(numbers [10]uint) string {
	t := make([]any, 10)
	for i := 0; i < 10; i++ {
		t[i] = numbers[i]
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", t...)
}

//

func CreatePhoneNumber37(numbers [10]uint) string {
	res := "("
	for i, r := range numbers {
		res += strconv.Itoa(int(r))
		if i == 2 {
			res += ") "
		} else if i == 5 {
			res += "-"
		}
	}
	return res
}

//

func CreatePhoneNumber38(numbers [10]uint) string {
	tmp := make([]interface{}, len(numbers))
	for i, val := range numbers {
		tmp[i] = val
	}
	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", tmp...)
}

//

type cellphone struct {
	p1 string
	p2 string
	p3 string
}

func CreatePhoneNumber39(numbers [10]uint) string {
	var phone []string
	phoneStruct := new(cellphone)

	for _, i := range numbers {
		phone = append(phone, strconv.Itoa(int(i)))
	}

	phoneStruct.p1 = strings.Join(phone[0:3], "")
	phoneStruct.p2 = strings.Join(phone[3:6], "")
	phoneStruct.p3 = strings.Join(phone[6:], "")

	return "(" + phoneStruct.p1 + ")" + " " + phoneStruct.p2 + "-" + phoneStruct.p3

}

func CreatePhoneNumber40(numbers [10]uint) string {
	var sb strings.Builder
	sb.WriteString("(")
	for _, i := range numbers[:3] {
		fmt.Println(i)
		sb.WriteString(fmt.Sprint(i))
	}
	sb.WriteString(") ")
	for _, k := range numbers[3:6] {
		sb.WriteString(fmt.Sprint(k))
	}
	sb.WriteString("-")
	for _, y := range numbers[6:10] {
		sb.WriteString(fmt.Sprint(y))
	}

	return sb.String()
}

//

func CreatePhoneNumber41(numbers [10]uint) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("(%v) %v-%v", numbers[0:3], numbers[3:6], numbers[6:]), "[", ""), "]", ""), " ", ""), ")", ") ")
}

func CreatePhoneNumber(numbers [10]uint) string {
	result := "("
	for i := 0; i < 10; i++ {
		if i == 3 {
			result = result + ") "
		}
		if i == 6 {
			result = result + "-"
		}
		result = fmt.Sprintf("%s%d", result, numbers[i])
	}
	return result
}

//

func CreatePhoneNumber42(numbers [10]uint) string {
	return fmt.Sprintf("(%s) %s-%s", Itos(numbers[0:3]), Itos(numbers[3:6]), Itos(numbers[6:10]))
}

func Itos(integers []uint) string {
	str := ""
	for _, i := range integers {
		str += fmt.Sprint(i)
	}
	return str
}

//

func CreatePhoneNumber43(numbers [10]uint) string {
	var phoneNumber string
	for index, num := range numbers {
		switch index {
		case 0:
			phoneNumber += ("(" + string(fmt.Sprint(num)))
		case 3:
			phoneNumber += (") " + string(fmt.Sprint(num)))
		case 6:
			phoneNumber += ("-" + string(fmt.Sprint(num)))
		default:
			phoneNumber += string(fmt.Sprint(num))
		}
	}
	return phoneNumber
}

//

func CreatePhoneNumber44(numbers [10]uint) string {
	// Create a string slice for neater sprintf later.
	stringSlice := make([]string, len(numbers))
	for i, val := range numbers {
		stringSlice[i] = strconv.FormatUint(uint64(val), 10)
	}

	res := strings.Join(stringSlice, "")

	return fmt.Sprintf("(%s) %s-%s", res[0:3], res[3:6], res[6:10])
}

//

func CreatePhoneNumber45(numbers [10]uint) string {

	// n := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// Convert slices to string and format as (xxx) xxx-xxxx
	f := fmt.Sprintf("%d%d%d", numbers[0], numbers[1], numbers[2])
	s := fmt.Sprintf("%d%d%d", numbers[3], numbers[4], numbers[5])
	t := fmt.Sprintf("%d%d%d%d", numbers[6], numbers[7], numbers[8], numbers[9])

	// Return the formatted phone number
	return fmt.Sprintf("(%s) %s-%s", f, s, t)

}

//

func CreatePhoneNumber46(numbers [10]uint) string {
	var snumbers []string
	for _, v := range numbers {
		snumbers = append(snumbers, strconv.Itoa(int(v)))
	}
	res := fmt.Sprint("(", strings.Join(snumbers[:3], ""), ") ", strings.Join(snumbers[3:6], ""), "-", strings.Join(snumbers[6:], ""))
	return res
}

//

func CreatePhoneNumber47(numbers [10]uint) string {
	mapk := make(map[int]string)
	mapk[0] = "("
	mapk[3] = ") "
	mapk[6] = "-"

	result := ""
	for i, num := range numbers {
		if str, ok := mapk[i]; ok {
			result = result + str
		}
		result = fmt.Sprintf("%s%v", result, num)
	}
	return result
}

//

func CreatePhoneNumber48(numbers [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d",
		numbers[0], numbers[1], numbers[2],
		numbers[3], numbers[4], numbers[5],
		numbers[6], numbers[7], numbers[8], numbers[9])
}

func main1() {
	numbers := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(CreatePhoneNumber(numbers))
}

//

// CreatePhoneNumber takes an array of 10 integers and returns a formatted phone number string.
func CreatePhoneNumber49(numbers [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d",
		numbers[0], numbers[1], numbers[2],
		numbers[3], numbers[4], numbers[5],
		numbers[6], numbers[7], numbers[8], numbers[9],
	)
}

//

func CreatePhoneNumber50(numbers [10]uint) string {
	var result strings.Builder
	for i, v := range numbers {
		str := strconv.FormatUint(uint64(v), 10)
		switch {
		case i == 0:
			result.WriteString("(")
			result.WriteString(str)
		case i == 3:
			result.WriteString(") ")
			result.WriteString(str)
		case i == 6:
			result.WriteString("-")
			result.WriteString(str)
		default:
			result.WriteString(str)
		}
	}
	return result.String()
}

//

func CreatePhoneNumber51(numbers [10]uint) string {
	code := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers[:3])), ""), "[]")
	firstPart := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers[3:6])), ""), "[]")
	lastPart := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers[6:10])), ""), "[]")
	return "(" + code + ") " + firstPart + "-" + lastPart
}

//

func CreatePhoneNumber52(numbers [10]uint) string {
	mas := []string{}
	for _, v := range numbers {
		mas = append(mas, strconv.Itoa(int(v)))
	}

	return fmt.Sprintf("(%s%s%s) %s%s%s-%s%s%s%s", mas[0], mas[1], mas[2], mas[3], mas[4], mas[5], mas[6], mas[7], mas[8], mas[9])
}

//

func CreatePhoneNumber53(numbers [10]uint) string {
	var result string = "("
	for i, v := range numbers {

		if i == 3 {
			result += ") "
		}

		if i == 6 {
			result += "-"
		}

		if v == 0 && i < 9 {
			result += "0"
			continue
		}

		result += string(rune(v + '0'))
	}

	return result
}

//

func CreatePhoneNumber54(numbers [10]uint) string {
	//   return
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4],
		numbers[5], numbers[6], numbers[7], numbers[8], numbers[9],
	)
}

//

func IntArrayToString(arr []uint) string {
	var result []string
	for _, num := range arr {
		result = append(result, fmt.Sprintf("%d", num))
	}
	return strings.Join(result, "")
}

func CreatePhoneNumber55(numbers [10]uint) string {
	return fmt.Sprintf("(%s) %s-%s", IntArrayToString(numbers[:3]), IntArrayToString(numbers[3:6]), IntArrayToString(numbers[6:]))
}

//

func CreatePhoneNumber56(numbers [10]uint) string {
	firstPage := ""
	secondPage := ""
	thirdPage := ""
	for i := 0; i < 10; i++ {
		if i < 3 {
			firstPage += strconv.FormatUint(uint64(numbers[i]), 10)
		} else if i >= 3 && i < 6 {
			secondPage += strconv.FormatUint(uint64(numbers[i]), 10)
		} else if i >= 6 {
			thirdPage += strconv.FormatUint(uint64(numbers[i]), 10)
		}
	}
	fullStr := "(" + firstPage + ") " + secondPage + "-" + thirdPage
	return fullStr

}

//

func CreatePhoneNumber57(numbers [10]uint) string {
	var b strings.Builder

	stringSl := make([]string, 10)
	for i, v := range numbers {
		stringSl[i] = strconv.FormatUint(uint64(v), 10)
	}

	b.WriteString("(")
	b.WriteString(strings.Join(stringSl[0:3], ""))
	b.WriteString(") ")
	b.WriteString(strings.Join(stringSl[3:6], ""))
	b.WriteString("-")
	b.WriteString(strings.Join(stringSl[6:10], ""))

	return b.String()

}

//

func CreatePhoneNumber58(numbers [10]uint) string {
	strNum := make([]string, len(numbers))

	for i, val := range numbers {
		strNum[i] = strconv.Itoa(int(val))
	}

	return fmt.Sprintf("(%s) %s-%s", strings.Join(strNum[:3], ""), strings.Join(strNum[3:6], ""), strings.Join(strNum[6:], ""))
}

//

func CreatePhoneNumber59(numbers [10]uint) string {
	var res strings.Builder
	res.Grow(14)

	res.WriteByte('(')
	for i := 0; i < 3; i++ {
		res.WriteByte(byte(numbers[i] + '0'))
	}
	res.WriteByte(')')
	res.WriteByte(' ')
	for i := 3; i < 6; i++ {
		res.WriteByte(byte(numbers[i] + '0'))
	}
	res.WriteByte('-')
	for i := 6; i < 10; i++ {
		res.WriteByte(byte(numbers[i] + '0'))
	}

	return res.String()
}

//

func CreatePhoneNumber60(n [10]uint) string {
	stringNum := make([]string, 10)

	for i, num := range n {
		stringNum[i] = fmt.Sprintf("%d", num)
	}

	part1 := strings.Join(stringNum[0:3], "")
	part2 := strings.Join(stringNum[3:6], "")
	part3 := strings.Join(stringNum[6:], "")

	return fmt.Sprintf("(%s) %s-%s", part1, part2, part3)
}

//

func CreatePhoneNumber61(numbers [10]uint) string {
	finalWord := ""

	strNumbers := make([]string, len(numbers))
	for i, num := range numbers {
		str := strconv.Itoa(int(num))
		if i == 0 {
			str = "(" + str
		} else if i == 2 {
			str = str + ") "
		} else if i == 5 {
			str = str + "-"
		}

		strNumbers[i] = str
	}
	result := strings.Join(strNumbers, "")
	finalWord = result
	return finalWord
}

//

func CreatePhoneNumber62(numbers [10]uint) string {
	result := "("
	for i := 0; i < 3; i++ {
		result += strconv.Itoa(int(numbers[i]))
	}
	result += ") "
	for i := 3; i < len(numbers); i++ {
		if i == 6 {
			result += "-"
		}
		result += strconv.Itoa(int(numbers[i]))
	}

	return result
}

//

func CreatePhoneNumber63(numbers [10]uint) string {
	res := "("
	for i, v := range numbers {

		if i == 3 {
			res += ") "
		} else if i == 6 {
			res += "-"
		}
		res += fmt.Sprint(v)
	}
	return res
}

//

func CreatePhoneNumber64(numbers [10]uint) string {
	phoneNumber := make([]rune, 14)
	phoneNumber[0] = '('
	phoneNumber[1] = rune(48 + numbers[0])
	phoneNumber[2] = rune(48 + numbers[1])
	phoneNumber[3] = rune(48 + numbers[2])
	phoneNumber[4] = ')'
	phoneNumber[5] = ' '
	phoneNumber[6] = rune(48 + numbers[3])
	phoneNumber[7] = rune(48 + numbers[4])
	phoneNumber[8] = rune(48 + numbers[5])
	phoneNumber[9] = '-'
	phoneNumber[10] = rune(48 + numbers[6])
	phoneNumber[11] = rune(48 + numbers[7])
	phoneNumber[12] = rune(48 + numbers[8])
	phoneNumber[13] = rune(48 + numbers[9])
	return string(phoneNumber)
}

//

func CreatePhoneNumber65(numbers [10]uint) string {
	var strNumbers []string
	for _, num := range numbers {
		strNumbers = append(strNumbers, fmt.Sprintf("%d", num))
	}

	areaCode := strings.Join(strNumbers[0:3], "")
	exchange := strings.Join(strNumbers[3:6], "")
	subscriber := strings.Join(strNumbers[6:10], "")

	return fmt.Sprintf("(%s) %s-%s", areaCode, exchange, subscriber)
}

//

func CreatePhoneNumber66(numbers [10]uint) string {
	if len(numbers) != 10 {
		return "Argument invalid"
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

//

func CreatePhoneNumber67(numbers [10]uint) string {
	var pn string = "("
	for i := range numbers {
		pn += fmt.Sprint(numbers[i])
		if i == 2 {
			pn += ") "
		}
		if i == 5 {
			pn += "-"
		}
	}

	return pn
}

//

func CreatePhoneNumber68(numbers [10]uint) string {
	var numbersStr string
	for _, char := range numbers {
		numbersStr = fmt.Sprintf("%s%d", numbersStr, char)
	}
	areaCode := "(" + numbersStr[:3] + ") "
	subCode := numbersStr[3:6] + "-" + numbersStr[6:10]
	return areaCode + subCode
}

//

func CreatePhoneNumber69(numbers [10]uint) string {
	phoneNumber := "("
	for i, num := range numbers {
		phoneNumber += fmt.Sprint(num)
		if i == 2 {
			phoneNumber += ") "
		} else if i == 5 {
			phoneNumber += "-"
		}
	}
	return phoneNumber
}

//

func CreatePhoneNumber70(numbers [10]uint) string {
	var result [14]byte
	result[0], result[4], result[5], result[9] = '(', ')', ' ', '-'
	positions := []int{1, 2, 3, 6, 7, 8, 10, 11, 12, 13}
	for i, pos := range positions {
		result[pos] = byte(numbers[i]) + '0'
	}
	return string(result[:])
}

//

func CreatePhoneNumber71(numbers [10]uint) (res string) {
	for i, v := range numbers {
		if i == 0 {
			res += "("
		}
		if i == 3 {
			res += ") "
		}
		if i == 6 {
			res += "-"
		}
		res += strconv.FormatUint(uint64(v), 10)
	}

	return
}

//

func CreatePhoneNumber72(numbers [10]uint) string {
	stringified := "("
	for i, v := range numbers {
		if i == 2 {
			stringified = stringified + fmt.Sprint(v) + ") "
		} else if i == 5 {
			stringified = stringified + fmt.Sprint(v) + "-"
		} else {
			stringified = stringified + fmt.Sprint(v)
		}
	}
	return stringified
}

//

func CreatePhoneNumber73(numbers [10]uint) (string, error) {
	for _, num := range numbers {
		if num > 9 {
			return "", fmt.Errorf("invalid number in input: %d", num)
		}
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d",
		numbers[0], numbers[1], numbers[2],
		numbers[3], numbers[4], numbers[5],
		numbers[6], numbers[7], numbers[8], numbers[9]), nil
}

//

func CreatePhoneNumber74(numbers [10]uint) string {
	r := strings.NewReplacer("[", "", " ", "", "]", "")

	p1 := r.Replace(fmt.Sprintf("%v", numbers[0:3]))
	p2 := r.Replace(fmt.Sprintf("%v", numbers[3:6]))
	p3 := r.Replace(fmt.Sprintf("%v", numbers[6:]))

	return fmt.Sprintf("(%v) %v-%v", p1, p2, p3)
}

//

func CreatePhoneNumber75(numbers [10]uint) (res string) {
	stringsArr := make([]string, 10)
	for i, number := range numbers {
		stringsArr[i] = fmt.Sprintf("%d", number)
	}
	fmt.Println(stringsArr[0:3])
	part1 := strings.Join(stringsArr[0:3], "")  // First 3 digits
	part2 := strings.Join(stringsArr[3:6], "")  // Next 3 digits
	part3 := strings.Join(stringsArr[6:10], "") // Last 4 digits

	return fmt.Sprintf("(%s) %s-%s", part1, part2, part3)
}

//

func CreatePhoneNumber76(numbers [10]uint) string {
	numbersConverted := []string{}
	for _, val := range numbers {
		numbersConverted = append(numbersConverted, strconv.FormatUint(uint64(val), 10))
	}
	first := numbersConverted[:3]
	second := numbersConverted[3:6]
	third := numbersConverted[6:]
	return fmt.Sprintf("(%s) %s-%s", strings.Join(first, ""), strings.Join(second, ""), strings.Join(third, ""))
}

//

func CreatePhoneNumber77(numbers [10]uint) string {
	var phone strings.Builder

	for _, number := range numbers {
		phone.WriteString(strconv.FormatUint(uint64(number), 10))
	}

	re := regexp.MustCompile(`^(\d{3})(\d{3})(\d{4})$`)
	matches := re.FindStringSubmatch(phone.String())

	return fmt.Sprintf("(%v) %v-%v", matches[1], matches[2], matches[3])
}

//

func CreatePhoneNumber78(numbers [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", toStr(numbers[:3]), toStr(numbers[3:6]), toStr(numbers[6:]))
}

func toStr(arr []uint) (res string) {
	for _, e := range arr {
		res += strconv.Itoa(int(e))
	}
	return res
}

//

func CreatePhoneNumber79(n [10]uint) string {
	a := fmt.Sprintf("(%d%d%d)", n[0], n[1], n[2])
	b := fmt.Sprintf("%d%d%d", n[3], n[4], n[5])
	c := fmt.Sprintf("%d%d%d%d", n[6], n[7], n[8], n[9])

	return fmt.Sprintf("%s %s-%s", a, b, c)
}

//

func CreatePhoneNumber80(numbers [10]uint) string {

	var s string
	s = fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	return s

}

//

func CreatePhoneNumber81(numbers [10]uint) string {
	result_phone := "("
	for i, num := range numbers {
		result_phone += fmt.Sprintf("%d", num)
		if i == 2 {
			result_phone += ") "
		} else if i == 5 {
			result_phone += "-"
		}
	}
	return result_phone
}

//

func CreatePhoneNumber82(numbers [10]uint) string {
	var slice []string
	for i := 0; i < len(numbers); i++ {
		s := strconv.Itoa(int(numbers[i]))
		slice = append(slice, s)
	}
	finale := "(" + string(slice[0]) + string(slice[1]) + string(slice[2]) + ")" + " " + string(slice[3]) + string(slice[4]) + string(slice[5]) + "-" + string(slice[6]) + string(slice[7]) + string(slice[8]) + string(slice[9])
	return finale

}

//

func CreatePhoneNumber83(numbers [10]uint) string {
	var s []string
	for _, v := range numbers {
		s = append(s, strconv.FormatUint(uint64(v), 10))
	}
	return fmt.Sprintf("(%s) %s-%s", strings.Join(s[:3], ""), strings.Join(s[3:6], ""), strings.Join(s[6:], ""))

}

//

func CreatePhoneNumber84(numbers [10]uint) string {
	var result string
	for i := 0; i < 10; i++ {
		if i != 0 && i != 2 && i != 5 {
			result += fmt.Sprintf("%d", numbers[i])
		} else {
			switch i {
			case 0:
				result += "(" + fmt.Sprintf("%d", numbers[i])
			case 2:

				result += fmt.Sprintf("%d", numbers[i]) + ") "
			case 5:
				result += fmt.Sprintf("%d", numbers[i]) + "-"
			}
		}
	}
	return result
}

//

func CreatePhoneNumber85(numbers [10]uint) string {
	bytes := make([]byte, 14)

	bytes[0] = '('
	bytes[4] = ')'
	bytes[5] = ' '
	bytes[9] = '-'

	cursor := 0
	for i := range bytes {
		switch i {
		case 0, 4, 5, 9:
			continue
		default:
			bytes[i] = byte('0' + numbers[cursor])
			cursor++
		}
	}

	return string(bytes)
}

//

func CreatePhoneNumber86(number [10]uint) string {
	res := fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", number[0], number[1], number[2], number[3], number[4], number[5], number[6], number[7], number[8], number[9])
	return res
}

//

func CreatePhoneNumber87(numbers [10]uint) string {
	num1 := strconv.FormatUint(uint64(numbers[0]), 10)
	num2 := strconv.FormatUint(uint64(numbers[1]), 10)
	num3 := strconv.FormatUint(uint64(numbers[2]), 10)
	num4 := strconv.FormatUint(uint64(numbers[3]), 10)
	num5 := strconv.FormatUint(uint64(numbers[4]), 10)
	num6 := strconv.FormatUint(uint64(numbers[5]), 10)
	num7 := strconv.FormatUint(uint64(numbers[6]), 10)
	num8 := strconv.FormatUint(uint64(numbers[7]), 10)
	num9 := strconv.FormatUint(uint64(numbers[8]), 10)
	num10 := strconv.FormatUint(uint64(numbers[9]), 10)

	return "(" + num1 + num2 + num3 + ") " + num4 + num5 + num6 + "-" + num7 + num8 + num9 + num10
}

//

func CreatePhoneNumber88(numbers [10]uint) string {
	code := ""
	second := ""
	last := ""

	for i, v := range numbers {

		if i < 3 {
			code += strconv.Itoa(int(v))
			continue
		}

		if i < 6 {
			second += strconv.Itoa(int(v))
			continue
		}

		if i < 10 {
			last += strconv.Itoa(int(v))
		}
	}

	return fmt.Sprintf("(%v) %v-%v", code, second, last)

}

//

func CreatePhoneNumber89(numbers [10]uint) string {

	n := []string{}

	for _, v := range numbers {
		n = append(n, strconv.Itoa(int(v)))
	}

	f := n[:3]
	s := n[3:6]
	t := n[6:]

	return fmt.Sprintf("(%s) %s-%s", strings.Join(f, ""), strings.Join(s, ""), strings.Join(t, ""))
}

//

func CreatePhoneNumber90(n [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", join1(n[:3]), join1(n[3:6]), join1(n[6:]))
}

func join1(n []uint) string {
	var s string
	for _, val := range n {
		s += strconv.FormatUint(uint64(val), 10)
	}
	return s
}

//

func CreatePhoneNumber91(numbers [10]uint) string {
	a := make([]any, len(numbers))

	for i, n := range numbers {
		a[i] = n
	}

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", a...)
}

//

func CreatePhoneNumber92(numbers [10]uint) string {
	var res string
	for i, num := range numbers {
		if i >= 6 {
			if i == 6 {
				res += "-"
			}
			res += strconv.Itoa(int(num))
		}
		if i >= 3 && i < 6 {
			if i == 3 {
				res += ") "
			}
			res += strconv.Itoa(int(num))
		}

		if i <= 2 {
			if i == 0 {
				res += "("
			}
			res += strconv.Itoa(int(num))
		}
	}
	return res
}

//

func CreatePhoneNumber93(numbers [10]uint) string {
	var numbersString string
	for i, number := range numbers {
		var numberString string
		numberString = strconv.Itoa(int(number))
		if i == 0 {
			numbersString += "("
		}
		if i == 3 {
			numbersString += ")"
			numbersString += " "

		}

		if i == 6 {
			numbersString += "-"

		}

		numbersString += numberString

	}
	return numbersString
}

//

func CreatePhoneNumber94(numbers [10]uint) string {
	var result string
	simbols := "() -"
	for i, v := range numbers {
		switch i {
		case 0:
			result = string(simbols[0]) + toString(v)
		case 2:
			result = result + toString(v) + string(simbols[1:3])
		case 5:
			result = result + toString(v) + string(simbols[3])
		default:
			result = result + toString(v)
		}
	}
	return result
}

func toString(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}

//

func CreatePhoneNumber95(numbers [10]uint) string {
	if len(numbers) == 0 {
		return ""

	}
	nums := make([]string, 10)
	for k, v := range numbers {
		nums[k] = strconv.Itoa(int(v))
	}
	return fmt.Sprintf("(%s) %s-%s", strings.Join(nums[0:3], ""), strings.Join(nums[3:6], ""), strings.Join(nums[6:], ""))
}

//(CreatePhoneNumber(10]uint{1,2,3,4,5,6,7,8,9,0})).To(Equal("(123) 456-7890"))

//

func CreatePhoneNumber96(numbers [10]uint) string {
	join := func(nums []uint) string {
		str := make([]string, len(nums))

		for i, num := range nums {
			str[i] = strconv.FormatUint(uint64(num), 10)
		}
		return strings.Join(str, "")
	}

	return fmt.Sprintf("(%v) %v-%v", join(numbers[0:3]), join(numbers[3:6]), join(numbers[6:10]))
}

//

func CreatePhoneNumber97(numbers [10]uint) string {
	var result string
	for i := range numbers {
		switch i {
		case 0:
			result += "("
			result += strconv.Itoa(int(numbers[i]))
		case 2:
			result += strconv.Itoa(int(numbers[i]))
			result += ") "
		case 5:
			result += strconv.Itoa(int(numbers[i]))
			result += "-"
		default:
			result += strconv.Itoa(int(numbers[i]))
		}
	}
	return result
}

//

func CreatePhoneNumber98(numbers [10]uint) string {
	start := numbers[:3]
	end := numbers[3:]
	resp := fmt.Sprintf("(%d%d%d) ", start[0], start[1], start[2])
	for i := 0; i < len(end); i++ {
		if i == 3 {
			resp += fmt.Sprintf("-%d", end[i])
			continue
		}
		resp += fmt.Sprintf("%d", end[i])
	}
	return resp
}

//

func CreatePhoneNumber99(numbers [10]uint) string {
	res := "("

	for i, n := range numbers {
		res = res + fmt.Sprintf("%d", n)

		if i == 2 {
			res = res + ") "
		} else if i == 5 {
			res = res + "-"
		}
	}

	return res
}

//

func CreatePhoneNumber100(numbers [10]uint) string {
	// Ensure the numbers are between 0 and 9.
	for _, num := range numbers {
		if num < 0 || num > 9 {
			return "Invalid input: all numbers must be between 0 and 9"
		}
	}

	// Format the phone number.
	phoneNumber := fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d",
		numbers[0], numbers[1], numbers[2], // Area code
		numbers[3], numbers[4], numbers[5], // First 3 digits
		numbers[6], numbers[7], numbers[8], numbers[9]) // Last 4 digits

	return phoneNumber

}

//

func CreatePhoneNumber101(numbers [10]uint) string {
	slice := []string{"("}
	for idx, num := range numbers {
		slice = append(slice, strconv.Itoa(int(num)))
		if idx == 2 {
			slice = append(slice, ") ")
		} else if idx == 5 {
			slice = append(slice, "-")
		}
	}
	return strings.Join(slice, "")
}

//

func CreatePhoneNumber102(numbers [10]uint) string {
	res := "("
	for i := 0; i < len(numbers); i++ {
		if i == 3 {
			res = res + ") "
		}
		if i == 6 {
			res = res + "-"
		}
		res = res + strconv.Itoa(int(numbers[i]))
	}
	return res

}

//

func CreatePhoneNumber103(numbers [10]uint) string {
	var code_len, first_part_len int = 3, 3
	var phone_number string = "("

	for i, number := range numbers {
		number_str := fmt.Sprint(number)

		if i == code_len {
			phone_number += ") "
		}
		if i == code_len+first_part_len {
			phone_number += "-"
		}

		phone_number += (number_str)
	}

	return phone_number
}

//

func CreatePhoneNumber104(numbers [10]uint) string {
	var numAny []any

	for _, v := range numbers {
		numAny = append(numAny, v)
	}

	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", numAny...)
}

//

func CreatePhoneNumber105(array [10]uint) string {
	return fmt.Sprintf(
		"(%d%d%d) %d%d%d-%d%d%d%d",
		array[0], array[1], array[2], array[3], array[4],
		array[5], array[6], array[7], array[8], array[9],
	)

}

//

func CreatePhoneNumber106(numbers [10]uint) string {

	var sb strings.Builder
	for index, number := range numbers {
		if index == 0 {
			sb.WriteString("(")
		}
		if index == 3 {
			sb.WriteString(") ")
		}
		if index == 6 {
			sb.WriteString("-")
		}
		sb.WriteString(strconv.Itoa(int(number)))
	}
	return sb.String()
}

//

func listInt2string(numbers []uint) string {
	result := ""
	for _, n := range numbers {
		result += strconv.Itoa(int(n))
	}
	return result
}

func CreatePhoneNumber107(numbers [10]uint) string {
	return fmt.Sprintf("(%s) %s-%s", listInt2string(numbers[0:3]), listInt2string(numbers[3:6]), listInt2string(numbers[6:10]))
}

func CreatePhoneNumber109(numbers [10]uint) string {
	sphone := fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d",
		numbers[0], numbers[1], numbers[2],
		numbers[3], numbers[4], numbers[5],
		numbers[6], numbers[7], numbers[8],
		numbers[9])
	return sphone
}

func CreatePhoneNumber110(numbers [10]uint) string {
	var sb strings.Builder

	sb.WriteString("(")
	for i := 0; i < 3; i++ {
		sb.WriteByte(byte(numbers[i]) + '0')
	}
	sb.WriteString(") ")
	for i := 3; i < 6; i++ {
		sb.WriteByte(byte(numbers[i]) + '0')
	}
	sb.WriteString("-")
	for i := 6; i < 10; i++ {
		sb.WriteByte(byte(numbers[i]) + '0')
	}
	return sb.String()
}

//

func CreatePhoneNumber111(numbers [10]uint) string {
	num := "("
	for i := 0; i < 10; i++ {
		num += strconv.Itoa(int(numbers[i]))
		if i == 2 {
			num += ") "
		}
		if i == 5 {
			num += "-"
		}
	}
	return num
}

//

func CreatePhoneNumber112(numbers [10]uint) string {
	var s string = ""
	for _, v := range numbers {
		s = s + strconv.Itoa(int(v))
	}
	return "(" + s[:3] + ") " + s[3:6] + "-" + s[6:]

}

//

func CreatePhoneNumber113(numbers [10]uint) string {
	return "(" + fmt.Sprintf("%v", numbers[0]) + fmt.Sprintf("%v", numbers[1]) + fmt.Sprintf("%v", numbers[2]) + ") " + fmt.Sprintf("%v", numbers[3]) + fmt.Sprintf("%v", numbers[4]) + fmt.Sprintf("%v", numbers[5]) + "-" + fmt.Sprintf("%v", numbers[6]) + fmt.Sprintf("%v", numbers[7]) + fmt.Sprintf("%v", numbers[8]) + fmt.Sprintf("%v", numbers[9])
}

//

func CreatePhoneNumber114(numbers [10]uint) string {
	var numStr string
	for _, v := range numbers {
		numStr += fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("(%s) %s-%s", numStr[:3], numStr[3:6], numStr[6:])
}

//

func CreatePhoneNumber115(numbers [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", ConvertToString(numbers[:3]), ConvertToString(numbers[3:6]), ConvertToString(numbers[6:]))
}

func ConvertToString116(part []uint) string {
	var result string

	for _, number := range part {
		str := strconv.Itoa(int(number))
		result += str
	}
	return result
}

//

func CreatePhoneNumber117(numbers [10]uint) string {
	// Define the ranges
	// We will need 3 items (one per phone number chunk)
	var ranges = [][]uint{
		numbers[:3],
		numbers[3:6],
		numbers[6:],
	}

	// We'll store the results here
	stringChunks := make([]string, 3)

	for i := 0; i < len(ranges); i++ {
		var converted string

		for j := 0; j < len(ranges[i]); j++ {
			converted += strconv.FormatInt(int64(ranges[i][j]), 10)
		}

		stringChunks[i] = converted
	}

	return fmt.Sprintf("(%v) %v-%v", stringChunks[0], stringChunks[1], stringChunks[2])
}

//

func CreatePhoneNumber118(numbers [10]uint) string {
	phoneTemplate := "(ddd) ddd-dddd"

	for _, digit := range numbers {
		phoneTemplate = strings.Replace(phoneTemplate, "d", strconv.Itoa(int(digit)), 1)
	}

	return phoneTemplate
}

//

func CreatePhoneNumber119(numbers [10]uint) string {
	format := "(xxx) xxx-xxxx"
	for _, num := range numbers {
		str := fmt.Sprintf("%d", num)
		format = strings.Replace(format, "x", str, 1)
	}
	return format
}

//

func CreatePhoneNumber120(numbers [10]uint) string {
	res := "("
	for idx, val := range numbers {
		if idx < 3 {
			res += fmt.Sprintf("%d", val)
		}
		if idx == 3 {
			res += ") "
		}
		if idx >= 3 && idx < 6 {
			res += fmt.Sprintf("%d", val)
		}
		if idx == 6 {
			res += "-"
		}
		if idx >= 6 {
			res += fmt.Sprintf("%d", val)
		}
	}
	return res
}

//

func CreatePhoneNumber121(numbers [10]uint) string {
	var main2 []string
	var holder []string
	for _, num := range numbers {
		holder = append(holder, strconv.Itoa(int(num)))
	}
	main2 = append([]string{"("}, holder[:3]...)
	main2 = append(main2, ")")
	main2 = append(main2, " ")
	main2 = append(main2, holder[3:6]...)
	main2 = append(main2, "-")
	main2 = append(main2, holder[6:]...)
	return strings.Join(main2, "")
}

//

func CreatePhoneNumber122(numbers [10]uint) string {

	var phoneNumber string

	for index, char := range numbers {
		switch {
		case index == 0:
			phoneNumber += "("
		case index == 3:
			phoneNumber += ") "
		case index == 6:
			phoneNumber += "-"
		}
		phoneNumber += strconv.Itoa(int(char))
	}
	return phoneNumber
}

//

func CreatePhoneNumber123(numbers [10]uint) string {
	ss := make([]string, len(numbers))
	for i, num := range numbers {
		ss[i] = fmt.Sprintf("%d", num)
	}
	return fmt.Sprintf("(%s) %s-%s", strings.Join(ss[:3], ""), strings.Join(ss[3:6], ""),
		strings.Join(ss[6:10], ""))
}

//

func CreatePhoneNumber124(numbers [10]uint) string {
	phoneNumber := "("
	for i, number := range numbers {
		if i == 3 {
			phoneNumber += ") "
		}
		if i == 6 {
			phoneNumber += "-"
		}
		temp := strconv.Itoa(int(number))
		phoneNumber += temp
	}
	return phoneNumber
}

//

func CreatePhoneNumber125(numbers [10]uint) string {
	var resultado string

	resultado += "("
	for i := 0; i < 3; i++ {
		resultado += strconv.FormatUint(uint64(numbers[i]), 10)
	}
	resultado += ") "

	for i := 3; i < 6; i++ {
		resultado += strconv.FormatUint(uint64(numbers[i]), 10)
	}
	resultado += "-"

	for i := 6; i < 10; i++ {
		resultado += strconv.FormatUint(uint64(numbers[i]), 10)
	}

	return resultado
}

//

func CreatePhoneNumber126(numbers [10]uint) string {

	ret := fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v",
		numbers[0],
		numbers[1],
		numbers[2],
		numbers[3],
		numbers[4],
		numbers[5],
		numbers[6],
		numbers[7],
		numbers[8],
		numbers[9],
	)

	return ret
}

//

func CreatePhoneNumber127(numbers [10]uint) string {
	phoneNumber := "("
	for i, number := range numbers {
		digit := strconv.Itoa(int(number))
		if i == 0 {
			phoneNumber += digit
		} else if i == 2 {
			phoneNumber += digit + ") "
		} else if i == 5 {
			phoneNumber += digit + "-"
		} else {
			phoneNumber += digit
		}
	}
	return phoneNumber
}

//

func CreatePhoneNumber128(numbers [10]uint) string {
	format := func(numbers []uint) string {
		s := ""
		for _, v := range numbers {
			s += strconv.Itoa(int(v))
		}
		return s
	}

	return fmt.Sprintf("(%s) %s-%s", format(numbers[:3]), format(numbers[3:6]), format(numbers[6:]))
}

//

// func CreatePhoneNumber129(n [10]uint) string {
// 	return "(" + strconv.Itoa(int(n[0])) + strconv.Itoa(int(n[1])) + strconv.Itoa(int(n[2])) + ")" + " " + strconv.Itoa(int(n[3])) + strconv.Itoa(int(n[4])) + strconv.Itoa(int(n[5])) + "-" + strconv.Itoa(int(n[6])) + strconv.Itoa(int(n[7])) + strconv.Itoa(int(n[8])) + strconv.Itoa(int(n[9]))

//

func CreatePhoneNumber130(numbers [10]uint) string {
	v := make([]any, 10)
	for i := range numbers {
		v[i] = numbers[i]
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", v...)
}

//

func CreatePhoneNumber131(numbers [10]uint) string {
	var myString string
	for i := 0; i < 10; i++ {
		myString += strconv.FormatUint(uint64(numbers[i]), 10)
	}

	var result = "(" + myString[0:3] + ") " + myString[3:6] + "-" + myString[6:10]
	return result
}

//

func CreatePhoneNumber133(numbers [10]uint) string {
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	res := "("
	for i, v := range numbers {
		res += nums[v]
		if i == 2 {
			res += ") "
		} else if i == 5 {
			res += "-"
		}
	}
	return res
}

//

func CreatePhoneNumber134(numbers [10]uint) string {
	var numbersStr []string
	for _, value := range numbers {
		valueInt := int(value)
		valueStr := strconv.Itoa(valueInt)
		numbersStr = append(numbersStr, valueStr)
	}
	a := fmt.Sprintf("(%s%s%s) %s%s%s-%s%s%s%s", numbersStr[0], numbersStr[1], numbersStr[2], numbersStr[3], numbersStr[4], numbersStr[5], numbersStr[6], numbersStr[7], numbersStr[8], numbersStr[9])
	return a
}

//

var posToChar = map[int]string{
	0: "(",
	3: ") ",
	6: "-",
}

func CreatePhoneNumber135(numbers [10]uint) string {
	resStr := ""
	for i, num := range numbers {
		if char, ok := posToChar[i]; ok {
			resStr += char
		}
		resStr += fmt.Sprintf("%d", num)
	}
	return resStr
}

//

func CreatePhoneNumber136(numbers [10]uint) string {
	var xByte strings.Builder
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			x := fmt.Sprintf("(%v", numbers[i])
			xByte.WriteString(x)
		} else if i == 2 {
			x := fmt.Sprintf("%v) ", numbers[i])
			xByte.WriteString(x)
		} else if i == 5 {
			x := fmt.Sprintf("%v-", numbers[i])
			xByte.WriteString(x)
		} else {
			xByte.WriteString(fmt.Sprintf("%v", numbers[i]))
		}
	}
	return xByte.String()
}

//

func CreatePhoneNumber137(numbers [10]uint) string {
	var s []string
	for _, v := range numbers {
		s = append(s, strconv.Itoa(int(v)))
	}

	formattedNumber := fmt.Sprintf("(%v) %v-%v", strings.Join(s[:3], ""), strings.Join(s[3:6], ""), strings.Join(s[6:], ""))
	fmt.Println(formattedNumber)
	return formattedNumber
}

//

func CreatePhoneNumber138(numbers [10]uint) string {
	res := "("
	for idx, val := range numbers {
		res += strconv.FormatUint(uint64(val), 10)
		if idx == 2 {
			res += ") "
		}
		if idx == 5 {
			res += "-"
		}
	}
	return res
}

//

func CreatePhoneNumber139(numbers [10]uint) string {
	bob := ""
	for i := 0; i < 10; i++ {
		if i == 0 {
			bob += "("
		}
		if i == 3 {
			bob += ") "
		}
		if i == 6 {
			bob += "-"
		}
		bob += strconv.Itoa(int(numbers[i]))
	}
	return bob
}

//

func CreatePhoneNumber140(numbers [10]uint) string {
	index := "(" + strconv.Itoa(int(numbers[0])) + strconv.Itoa(int(numbers[1])) + strconv.Itoa(int(numbers[2])) + ") "
	index2 := strconv.Itoa(int(numbers[3])) + strconv.Itoa(int(numbers[4])) + strconv.Itoa(int(numbers[5])) + "-"
	index3 := strconv.Itoa(int(numbers[6])) + strconv.Itoa(int(numbers[7])) + strconv.Itoa(int(numbers[8])) + strconv.Itoa(int(numbers[9]))

	return index + index2 + index3
}

//

func CreatePhoneNumber141(numbers [10]uint) string {
	var strNumbers []string
	for _, num := range numbers {
		strNumbers = append(strNumbers, fmt.Sprintf("%d", num))
	}

	return fmt.Sprintf("(%s) %s-%s", strings.Join(strNumbers[0:3], ""), strings.Join(strNumbers[3:6], ""), strings.Join(strNumbers[6:10], ""))
}

//

func CreatePhoneNumber142(numbers [10]uint) string {
	var stringParts [3]string
	numberParts := [3][]uint{
		numbers[:3],
		numbers[3:6],
		numbers[6:],
	}
	for i, nums := range numberParts {
		for _, num := range nums {
			stringParts[i] += strconv.Itoa(int(num))
		}
	}
	return "(" + stringParts[0] + ") " + stringParts[1] + "-" + stringParts[2]
}

//

func CreatePhoneNumber143(numbers [10]uint) string {
	m := make([]string, len(numbers))

	for k, v := range numbers {
		m[k] = fmt.Sprintf("%v", v)
	}

	return "(" + strings.Join(m[:3], "") + ") " + strings.Join(m[3:6], "") + "-" + strings.Join(m[6:], "")
}

//

func CreatePhoneNumber144(numbers [10]uint) string {
	var digits []string
	for _, num := range numbers {
		digits = append(digits, strconv.Itoa(int(num)))
	}

	return fmt.Sprintf("(%s) %s-%s", strings.Join(digits[:3], ""), strings.Join(digits[3:6], ""), strings.Join(digits[6:], ""))
}

//

func CreatePhoneNumber145(numbers [10]uint) string {

	var builder strings.Builder
	builder.Grow(14)

	// Формируем строку вручную
	builder.WriteByte('(')
	builder.WriteByte(byte(numbers[0]) + '0')
	builder.WriteByte(byte(numbers[1]) + '0')
	builder.WriteByte(byte(numbers[2]) + '0')
	builder.WriteByte(')')
	builder.WriteByte(' ')
	builder.WriteByte(byte(numbers[3]) + '0')
	builder.WriteByte(byte(numbers[4]) + '0')
	builder.WriteByte(byte(numbers[5]) + '0')
	builder.WriteByte('-')
	builder.WriteByte(byte(numbers[6]) + '0')
	builder.WriteByte(byte(numbers[7]) + '0')
	builder.WriteByte(byte(numbers[8]) + '0')
	builder.WriteByte(byte(numbers[9]) + '0')

	return builder.String()

}

//

func CreatePhoneNumber1455(numbers [10]uint) string {
	var result string
	for i, a := range numbers {
		if i == 0 {
			result += fmt.Sprintf("(%d", a)
			continue
		}
		if i == 2 {
			result += fmt.Sprintf("%d) ", a)
			continue
		}
		if i == 5 {
			result += fmt.Sprintf("%d-", a)
			continue
		} else {
			result += fmt.Sprintf("%d", a)
		}

	}
	return result
}

//

func CreatePhoneNumber146(numbers [10]uint) string {
	numsFmt := []string{}
	for _, num := range numbers {
		numsFmt = append(numsFmt, strconv.Itoa(int(num)))
	}

	return fmt.Sprintf(
		"(%v) %v-%v",
		numsFmt[0]+numsFmt[1]+numsFmt[2],
		numsFmt[3]+numsFmt[4]+numsFmt[5],
		numsFmt[6]+numsFmt[7]+numsFmt[8]+numsFmt[9],
	)
}

//

func CreatePhoneNumber147(numbers [10]uint) string {
	result := "("
	for i := 0; i < len(numbers); i++ {
		result += fmt.Sprint(numbers[i])
		if i == 2 {
			result += ") "
		}
		if i == 5 {
			result += "-"
		}
	}
	return result
}

//

func CreatePhoneNumber148(numbers [10]uint) (r string) {
	for i, v := range numbers {
		var vString string = string(rune(v + '0'))

		switch i {
		case 0:
			r += "(" + vString
		case 2:
			r += vString + ") "
		case 5:
			r += vString + "-"
		default:
			r += vString
		}
	}

	return
}

//

func CreatePhoneNumber149(numbers [10]uint) string {
	first := stringNumbers(numbers[:3])
	second := stringNumbers(numbers[3:6])
	third := stringNumbers(numbers[6:10])
	return fmt.Sprintf("(%s) %s-%s", first, second, third)
}

func stringNumbers(numbers []uint) string {
	var buffer bytes.Buffer
	for _, v := range numbers {
		buffer.WriteString(fmt.Sprintf("%d", v))
	}
	return buffer.String()
}

//

func CreatePhoneNumber150(numbers [10]uint) string {
	prefixPhoneNumber := []string{}
	midPrefixPhoneNumber := []string{}
	sufixPhoneNumber := []string{}
	halfFormatedPhoneNumber := []string{}
	formatedPhoneNumber := []string{}

	// Convert the first 3 elements to strings and append to the slice
	for _, num := range numbers[:3] {
		prefixPhoneNumber = append(prefixPhoneNumber, fmt.Sprintf("%d", num))
	}

	for _, v := range numbers[3:6] {
		midPrefixPhoneNumber = append(midPrefixPhoneNumber, fmt.Sprintf("%d", v))
	}

	for _, v := range numbers[6:] {
		sufixPhoneNumber = append(sufixPhoneNumber, fmt.Sprintf("%d", v))
	}

	formatedPrefixPhoneNumber := fmt.Sprintf("(%s)", strings.Join(prefixPhoneNumber, ""))
	formatedMidSectionPhoneNumber := strings.Join(midPrefixPhoneNumber, "")
	formattedSuffixPhoneNumber := strings.Join(sufixPhoneNumber, "")

	// append the mid section and last section of the lists of numbers
	halfFormatedPhoneNumber = append(halfFormatedPhoneNumber, formatedMidSectionPhoneNumber)
	halfFormatedPhoneNumber = append(halfFormatedPhoneNumber, formattedSuffixPhoneNumber)
	formatedHalfFormatedPhoneNumber := strings.Join(halfFormatedPhoneNumber, "-")

	formatedPhoneNumber = append(formatedPhoneNumber, formatedPrefixPhoneNumber)
	formatedPhoneNumber = append(formatedPhoneNumber, formatedHalfFormatedPhoneNumber)

	return strings.Join(formatedPhoneNumber, " ")
}

//

func CreatePhoneNumber151(numbers [10]uint) string {
	var phoneNumber string = ""
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			phoneNumber += "("
		} else if i == 3 {
			phoneNumber += ") "
		} else if i == 6 {
			phoneNumber += "-"
		}

		phoneNumber += string(rune(numbers[i] + '0'))
	}
	return phoneNumber
}

//

func CreatePhoneNumber152(numbers [10]uint) string {
	// Your code here!
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

//

func CreatePhoneNumber153(numbers [10]uint) string {
	var inputBuffer bytes.Buffer
	inputBuffer.WriteString("(")

	for i := 0; i < len(numbers); i++ {
		currentNumber := numbers[i]
		if i == 3 {
			inputBuffer.WriteString(
				fmt.Sprintf(") %d", currentNumber),
			)
		} else if i == 6 {
			inputBuffer.WriteString(
				fmt.Sprintf("-%d", currentNumber),
			)
		} else {
			inputBuffer.WriteString(fmt.Sprintf("%d", currentNumber))
		}
	}

	return inputBuffer.String()
}

//

func joinNumbers(numbers []uint) string {
	var stringBuilder strings.Builder
	for _, num := range numbers {
		stringBuilder.WriteString(strconv.Itoa(int(num)))
	}
	return stringBuilder.String()
}

func CreatePhoneNumber154(numbers [10]uint) string {
	areaCode := joinNumbers(numbers[:3])
	centralOfficeCode := joinNumbers(numbers[3:6])
	lineNumber := joinNumbers(numbers[6:])

	return fmt.Sprintf("(%s) %s-%s", areaCode, centralOfficeCode, lineNumber)
}

//

func CreatePhoneNumber155(numbers [10]uint) string {
	convSlice := make([]any, 0, len(numbers))
	for _, v := range numbers {
		convSlice = append(convSlice, strconv.Itoa(int(v)))
	}
	return fmt.Sprintf("(%s%s%s) %s%s%s-%s%s%s%s", convSlice...)
}

//

func CreatePhoneNumber156(numbers [10]uint) string {

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}
func main() {
	arr := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(strings.Fields(CreatePhoneNumber(arr)))
}

//

func CreatePhoneNumber157(numbers [10]uint) string {
	var phonenum string
	for i, v := range numbers {
		xyi := ""
		switch {
		case i == 0:
			xyi = "("
		case i == 3:
			xyi = ") "
		case i == 6:
			xyi = "-"
		}
		phonenum += xyi + (strconv.FormatUint(uint64(v), 10))
	}
	return phonenum
}

//

func toString2(arr []uint) string {
	str := strings.Trim(fmt.Sprintf("%v", arr), "[]")
	return strings.ReplaceAll(str, " ", "")
}

func CreatePhoneNumber1588(numbers [10]uint) string {
	return fmt.Sprintf("(%s) %s-%s", toString2(numbers[0:3]), toString2(numbers[3:6]), toString2(numbers[6:10]))
}

//

func CreatePhoneNumber158(numbers [10]uint) string {
	StrPhoneNumber := fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	return StrPhoneNumber
}

//

func CreatePhoneNumber159(numbers [10]uint) string {
	phoneNumber := "("
	for i, num := range numbers {
		phoneNumber += fmt.Sprintf("%v", num)
		switch i {
		case 2:
			phoneNumber += ") "
		case 5:
			phoneNumber += "-"
		}
	}
	return phoneNumber
}

//

func CreatePhoneNumber160(numbers [10]uint) string {
	var s string
	for i, n := range numbers {
		if i == 0 {
			s += "("
		}
		s += fmt.Sprint(n)
		if i == 2 {
			s += ") "
		}
		if i == 5 {
			s += "-"
		}
	}
	return s
}

//

func CreatePhoneNumber161(numbers [10]uint) string {
	resp := "("
	for i, n := range numbers {
		a := fmt.Sprintf("%d", n)
		resp += a
		if i == 2 {
			resp += ") "
		} else if i == 5 {
			resp += "-"
		}
	}
	return resp
}

//

func CreatePhoneNumber162(numbers [10]uint) string {
	number := make([]byte, 1, 14)
	number[0] = '('
	for i, _ := range numbers {
		number = append(number, '0'+byte(numbers[i]))
		if i == 2 {
			number = append(number, ')')
			number = append(number, ' ')
		}
		if i == 5 {
			number = append(number, '-')
		}
	}
	return string(number[:])
}

//

func CreatePhoneNumber163(n [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
}

func F(x uint) string {
	return strconv.FormatUint(uint64(x), 10)

}

//

func CreatePhoneNumber164(numbers [10]uint) string {
	var phoneNumber strings.Builder
	for i, n := range numbers {
		switch i {
		case 0:
			phoneNumber.WriteString("(")
		case 3:
			phoneNumber.WriteString(") ")
		case 6:
			phoneNumber.WriteString("-")
		}
		phoneNumber.WriteString(strconv.FormatUint(uint64(n), 10))
	}
	return phoneNumber.String()
}

//

func CreatePhoneNumber165(numbers [10]uint) string {
	var arr [10]string
	for index, value := range numbers {
		v := strconv.FormatUint(uint64(value), 10)
		arr[index] = v
	}
	return `(` + arr[0] + arr[1] + arr[2] + `) ` + arr[3] + arr[4] + arr[5] + `-` + arr[6] + arr[7] + arr[8] + arr[9]
}

//

func CreatePhoneNumber166(numbers [10]uint) string {
	str := ""
	for _, n := range numbers {
		str = fmt.Sprintf("%s%d", str, n)
	}
	return "(" + str[:3] + ") " + str[3:6] + "-" + str[6:]

}

//

func CreatePhoneNumber167(numbers [10]uint) string {
	foo := "("
	for i, num := range numbers {
		str := strconv.FormatUint(uint64(num), 10)
		if i == 3 {
			foo += ") " + str
		} else if i == 6 {
			foo += "-" + str
		} else {
			foo += str
		}
	}
	return foo

}

//

func CreatePhoneNumber168(numbers [10]uint) string {
	var ans string = "("

	for i, num1 := range numbers {
		var num int
		num = int(num1)
		if i < 2 {
			ans += strconv.Itoa(num)
			fmt.Println(ans)

		}
		if i == 2 {
			ans += strconv.Itoa(num) + ") "
		}
		if i < 5 && i > 2 {
			ans += strconv.Itoa(num)

		}
		if i == 5 {
			ans += strconv.Itoa(num) + "-"
		}
		if i <= 9 && i > 5 {
			ans += strconv.Itoa(num)

		}

	}

	return ans

}

//

func CreatePhoneNumber169(numbers [10]uint) string {
	x := new(strings.Builder)
	for i, v := range numbers {
		if i == 0 {
			x.WriteString("(")
			x.WriteString(strconv.Itoa(int(v)))
		} else if i == 2 {
			x.WriteString(strconv.Itoa(int(v)))
			x.WriteString(")")
			x.WriteString(" ")
		} else if i == 6 {
			x.WriteString("-")
			x.WriteString(strconv.Itoa(int(v)))
		} else {
			x.WriteString(strconv.Itoa(int(v)))
		}
	}

	return x.String()
}

//

func CreatePhoneNumber170(numbers [10]uint) string {
	var area_code, prefix, subscriber string
	for i, v := range numbers {
		switch {
		case i < 3:
			area_code += fmt.Sprintf("%d", v)
		case i < 6:
			prefix += fmt.Sprintf("%d", v)
		default:
			subscriber += fmt.Sprintf("%d", v)
		}
	}

	return fmt.Sprintf("(%s) %s-%s", area_code, prefix, subscriber)
}

//

func CreatePhoneNumber171(numbers [10]uint) string {
	nums := make([]interface{}, len(numbers))
	for i, v := range numbers {
		nums[i] = v
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", nums...)
}

//

func CreatePhoneNumber172(numbers [10]uint) string {
	var sl1 []uint = numbers[0:3]
	var sl2 []uint = numbers[3:6]
	var sl3 []uint = numbers[6:10]
	first := "(" + convertUItoS(sl1) + ") "
	second := convertUItoS(sl2) + "-"
	third := convertUItoS(sl3)
	concat := first + second + third
	return concat
}

func convertUItoS(sl []uint) string {
	conv := ""
	for i := 0; i < len(sl); i++ {
		str := strconv.Itoa(int(sl[i]))
		conv += str
	}
	return conv
}

//

func SliceToString(nums []uint) string {
	fs := ""

	for _, num := range nums {
		fs += strconv.FormatUint(uint64(num), 10)
	}

	return fs
}

func CreatePhoneNumber173(numbers [10]uint) string {
	str := fmt.Sprintf("(%+v) %+v-%+v", SliceToString(numbers[0:3]), SliceToString(numbers[3:6]), SliceToString(numbers[6:10]))
	return str
}

//

func CreatePhoneNumber174(numbers [10]uint) string {
	var a string = "("
	for i := 0; i < 10; i++ {
		if i == 3 {
			a = a + ") "

		}
		if i == 6 {
			a += "-"
		}

		a = a + strconv.FormatUint(uint64(numbers[i]), 10)
	}
	return a
}

//

func getPart(s []string, first int, last int) string {
	return strings.Join(s[first:last], "")
}

func CreatePhoneNumber175(numbers [10]uint) string {
	chars := []string{}

	for _, n := range numbers {
		chars = append(chars, strconv.Itoa(int(n)))
	}

	firstPart := getPart(chars, 0, 3)
	secondPart := getPart(chars, 3, 6)
	thirdPart := getPart(chars, 6, 10)

	return fmt.Sprintf("(%v) %v-%v", firstPart, secondPart, thirdPart)
}

//

func CreatePhoneNumber176(numbers [10]uint) string {
	format := "("
	for _, v := range numbers[0:3] {
		format += fmt.Sprintf("%d", v)
	}
	format += ") "

	for _, v := range numbers[3:6] {
		format += fmt.Sprintf("%d", v)
	}
	format += "-"

	for _, v := range numbers[6:] {
		format += fmt.Sprintf("%d", v)
	}

	return format
}

//

func CreatePhoneNumber177(a [10]uint) string {
	str := "("
	for i := range a {
		switch i {
		case 0, 1, 2:
			str += fmt.Sprintf("%d", a[i])
		case 3:
			str += ") "
			str += fmt.Sprintf("%d", a[i])
		case 4, 5:
			str += fmt.Sprintf("%d", a[i])
		case 6:
			str += "-"
			str += fmt.Sprintf("%d", a[i])
		case 7, 8, 9:
			str += fmt.Sprintf("%d", a[i])
		default:
			fmt.Println("Error")
		}
	}
	return str
}

//

func CreatePhoneNumber178(n [10]uint) string {

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", []interface{}{
		n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9]}...)

}

//

func CreatePhoneNumber179(numbers [10]uint) string {
	var stringNumbers []string
	for _, value := range numbers {
		stringNumbers = append(stringNumbers, strconv.FormatUint(uint64(value), 10))
	}
	return fmt.Sprintf("(%v) %v-%v", strings.Join(stringNumbers[0:3], ""), strings.Join(stringNumbers[3:6], ""), strings.Join(stringNumbers[6:], ""))
}

//

func CreatePhoneNumber180(numbers [10]uint) string {
	numbersStr := [10]string{}

	for i := range numbers {
		numbersStr[i] = strconv.Itoa(int(numbers[i]))
	}

	return fmt.Sprintf("(%s) %s-%s", strings.Join(numbersStr[0:3], ""), strings.Join(numbersStr[3:6], ""), strings.Join(numbersStr[6:10], ""))
}

//

func CreatePhoneNumber181(numbers [10]uint) string {
	slice := []rune("(123) 456-7890")
	for i, n := range numbers {
		switch i {
		case 0:
			slice[1] = rune(n) + '0'
		case 1:
			slice[2] = rune(n) + '0'
		case 2:
			slice[3] = rune(n) + '0'
		case 3:
			slice[6] = rune(n) + '0'
		case 4:
			slice[7] = rune(n) + '0'
		case 5:
			slice[8] = rune(n) + '0'
		case 6:
			slice[10] = rune(n) + '0'
		case 7:
			slice[11] = rune(n) + '0'
		case 8:
			slice[12] = rune(n) + '0'
		case 9:
			slice[13] = rune(n) + '0'
		}
	}
	return string(slice)
}

//

func CreatePhoneNumber182(numbers [10]uint) string {
	result_string := []string{}
	result_string = append(result_string, "(")
	for i, v := range numbers {
		result_string = append(result_string, strconv.Itoa(int(v)))
		if i == 2 {
			result_string = append(result_string, ")", " ")
		}
		if i == 5 {
			result_string = append(result_string, "-")
		}
	}
	return strings.Join(result_string, "")
}

//

func CreatePhoneNumber183(numbers [10]uint) string {
	s := ""
	for _, n := range numbers {
		s = fmt.Sprintf("%s%d", s, n)
	}
	return fmt.Sprintf("(%s) %s-%s", s[0:3], s[3:6], s[6:len(numbers)])
}

//

func CreatePhoneNumber184(numbers [10]uint) string {
	final := fmt.Sprintf("(%s) %s-%s",
		arrayToString(numbers[0:3], ""),
		arrayToString(numbers[3:6], ""),
		arrayToString(numbers[6:10], ""),
	)

	return final
}
func arrayToString(a []uint, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

//

func CreatePhoneNumber185(numbers [10]uint) (tel string) {
	for i := range numbers {
		switch i {
		case 0:
			tel += "("
		case 3:
			tel += ") "
		case 6:
			tel += "-"
		}
		tel += strconv.Itoa(int(numbers[i]))
	}
	return
}

//

func CreatePhoneNumber186(n [10]uint) (tel string) {
	for i := range n {
		if i == 0 {
			tel += "("
		}
		tel += fmt.Sprint(n[i])
		if i == 2 {
			tel += ") "
		}
		if i == 5 {
			tel += "-"
		}
	}
	return
}

//

func CreatePhoneNumber187(numbers [10]uint) string {
	numberPhone := []string{}
	for _, num := range numbers {
		numberPhone = append(numberPhone, strconv.Itoa(int(num)))
	}
	return "(" + strings.Join(numberPhone[:3], "") + ") " + strings.Join(numberPhone[3:6], "") + "-" + strings.Join(numberPhone[6:], "")
}

//

func CreatePhoneNumber188(numbers [10]uint) string {
	numString := ""
	for _, num := range numbers {
		numString += strconv.FormatUint(uint64(num), 10)
	}
	return "(" + numString[:3] + ") " + numString[3:6] + "-" + numString[6:]
}

//

func CreatePhoneNumber189(numbers [10]uint) string {
	template := "(%s) %s-%s"
	var first string
	var second string
	var third string

	for i, number := range numbers {
		if i < 3 {
			first += strconv.Itoa(int(number))
			continue
		}
		if i < 6 {
			second += strconv.Itoa(int(number))
			continue
		}
		if i < 10 {
			third += strconv.Itoa(int(number))
			continue
		}
	}

	return fmt.Sprintf(template, first, second, third)
}

//

func CreatePhoneNumber190(numbers [10]uint) string {
	var number string
	for i, v := range numbers {
		if i == 0 {
			number += "("
		}
		if i == 3 {
			number += ") "
		}
		if i == 6 {
			number += "-"
		}
		number += strconv.Itoa(int(v))
	}

	return number
}

//

func CreatePhoneNumber191(numbers [10]uint) string {
	stringedNums := make([]string, len(numbers))
	for i, num := range numbers {
		stringedNums[i] = fmt.Sprintf("%d", num)
	}
	// Format the numbers into (xxx) xxx-xxxx

	return fmt.Sprintf("(%s%s%s) %s%s%s-%s%s%s%s", stringedNums[0], stringedNums[1], stringedNums[2], stringedNums[3], stringedNums[4], stringedNums[5], stringedNums[6], stringedNums[7], stringedNums[8], stringedNums[9])
}

//

func CreatePhoneNumber192(numbers [10]uint) string {
	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

func main3() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

//

func CreatePhoneNumber193(numbers [10]uint) string {

	stringsSlice := make([]string, len(numbers))

	for index, value := range numbers {
		stringsSlice[index] = strconv.FormatUint(uint64(value), 10)
	}

	part1 := "(" + strings.Join(stringsSlice[0:3], "") + ")"
	part2 := " " + strings.Join(stringsSlice[3:6], "") + "-"
	part3 := strings.Join(stringsSlice[6:], "")

	return part1 + part2 + part3

}

//

func CreatePhoneNumber194(n [10]uint) string {
	res := fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
	return res
}

//

func CreatePhoneNumber195(numbers [10]uint) string {

	v1 := fmt.Sprintf("%d%d%d", numbers[0], numbers[1], numbers[2])
	v2 := fmt.Sprintf("%d%d%d", numbers[3], numbers[4], numbers[5])
	v3 := fmt.Sprintf("%d%d%d%d", numbers[6], numbers[7], numbers[8], numbers[9])
	return fmt.Sprintf("(%s) %s-%s", v1, v2, v3)
}

//

func CreatePhoneNumber196(numbers [10]uint) string {
	// (123) 456-7890
	var formattedNumber string
	for _, number := range numbers {
		formattedNumber += strconv.Itoa(int(number))
	}
	return fmt.Sprintf("(%v) %v-%v", formattedNumber[0:3], formattedNumber[3:6], formattedNumber[6:])
}

//

func CreatePhoneNumber197(numbers [10]uint) string {
	s := ""
	for _, n := range numbers {
		s += fmt.Sprintf("%d", n)
	}
	return fmt.Sprintf("(%s) %s-%s", s[0:3], s[3:6], s[6:10])
}

//

func CreatePhoneNumber198(numbers [10]uint) string {

	var strNumbers []string
	for _, num := range numbers {
		strNumbers = append(strNumbers, strconv.Itoa(int(num)))
	}

	phoneNumber := fmt.Sprintf("(%s%s%s) %s%s%s-%s%s%s%s",
		strNumbers[0], strNumbers[1], strNumbers[2],
		strNumbers[3], strNumbers[4], strNumbers[5],
		strNumbers[6], strNumbers[7], strNumbers[8], strNumbers[9],
	)

	return phoneNumber
}

//

func CreatePhoneNumber199(numbers [10]uint) string {
	numberString := ""
	for _, num := range numbers {
		numberString += strconv.Itoa(int(num))
	}

	return fmt.Sprintf("(%s) %s-%s", numberString[0:3], numberString[3:6], numberString[6:10])
}

//

func CreatePhoneNumber200(numbers [10]uint) (phoneNumber string) {
	var a []any
	for _, number := range numbers {
		a = append(a, number)
	}

	phoneNumber = fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", a...)
	return
}

//

func CreatePhoneNumber201(numbers [10]uint) string {

	concat := ""
	for _, v := range numbers {
		concat += strconv.Itoa(int(v))
	}

	log.Println(concat)

	result := "(" + string(concat[0:3]) + ")" + " " + string(concat[3:6]) + "-" + string(concat[6:])
	return result
}

//

func CreatePhoneNumber202(numbers [10]uint) string {
	var a string
	for i := 0; i < len(numbers); i++ {
		a += strconv.Itoa(int(numbers[i]))
	}
	a = "(" + a[:3] + ") " + a[3:6] + "-" + a[6:]
	return a
}

//

func CreatePhoneNumber203(numbers [10]uint) string {
	str := "("
	for i := 0; i < 3; i++ {
		str += strconv.Itoa(int(numbers[i]))
	}
	str += ") "
	for i := 3; i < 6; i++ {
		str += strconv.Itoa(int(numbers[i]))
	}
	str += "-"
	for i := 6; i < 10; i++ {
		str += strconv.Itoa(int(numbers[i]))
	}
	return str
}

//

func CreatePhoneNumber204(numbers [10]uint) string {
	//loop through the array
	phoneNumber := "("
	for i := 0; i < len(numbers); i++ {
		if i <= 2 {
			phoneNumber += strconv.Itoa(int(numbers[i]))
			//add the last bracket after the first 3 numbers
			if i == 2 {
				phoneNumber += ") "
			}
			continue
		}

		if i > 2 {
			//add the second set of three numbers
			phoneNumber += strconv.Itoa(int(numbers[i]))
			if i == 5 {
				phoneNumber += "-"
			}
			continue
		}

		phoneNumber += strconv.Itoa(int(numbers[i]))
	}

	return phoneNumber
}

//

func CreatePhoneNumber205(numbers [10]uint) string {
	var str = strings.ReplaceAll(fmt.Sprint(numbers), " ", "")[1 : len(numbers)+1]
	var ret = "(" + str[:3] + ") " + str[3:6] + "-" + str[6:]
	return ret
}

//

func CreatePhoneNumber206(numbers [10]uint) string {
	result := ""
	count := 0
	for _, i := range numbers {
		if count == 0 {
			result += "("
		} else if count == 3 {
			result += ") "
		} else if count == 6 {
			result += "-"
		}
		result += strconv.Itoa(int(i))
		count += 1
	}
	return result
}

//

func CreatePhoneNumber207(numbers [10]uint) string {
	out := "("
	for ind, el := range numbers {
		if ind == 3 {
			out += ") " + strconv.Itoa(int(el))
			continue
		}
		if ind == 6 {
			out += "-" + strconv.Itoa(int(el))
			continue
		}
		out += strconv.Itoa(int(el))
	}

	return out

}

//

func CreatePhoneNumber208(numbers [10]uint) string {
	var first string
	var second string
	var end string

	for k, v := range numbers {
		if k >= 0 && k < 3 {
			first += strconv.Itoa(int(v))
		}

		if k >= 3 && k < 6 {
			second += strconv.Itoa(int(v))
		}

		if k >= 6 && k < 10 {
			end += strconv.Itoa(int(v))
		}
	}

	return fmt.Sprintf("(%s) %s-%s", first, second, end)
}

//

func CreatePhoneNumber209(numbers [10]uint) string {
	var result string = "("
	for i := 0; i < len(numbers); i++ {
		if i == 3 {
			result += ") "
		}
		if i == 6 {
			result += "-"
		}
		result += fmt.Sprintf("%d", numbers[i])
	}
	return result
}

//

func ConvertToString(array []uint) string {
	string := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ""), "[]")
	return string
}

func CreatePhoneNumber210(numbers [10]uint) string {
	result := "(" + ConvertToString(numbers[:3]) + ") " + ConvertToString(numbers[3:6]) + "-" + ConvertToString(numbers[6:])
	return result
}

//

func fmap[T1 any, T2 any](fn func(T1) T2, input []T1) []T2 {
	res := make([]T2, len(input))
	for i, item := range input {
		res[i] = fn(item)
	}
	return res
}

func CreatePhoneNumber211(numbers [10]uint) string {
	uintToString := func(i uint) string {
		return fmt.Sprintf("%d", i)
	}

	str := strings.Join(fmap(uintToString, numbers[:]), "")
	return fmt.Sprintf("(%s) %s-%s", str[:3], str[3:6], str[6:])
}

//

func split(numbers [10]uint) []string {
	return []string{
		toString3(numbers[:3]),
		toString3(numbers[3:6]),
		toString3(numbers[6:]),
	}
}

func toString3(numbers []uint) string {
	var b strings.Builder
	for _, num := range numbers {
		fmt.Fprintf(&b, "%d", num)
	}
	return b.String()
}

func CreatePhoneNumber212(numbers [10]uint) string {
	parts := split(numbers)
	return fmt.Sprintf("(%s) %s-%s", parts[0], parts[1], parts[2])
}

//

func CreatePhoneNumber213(numbers [10]uint) string {

	phoneNumber := "(" +
		string(rune(numbers[0]+'0')) +
		string(rune(numbers[1]+'0')) +
		string(rune(numbers[2]+'0')) +
		") " +
		string(rune(numbers[3]+'0')) +
		string(rune(numbers[4]+'0')) +
		string(rune(numbers[5]+'0')) +
		"-" +
		string(rune(numbers[6]+'0')) +
		string(rune(numbers[7]+'0')) +
		string(rune(numbers[8]+'0')) +
		string(rune(numbers[9]+'0'))

	return phoneNumber

}

//

func CreatePhoneNumber214(numbers [10]uint) string {
	a := make([]any, len(numbers))
	for i, num := range numbers {
		a[i] = num
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", a...)
}

//

func CreatePhoneNumber215(numbers [10]uint) string {
	res := "("
	for i, v := range numbers {
		res += strconv.Itoa(int(v))
		if i == 2 {
			res += ") "
		} else if i == 5 {
			res += "-"
		}
	}
	return res
}

//

func CreatePhoneNumber216(numbers [10]uint) string {
	var areaCode string
	for _, n := range numbers[:3] {
		areaCode += strconv.FormatUint(uint64(n), 10)
	}

	var phonePrefix string
	for _, n := range numbers[3:6] {
		phonePrefix += strconv.FormatUint(uint64(n), 10)
	}

	var lineNumber string
	for _, n := range numbers[6:] {
		lineNumber += strconv.FormatUint(uint64(n), 10)
	}

	return fmt.Sprintf("(%s) %s-%s", areaCode, phonePrefix, lineNumber)
}

//

func CreatePhoneNumber217(numbers [10]uint) string {
	var out strings.Builder
	out.WriteString("(")
	for _, val := range numbers[:3] {
		out.WriteString(strconv.FormatUint(uint64(val), 10))
	}
	out.WriteString(") ")
	for _, val := range numbers[3:6] {
		out.WriteString(strconv.FormatUint(uint64(val), 10))
	}
	out.WriteString("-")
	for _, val := range numbers[6:len(numbers)] {
		out.WriteString(strconv.FormatUint(uint64(val), 10))
	}
	return out.String()
}

//

func CreatePhoneNumber218(numbers [10]uint) (output string) {
	for i, el := range numbers {
		switch {
		case i == 0:
			output += "("
			output += fmt.Sprintf("%d", el)
		case i == 2:
			output += fmt.Sprintf("%d", el)
			output += ")"
			output += " "
		case i == 5:
			output += fmt.Sprintf("%d", el)
			output += "-"
		default:
			output += fmt.Sprintf("%d", el)
		}
	}
	return output
}

//

func CreatePhoneNumber219(numbers [10]uint) string {
	return "(" + strings.Trim(strings.ReplaceAll(fmt.Sprint(numbers[0:3]), " ", ""), "[]") + ") " + strings.Trim(strings.ReplaceAll(fmt.Sprint(numbers[3:6]), " ", ""), "[]") + "-" + strings.Trim(strings.ReplaceAll(fmt.Sprint(numbers[6:]), " ", ""), "[]")
}

//

func CreatePhoneNumber220(numbers [10]uint) string {

	return fmt.Sprintf("(%s) %s-%s",
		uintSliceToString(numbers[0:3]),
		uintSliceToString(numbers[3:6]),
		uintSliceToString(numbers[6:]),
	)

}

func uintSliceToString(slice []uint) string {
	str := ""
	for _, num := range slice {
		str += strconv.FormatUint(uint64(num), 10)
	}
	return str

}

//

func CreatePhoneNumber222(numbers [10]uint) string {
	fmt.Printf("%v", numbers)
	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

//

func CreatePhoneNumber223(n [10]uint) string {
	return "(" + JoinN(&n, 0, 2) + ") " + JoinN(&n, 3, 5) + "-" + JoinN(&n, 6, 9)
}

func JoinN(n *[10]uint, f uint, t uint) (word string) {
	for i := f; i <= t; i++ {
		word = word + strconv.Itoa(int(n[i]))
	}
	return
}

//

func CreatePhoneNumber224(numbers [10]uint) string {
	var numbersStrings []string
	for _, i := range numbers {
		numbersStrings = append(numbersStrings, strconv.Itoa(int(i)))
	}
	var numbersString = strings.Join(numbersStrings, "")
	return fmt.Sprintf("(%s) %s-%s", numbersString[0:3], numbersString[3:6], numbersString[6:10])
}

//

func CreatePhoneNumber225(numbers [10]uint) string {
	tmp := make([]any, 10)
	for idx, number := range numbers {
		tmp[idx] = number
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", tmp...)
}

//

func CreatePhoneNumber226(numbers [10]uint) string {
	var result string
	result = "("
	for index, element := range numbers {
		if index == 3 {
			result += ") "
		}
		if index == 6 {
			result += "-"
		}
		result += strconv.Itoa(int(element))
	}
	return result
}

//

func CreatePhoneNumber227(numbers [10]uint) string {
	answer := "("
	for _, value := range numbers[:3] {
		answer += strconv.Itoa(int(value))
	}
	answer += ") "

	for _, value := range numbers[3:6] {
		answer += strconv.Itoa(int(value))
	}
	answer += "-"

	for _, value := range numbers[6:] {
		answer += strconv.Itoa(int(value))
	}
	return answer
}

//

func CreatePhoneNumber228(numbers [10]uint) string {

	var s string
	for _, v := range numbers {
		s += strconv.FormatUint(uint64(v), 10)
	}
	return fmt.Sprintf("(%s) %s-%s", s[:3], s[3:6], s[6:])

}

//

func CreatePhoneNumber229(numbers [10]uint) string {
	group1 := fmt.Sprintf("%d%d%d", numbers[0], numbers[1], numbers[2])
	group2 := fmt.Sprintf("%d%d%d", numbers[3], numbers[4], numbers[5])
	group3 := fmt.Sprintf("%d%d%d%d", numbers[6], numbers[7], numbers[8], numbers[9])

	return fmt.Sprintf("(%s) %s-%s", group1, group2, group3)
}

//

func CreatePhoneNumber230(numbers [10]uint) string {
	result := ""

	for i, v := range numbers {
		strNum := strconv.Itoa(int(v))
		switch i {
		case 0:
			result += "(" + strNum
		case 2:
			result += strNum + ")"
		case 3:
			result += " " + strNum
		case 6:
			result += "-" + strNum
		default:
			result += strNum
		}

	}
	return result
}

//

func CreatePhoneNumber231(numbers [10]uint) string {
	var values string
	for _, i := range numbers {
		values += strconv.FormatUint(uint64(i), 10)
	}
	return fmt.Sprintf("(%v) %v-%v", values[:3], values[3:6], values[6:])

}

//

func CreatePhoneNumber232(numbers [10]uint) string {
	phoneNumber := make([]byte, 14)
	phoneNumber[0] = '('
	phoneNumber[1] = ('0' + byte(numbers[0]))
	phoneNumber[2] = ('0' + byte(numbers[1]))
	phoneNumber[3] = ('0' + byte(numbers[2]))
	phoneNumber[4] = ')'
	phoneNumber[5] = ' '
	phoneNumber[6] = ('0' + byte(numbers[3]))
	phoneNumber[7] = ('0' + byte(numbers[4]))
	phoneNumber[8] = ('0' + byte(numbers[5]))
	phoneNumber[9] = '-'
	phoneNumber[10] = ('0' + byte(numbers[6]))
	phoneNumber[11] = ('0' + byte(numbers[7]))
	phoneNumber[12] = ('0' + byte(numbers[8]))
	phoneNumber[13] = ('0' + byte(numbers[9]))

	return string(phoneNumber)
}

//

func JoinArray(data any, delim string) string {
	s := fmt.Sprint(data)
	f := strings.Fields(s)
	j := strings.Join(f, delim)
	r := strings.Trim(j, "[]")

	return r
}

func CreatePhoneNumber233(numbers [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", JoinArray(numbers[0:3], ""), JoinArray(numbers[3:6], ""), JoinArray(numbers[6:], ""))
}

//

func CreatePhoneNumber234(numbers [10]uint) string {
	var s []string
	for _, n := range numbers {
		s = append(s, strconv.FormatUint(uint64(n), 10))
	}

	code := "(" + s[0] + s[1] + s[2] + ")"
	p1 := s[3] + s[4] + s[5]
	p2 := s[6] + s[7] + s[8] + s[9]

	return code + " " + p1 + "-" + p2
}

//

func CreatePhoneNumber235(numbers [10]uint) string {
	var phone string
	phone = phone + "("
	for i := 0; i < 3; i++ {
		phone = phone + strconv.FormatUint(uint64(numbers[i]), 10)
	}
	phone = phone + ") "
	for i := 3; i < 6; i++ {
		phone = phone + strconv.FormatUint(uint64(numbers[i]), 10)
	}
	phone = phone + "-"
	for i := 6; i < 10; i++ {
		phone = phone + strconv.FormatUint(uint64(numbers[i]), 10)
	}
	return phone
}

//

func CreatePhoneNumber236(numbers [10]uint) string {
	b := strings.Trim(strings.Join(strings.Split(fmt.Sprint(numbers), " "), ""), "[]")
	return fmt.Sprintf("(%s) %s-%s", b[0:3], b[3:6], b[6:])
}

//

func CreatePhoneNumber237(numbers [10]uint) string {
	values := make([]interface{}, len(numbers))

	for i := range numbers {
		values[i] = numbers[i]
	}

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", values...)
}

//

func CreatePhoneNumber238(numbers [10]uint) string {
	phoneNumber := ""
	for _, num := range numbers {
		phoneNumber += strconv.Itoa(int(num))
	}

	formattedPhoneNumber := fmt.Sprintf("(%s) %s-%s",
		phoneNumber[:3],
		phoneNumber[3:6],
		phoneNumber[6:],
	)

	return formattedPhoneNumber
}

//

func CreatePhoneNumber239(numbers [10]uint) string {
	ph_num := "("
	for i := 0; i < len(numbers); i++ {
		if i == 3 {
			ph_num += ") "
		}
		if i == 6 {
			ph_num += "-"
		}
		ph_num += string(rune(numbers[i]) + '0')
	}
	return ph_num

}

//

func CreatePhoneNumber240(numbers [10]uint) string {
	result := ""
	str := fmt.Sprintf("%d", numbers)
	for i, v := range str {
		switch i {
		case 0:
			result += "("
		case 1:
			result += string(v)
		case 3:
			result += string(v)
		case 5:
			result += string(v)
		case 6:
			result += ") "
		case 7:
			result += string(v)
		case 9:
			result += string(v)
		case 11:
			result += string(v)
		case 12:
			result += "-"
		case 13:
			result += string(v)
		case 15:
			result += string(v)
		case 17:
			result += string(v)
		case 19:
			result += string(v)
		}
	}

	return result
}

//

func CreatePhoneNumber241(numbers [10]uint) (ans string) {
	ans = "("
	for i := 0; i < len(numbers); i++ {
		ans += fmt.Sprintf("%d", numbers[i])
		if i == 2 {
			ans += ") "
		} else if i == 5 {
			ans += "-"
		}
	}
	return
}

//

func CreatePhoneNumber242(numbers [10]uint) string {
	// Convert numbers to strings
	numStrings := make([]string, len(numbers))
	for i, num := range numbers {
		numStrings[i] = strconv.FormatUint(uint64(num), 10)
	}

	// Construct the phone number string
	phoneNumber := "(" + numStrings[0] + numStrings[1] + numStrings[2] + ") " +
		numStrings[3] + numStrings[4] + numStrings[5] + "-" +
		numStrings[6] + numStrings[7] + numStrings[8] + numStrings[9]

	return phoneNumber
}

//

func arr2str(arr []uint) string {
	string := ""
	for _, v := range arr {
		string += strconv.FormatUint(uint64(v), 10)
	}
	return string
}

func CreatePhoneNumber243(numbers [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", arr2str(numbers[:3]), arr2str(numbers[3:6]), arr2str(numbers[6:]))
}

//

func CreatePhoneNumber244(numbers [10]uint) string {
	numbers_strings := make([]string, len(numbers))
	for i, v := range numbers {
		numbers_strings[i] = fmt.Sprint(v)
	}
	return fmt.Sprintf(
		"(%s) %s-%s",
		strings.Join(numbers_strings[0:3], ""),
		strings.Join(numbers_strings[3:6], ""),
		strings.Join(numbers_strings[6:], ""),
	)
}

//

func CreatePhoneNumber245(numbers [10]uint) string {
	result := ""
	for i, val := range numbers {
		switch i {
		case 0:
			result += fmt.Sprintf("(%d", val)
		case 2:
			result += fmt.Sprintf("%d) ", val)
		case 6:
			result += fmt.Sprintf("-%d", val)
		default:
			result += fmt.Sprintf("%d", val)
		}
	}
	return result
}

//

func CreatePhoneNumber246(numbers [10]uint) string {
	var str string
	for i, val := range numbers {
		if i == 0 {
			str += "("
		} else if i == 3 {
			str += ") "
		} else if i == 6 {
			str += "-"
		}
		str += strconv.Itoa(int(val))
	}
	return str
}

//

func CreatePhoneNumber247(numbers [10]uint) string {
	var out string
	out += "("
	for i := 0; i < 3; i++ {
		out += strconv.Itoa(int(numbers[i]))
	}
	out += ") "
	for i := 3; i < 6; i++ {
		out += strconv.Itoa(int(numbers[i]))
	}
	out += "-"
	for i := 6; i < 10; i++ {
		out += strconv.Itoa(int(numbers[i]))
	}
	return out
}

//

func CreatePhoneNumber248(numbers [10]uint) string {
	return fmt.Sprintf("(%01d%01d%01d) %01d%01d%01d-%01d%01d%01d%01d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

//

func CreatePhoneNumber249(numbers [10]uint) string {
	var phoneNumber string = "("
	for i := 0; i < 10; i++ {
		if i == 2 {
			phoneNumber += fmt.Sprint(numbers[i]) + ") "
		} else if i == 5 {
			phoneNumber += fmt.Sprint(numbers[i]) + "-"
		} else {
			phoneNumber += fmt.Sprint(numbers[i])
		}
	}
	return phoneNumber
}

//

func CreatePhoneNumber250(numbers [10]uint) string {
	phone := "("

	for i, v := range numbers {
		newValue := strconv.Itoa(int(v))
		if i == 2 {
			newValue += ") "
		} else if i == 5 {
			newValue += "-"
		} else {
			newValue += ""
		}
		phone += newValue
	}

	return phone
}

func CreatePhoneNumber251(numbers [10]uint) string {
	result := "("
	for k, v := range numbers {
		num := strconv.FormatUint(uint64(v), 10)
		result += num
		if k == 2 {
			result += ") "
		}
		if k == 5 {
			result += "-"
		}
	}
	return result
}

//

func CreatePhoneNumber252(numbers [10]uint) string {
	var firstPart string = fmt.Sprintf("(%v%v%v)", numbers[0], numbers[1], numbers[2])
	var secondPart string = fmt.Sprintf(" %v%v%v-%v%v%v%v", numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	return fmt.Sprintf("%v%v", firstPart, secondPart)
}

//

func CreatePhoneNumber253(numbers [10]uint) string {
	builder := strings.Builder{}
	builder.WriteString("(")
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[0]))))
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[1]))))
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[2]))))
	builder.WriteString(") ")
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[3]))))
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[4]))))
	builder.WriteString(fmt.Sprintf("%s-", strconv.Itoa(int(numbers[5]))))
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[6]))))
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[7]))))
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[8]))))
	builder.WriteString(fmt.Sprintf("%s", strconv.Itoa(int(numbers[9]))))
	phone := builder.String()
	return phone
}

//

func CreatePhoneNumber254(numbers [10]uint) string {
	fmt.Println(numbers)
	var myArr [10]string

	for i := 0; i < 10; i++ {
		myArr[i] = fmt.Sprint(numbers[i])
	}
	// myArr[9] = "0"
	var firstPart string = fmt.Sprintf("(%s%s%s)", myArr[0], myArr[1], myArr[2])
	var secondPart string = fmt.Sprintf(" %s%s%s-%s%s%s%s", myArr[3], myArr[4], myArr[5], myArr[6], myArr[7], myArr[8], myArr[9])
	return fmt.Sprintf("%s%s", firstPart, secondPart)
}

//

func CreatePhoneNumber255(numbers [10]uint) string {
	var firstPart string = strconv.Itoa(int(numbers[0])) + strconv.Itoa(int(numbers[1])) + strconv.Itoa(int(numbers[2]))
	var secondPart string = strconv.Itoa(int(numbers[3])) + strconv.Itoa(int(numbers[4])) + strconv.Itoa(int(numbers[5]))
	var thirdPart string = strconv.Itoa(int(numbers[6])) +
		strconv.Itoa(int(numbers[7])) +
		strconv.Itoa(int(numbers[8])) +
		strconv.Itoa(int(numbers[9]))
	var result string = fmt.Sprintf("(%s) %s-%s", firstPart, secondPart, thirdPart)
	return result
}

//

func CreatePhoneNumber256(numbers [10]uint) string {
	areaCode := strings.Trim(strings.Replace(fmt.Sprint(numbers[:3]), " ", "", -1), "[]")
	telP1 := strings.Trim(strings.Replace(fmt.Sprint(numbers[3:6]), " ", "", -1), "[]")
	telP2 := strings.Trim(strings.Replace(fmt.Sprint(numbers[6:]), " ", "", -1), "[]")
	return "(" + areaCode + ")" + " " + telP1 + "-" + telP2
}

//

func CreatePhoneNumber257(numbers [10]uint) string {
	res := ""
	res += string('(')
	for i := 0; i < 3; i++ {
		asciiSymbol := strconv.Itoa(int(numbers[i]))
		res += asciiSymbol
	}
	res += ") "
	for i := 3; i < 6; i++ {
		asciiSymbol := strconv.Itoa(int(numbers[i]))
		res += asciiSymbol
	}
	res += "-"
	for i := 6; i < 10; i++ {
		asciiSymbol := strconv.Itoa(int(numbers[i]))
		res += asciiSymbol
	}
	return res
}

//

func CreatePhoneNumber258(numbers [10]uint) string {

	codeArea := ""
	for _, v := range numbers[0:3] {
		codeArea += fmt.Sprintf("%d", v)
	}

	prefix := ""

	for _, v := range numbers[3:6] {
		prefix += fmt.Sprintf("%d", v)
	}

	suffix := ""

	for _, v := range numbers[6:] {
		suffix += fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("(%s) %s-%s", codeArea, prefix, suffix)
}

//

func CreatePhoneNumber259(numbers [10]uint) string {
	var str string
	for i, num := range numbers {
		if i == 0 {
			str += "(" + strconv.Itoa(int(num))
		} else if i == 2 {
			str += strconv.Itoa(int(num)) + ") "
		} else if i == 5 {
			str += strconv.Itoa(int(num)) + "-"
		} else {
			str += strconv.Itoa(int(num))
		}
	}
	return str
}

//

var phoneNumFmt = "(%d%d%d) %d%d%d-%d%d%d%d"

func CreatePhoneNumber260(numbers [10]uint) string {
	vals := make([]any, len(numbers))
	for i := range vals {
		vals[i] = numbers[i]
	}
	return fmt.Sprintf(phoneNumFmt, vals...)
}

//

func CreatePhoneNumber261(numbers [10]uint) string {

	res := "("

	for i := range numbers {
		if i < 3 {
			res += fmt.Sprintf("%d", numbers[i])
		} else if i == 3 {
			res += ") "
		}
		if i < 6 && i >= 3 {
			res += fmt.Sprintf("%d", numbers[i])
		} else if i == 6 {
			res += "-"
		}
		if i >= 6 {
			res += fmt.Sprintf("%d", numbers[i])
		}
	}
	return res
}

//

func CreatePhoneNumber262(numbers [10]uint) string {
	var Bank []string
	for i := 0; i < len(numbers); i++ {
		dwarf := numbers[i]
		Bank = append(Bank, strconv.FormatUint(uint64(dwarf), 10))
	}
	return "(" + Bank[0] + Bank[1] + Bank[2] + ")" + " " + Bank[3] + Bank[4] + Bank[5] + "-" + Bank[6] + Bank[7] + Bank[8] + Bank[9]
}

//

func CreatePhoneNumber263(numbers [10]uint) string {
	strVals := convertToString(numbers)
	return fmt.Sprintf("(%s) %s-%s", strVals[0:3], strVals[3:6], strVals[6:10])
}

func convertToString(numbers [10]uint) string {
	var result string
	for _, value := range numbers {
		result += strconv.FormatUint(uint64(value), 10)
	}
	return result
}

//

func CreatePhoneNumber264(numbers [10]uint) string {
	numberString := "("
	for idx, val := range numbers {
		strVal := strconv.FormatUint(uint64(val), 10)
		switch idx {
		case 2:
			numberString += strVal + ") "
			break
		case 5:
			numberString += strVal + "-"
			break
		default:
			numberString += strVal
			break
		}
	}
	return numberString
}

//

func CreatePhoneNumber265(numbers [10]uint) string {
	phoneNumber := "("
	for i, n := range numbers {
		ns := strconv.FormatUint(uint64(n), 10)
		if i <= 2 {
			phoneNumber = phoneNumber + ns
		}
		if i == 3 {
			phoneNumber = phoneNumber + ")" + " " + ns
		}
		if i >= 4 && i <= 5 {
			phoneNumber = phoneNumber + ns
		}
		if i == 6 {
			phoneNumber = phoneNumber + "-" + ns
		}
		if i >= 7 {
			phoneNumber = phoneNumber + ns
		}
	}

	return phoneNumber
}

//

func CreatePhoneNumber266(n [10]uint) string {
	var sn []string
	for i := 0; i < 10; i++ {
		sn = append(sn, strconv.Itoa(int(n[i])))
	}
	return fmt.Sprintf("(%s) %s-%s",
		sn[0]+sn[1]+sn[2],
		sn[3]+sn[4]+sn[5],
		sn[6]+sn[7]+sn[8]+sn[9])
}

//

func CreatePhoneNumber267(numbers [10]uint) string {
	sliceStrings := make([]string, len(numbers))
	for index, value := range numbers {
		sliceStrings[index] = strconv.FormatUint(uint64(value), 10)
	}
	str := fmt.Sprintf("(%s) %s-%s", strings.Join(sliceStrings[:3], ""), strings.Join(sliceStrings[3:6], ""), strings.Join(sliceStrings[6:10], ""))
	return str
}

//

func CreatePhoneNumber268(numbers [10]uint) string {
	prefix := Join(numbers[0:3])
	first_part := Join(numbers[3:6])
	second_part := Join(numbers[6:10])
	return fmt.Sprintf("(%s) %s-%s", prefix, first_part, second_part)
}

func Join(numbers []uint) string {
	var ret string

	for _, val := range numbers {
		ret += fmt.Sprintf("%d", val)
	}
	return ret
}

//

func CreatePhoneNumber269(numbers [10]uint) string {
	var phoneNumber = "("
	for i, digit := range numbers {
		if i == 2 {
			phoneNumber += fmt.Sprintf("%d) ", digit)
		} else if i == 5 {
			phoneNumber += fmt.Sprintf("%d-", digit)
		} else {
			phoneNumber += strconv.Itoa(int(digit))
		}
	}
	return phoneNumber
}

//

func CreatePhoneNumber270(numbers [10]uint) (res string) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			res += "("
		case 3:
			res += ") "
		case 6:
			res += "-"
		}
		res += strconv.Itoa(int(numbers[i]))
	}
	return res
}

//

func CreatePhoneNumber271(numbers [10]uint) string {
	myStr := "("
	for i := 0; i < len(numbers); i++ {
		myStr = myStr + strconv.Itoa(int(numbers[i]))
		if i == 2 {
			myStr = myStr + ")"
		}
		if i == 2 {
			myStr = myStr + " "
		}
		if i == 5 {
			myStr = myStr + "-"
		}
	}
	return myStr
}

//

func CreatePhoneNumber272(numbers [10]uint) string {
	first := numbers[0:3]
	second := numbers[3:6]
	third := numbers[6:]

	phoneString := fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", first[0], first[1], first[2], second[0], second[1], second[2], third[0], third[1], third[2], third[3])

	return phoneString
}

//

func CreatePhoneNumber273(numbers [10]uint) string {
	acc := ""
	for _, v := range numbers {
		acc += fmt.Sprint(v)
	}
	r, _ := regexp.Compile(`(\d{3})(\d{3})(\d{4})`)
	m := r.FindStringSubmatch(acc)
	return fmt.Sprintf("(%s) %s-%s", m[1], m[2], m[3])
}

//

func CreatePhoneNumber274(numbers [10]uint) string {
	return fmt.Sprintf("(%s) %s-%s", uint2str(numbers[:3]), uint2str(numbers[3:6]), uint2str(numbers[6:]))
}

func uint2str(n []uint) string {
	res := ""
	for _, v := range n {
		res += fmt.Sprint(v)
	}
	return res
}

//

func CreatePhoneNumber275(numbers [10]uint) string {
	result := ""
	for index, value := range numbers {
		if index == 0 {
			result += "("
		} else if index == 3 {
			result += ") "
		} else if index == 6 {
			result += "-"
		}
		result += strconv.Itoa(int(value))
	}
	return result
}

//

func CreatePhoneNumber276(numbers [10]uint) string {
	var b strings.Builder
	for _, n := range numbers {
		fmt.Fprintf(&b, "%s", fmt.Sprint(n))
	}
	str := b.String()

	return fmt.Sprintf("(%s) %s-%s", str[:3], str[3:6], str[6:])
}

//

func CreatePhoneNumber277(numbers [10]uint) string {

	var result strings.Builder

	first := fmtNum(numbers[0:3])
	result.WriteString("(")
	result.WriteString(first)
	result.WriteString(")")
	result.WriteString(" ")

	second := fmtNum(numbers[3:6])
	result.WriteString(second)
	result.WriteString("-")
	result.WriteString(fmtNum(numbers[6:]))
	return result.String()
}
func fmtNum(a []uint) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(int(v))
		if unicode.IsNumber(rune(v)) {
		}
	}

	return strings.Join(b, "")
}

//

func CreatePhoneNumber278(numbers [10]uint) string {
	var phoneNumber string
	for i, num := range numbers {
		switch i {
		case 0:
			phoneNumber += "(" + fmt.Sprint(num)
		case 2:
			phoneNumber += fmt.Sprint(num) + ")" + " "
		case 5:
			phoneNumber += fmt.Sprint(num) + "-"
		default:
			phoneNumber += fmt.Sprint(num)
		}
	}
	return phoneNumber
}

//

func CreatePhoneNumber279(numbers [10]uint) string {
	res := "("
	var digit string
	for i, val := range numbers {
		if i == 3 {
			res += ") "
		} else if i == 6 {
			res += "-"
		}
		digit = strconv.FormatUint(uint64(val), 10)
		res += digit
	}
	return res
}

//

func CreatePhoneNumber280(numbers [10]uint) string {
	var mask = "(xxx) xxx-xxxx"

	for i := 0; i < len(numbers); i++ {
		mask = strings.Replace(mask, "x", strconv.Itoa(int(numbers[i])), 1)
	}

	return mask
}

//

func CreatePhoneNumber281(numbers [10]uint) string {
	var phoneNumber string
	for i := 0; i < len(numbers); i++ {
		num := numbers[i]

		switch i {
		case 0:
			phoneNumber += fmt.Sprintf("(%d", num)
		case 2:
			phoneNumber += fmt.Sprintf("%d) ", num)
		case 5:
			phoneNumber += fmt.Sprintf("%d-", num)
		default:
			phoneNumber += fmt.Sprintf("%d", num)
		}
	}

	return phoneNumber
}

//

func CreatePhoneNumber282(numbers [10]uint) string {
	s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ""), "[]")

	re := regexp.MustCompile(`([\d]{3})([\d]{3})([\d]{4})`)
	s = re.ReplaceAllString(s, "($1) $2-$3")
	return s
}

//

func CreatePhoneNumber283(numbers [10]uint) string {
	numb := "("
	for i := 0; i < 3; i++ {
		numb += strconv.Itoa(int(numbers[i]))
	}
	numb += ") "
	for j := 3; j < len(numbers); j++ {
		numb += strconv.Itoa(int(numbers[j]))
		if j == 5 {
			numb += "-"
		}
	}
	return numb
}

//

func CreatePhoneNumber284(numbers [10]uint) string {
	result := "("
	for i, num := range numbers {
		result += strconv.Itoa(int(num))
		if i == 2 {
			result += ") "
		} else if i == 5 {
			result += "-"
		}
	}
	return result
}

//

func CreatePhoneNumber285(numbers [10]uint) string {
	result := "("
	for i := 0; i < len(numbers); i++ {
		if i == 3 {
			result += ") "
		} else if i == 6 {
			result += "-"
		}
		result += strconv.Itoa(int(numbers[i]))
	}
	return result
}

//

func CreatePhoneNumber286(numbers [10]uint) string {
	numInterface := make([]interface{}, len(numbers))
	for i, num := range numbers {
		numInterface[i] = num
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numInterface...)
}

//

func CreatePhoneNumber287(numbers [10]uint) string {
	args := []interface{}{}
	for _, i := range numbers {
		args = append(args, i)
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", args...)
}

//

func toInterfaceSlice(numbers [10]uint) []interface{} {
	s := make([]interface{}, len(numbers))
	for i, v := range numbers {
		s[i] = v
	}
	return s
}

func CreatePhoneNumber288(numbers [10]uint) string {
	args := toInterfaceSlice(numbers)
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", args...)
}

//

func CreatePhoneNumber289(numbers [10]uint) string {
	var a string = fmt.Sprintf("%d%d%d", numbers[0], numbers[1], numbers[2])
	var b string = fmt.Sprintf("%d%d%d", numbers[3], numbers[4], numbers[5])
	var c string = fmt.Sprintf("%d%d%d%d", numbers[6], numbers[7], numbers[8], numbers[9])
	var res string = "(" + a + ")" + " " + b + "-" + c
	return res
}

//

func CreatePhoneNumber290(numbers [10]uint) string {
	var str string

	for _, i := range numbers {
		str += strconv.FormatUint(uint64(i), 10)
	}

	return fmt.Sprintf("(%v) %v-%v", str[:3], str[3:6], str[6:])
}

//

func CreatePhoneNumber291(numbers [10]uint) string {
	var l1 = []string{"("}
	for x := 0; x < 3; x++ {
		l1 = append(l1, strconv.FormatInt(int64(numbers[x]), 10))
	}
	l1 = append(l1, ") ")

	var l2 = []string{}
	for x := 3; x < 6; x++ {
		l2 = append(l2, strconv.FormatInt(int64(numbers[x]), 10))
	}
	l2 = append(l2, "-")

	var l3 = []string{}
	for x := 6; x < 10; x++ {
		l3 = append(l3, strconv.FormatInt(int64(numbers[x]), 10))
	}

	l1s := strings.Join(l1, "")
	l2s := strings.Join(l2, "")
	l3s := strings.Join(l3, "")
	return l1s + l2s + l3s
}

//

func CreatePhoneNumber292(numbers [10]uint) string {
	var result string = fmt.Sprintf("(%s) %s-%s", joinArrIntoStrings(numbers[0:3]), joinArrIntoStrings(numbers[3:6]), joinArrIntoStrings(numbers[6:10]))
	return result
}

func joinArrIntoStrings(numbers []uint) string {
	newList := make([]string, len(numbers))
	for i, numberUnity := range numbers {
		newList[i] = strconv.FormatUint(uint64(numberUnity), 10)
	}
	return strings.Join(newList, "")
}

//

func CreatePhoneNumber293(numbers [10]uint) string {
	var numberStr string
	for idx, number := range numbers {
		switch idx {
		case 0:
			numberStr += `(`
		case 3:
			numberStr += `) `
		case 6:
			numberStr += `-`
		}
		numberStr += strconv.FormatUint(uint64(number), 10)
	}
	return numberStr
}

//

func convertNumberArrayToString(numbers []uint) string {
	return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

func CreatePhoneNumber294(numbers [10]uint) string {
	first3 := convertNumberArrayToString(numbers[:3])
	second3 := convertNumberArrayToString(numbers[3:6])
	last4 := convertNumberArrayToString(numbers[6:])

	return fmt.Sprintf("(%s) %s-%s", first3, second3, last4)
}

//

func CreatePhoneNumber295(numbers [10]uint) string {
	var s string
	for _, v := range numbers {
		s += strconv.Itoa(int(v))
	}
	return fmt.Sprintf("(%s) %s-%s", s[0:3], s[3:6], s[6:10])
}

//

func CreatePhoneNumber2966(numbers [10]uint) string {
	ans := "("
	for i := 0; i < 3; i++ {
		ans += strconv.FormatUint(uint64(numbers[i]), 10)
	}
	ans += ") "
	for i := 3; i < 6; i++ {
		ans += strconv.FormatUint(uint64(numbers[i]), 10)
	}
	ans += "-"
	for i := 6; i < 10; i++ {
		ans += strconv.FormatUint(uint64(numbers[i]), 10)
	}

	return ans
}

//

func CreatePhoneNumber296(numbers [10]uint) string {
	first, second, third := numbers[:3], numbers[3:6], numbers[6:]
	return fmt.Sprintf("(%s) %s-%s", arrToStr(first), arrToStr(second), arrToStr(third))
}

func arrToStr(arr []uint) (str string) {
	for _, v := range arr {
		str += strconv.Itoa(int(v))
	}
	return
}

//

func CreatePhoneNumber297(numbers [10]uint) string {

	nums := make([]string, len(numbers))
	for i, v := range numbers {
		nums[i] = fmt.Sprint(v)
	}
	return fmt.Sprintf("(%s) %s-%s", strings.Join(nums[:3], ""), strings.Join(nums[3:6], ""), strings.Join(nums[6:], ""))

}

//

func CreatePhoneNumber298(numbers [10]uint) string {
	var ddd, phone1, phone2 string

	for i := 0; i < len(numbers); i++ {
		// Converte `uint` para `int` antes de chamar `strconv.Itoa`
		if i < 3 {
			ddd += strconv.Itoa(int(numbers[i]))
		} else if i >= 3 && i <= 5 {
			phone1 += strconv.Itoa(int(numbers[i]))
		} else {
			phone2 += strconv.Itoa(int(numbers[i]))
		}
	}

	return "(" + ddd + ") " + phone1 + "-" + phone2
}

//

func CreatePhoneNumber299(numbers [10]uint) string {
	var phoneNum string
	for i, num := range numbers {
		if i == 0 {
			phoneNum = "(" + phoneNum
		} else if i == 3 {
			phoneNum = phoneNum + ") "
		} else if i == 6 {
			phoneNum = phoneNum + "-"
		}
		phoneNum = phoneNum + fmt.Sprint(num)
	}

	return phoneNum

}

//

func CreatePhoneNumber300(numbers [10]uint) string {
	var phoneNumber string
	for idx, number := range numbers {
		switch idx {
		case 0:
			phoneNumber += "("
		case 3:
			phoneNumber += ") "
		case 6:
			phoneNumber += "-"
		}

		numberStr := strconv.Itoa(int(number))
		phoneNumber += numberStr

	}
	return phoneNumber
}

//

func CreatePhoneNumber301(numbers [10]uint) string {
	strNum := strings.Replace(
		fmt.Sprint(numbers),
		" ", "", -1)
	strNum = strings.Trim(strNum, "[]")

	return fmt.Sprintf("(%s) %s-%s", strNum[0:3], strNum[3:6], strNum[6:])
}

//

func CreatePhoneNumber302(numbers [10]uint) string {
	strNum := ""
	for i := 0; i < len(numbers); i++ {
		strNum += strconv.FormatUint(uint64(numbers[i]), 10)
	}

	return fmt.Sprintf("(%s) %s-%s", strNum[0:3], strNum[3:6], strNum[6:10])
}

//

func getString(arr []uint) string {
	sb := strings.Builder{}
	for _, num := range arr {
		sb.WriteString(strconv.Itoa(int(num)))
	}

	return sb.String()
}

func CreatePhoneNumber303(numbers [10]uint) string {
	code := getString(numbers[:3])
	left := getString(numbers[3:6])
	right := getString(numbers[6:])

	return fmt.Sprintf("(%s) %s-%s", code, left, right)
}

//

func CreatePhoneNumber304(numbers [10]uint) string {
	var phoneNumber string
	for i, num := range numbers {
		switch i {
		case 0:
			phoneNumber += fmt.Sprintf("(%d", num)
		case 2:
			phoneNumber += fmt.Sprintf("%d)", num)
		case 3:
			phoneNumber += fmt.Sprintf(" %d", num)
		case 6:
			phoneNumber += fmt.Sprintf("-%d", num)
		default:
			phoneNumber += fmt.Sprintf("%d", num)
		}
	}
	return phoneNumber
}

//

func CreatePhoneNumber305(numbers [10]uint) string {

	var s []string

	for i, v := range numbers {
		var i2 int = int(v)
		switch i {
		case 0:
			s = append(s, "(")
		case 3:
			s = append(s, ") ")
		case 6:
			s = append(s, "-")
		}
		s = append(s, strconv.Itoa(i2))
	}

	return strings.Join(s, "")
}

//

func CreatePhoneNumber306(numbers [10]uint) string {

	var phone = []string{}
	phone = append(phone, "(")
	var index = numbers[:3]
	fmt.Println(index)

	for _, i := range index {
		j := strconv.FormatUint(uint64(i), 16)
		phone = append(phone, j)
	}
	phone = append(phone, ") ")

	var operator = numbers[3:6]
	for _, i := range operator {
		j := strconv.FormatUint(uint64(i), 16)
		phone = append(phone, j)
	}

	phone = append(phone, "-")

	var number = numbers[6:]
	for _, i := range number {
		j := strconv.FormatUint(uint64(i), 16)
		phone = append(phone, j)
	}

	return strings.Join(phone, "")
}

//

func CreatePhoneNumber307(numbers [10]uint) string {
	var s []string
	for _, v := range numbers {
		s = append(s, strconv.Itoa(int(v)))
	}

	result := "(" + strings.Join(s[0:3], "") + ")" + " " + strings.Join(s[3:6], "") + "-" + strings.Join(s[6:], "")
	return result
}

//

func CreatePhoneNumber308(numbers [10]uint) string {
	tel := []string{}
	for _, v := range numbers {
		x := strconv.FormatUint(uint64(v), 10)
		tel = append(tel, x)
	}
	res := "(" + tel[0] + tel[1] + tel[2] + ") " + tel[3] + tel[4] + tel[5] + "-" + tel[6] + tel[7] + tel[8] + tel[9]
	return res

}

//

func CreatePhoneNumber309(numbers [10]uint) string {
	var phoneNumber string
	phoneNumber += "("

	for i := 0; i < 3; i++ {
		phoneNumber += strconv.Itoa(int(numbers[i]))
	}
	phoneNumber += ") "
	for i := 3; i < 6; i++ {
		phoneNumber += strconv.Itoa(int(numbers[i]))
	}
	phoneNumber += "-"
	for i := 6; i < 10; i++ {
		phoneNumber += strconv.Itoa(int(numbers[i]))
	}

	return phoneNumber
}

//

func CreatePhoneNumber310(numbers [10]uint) string {
	str := ArrayToString4(numbers)
	return fmt.Sprintf("(%s) %s-%s", str[0:3], str[3:6], str[6:])
}

func ArrayToString4(numbers interface{}) string {
	return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

//

func CreatePhoneNumber311(numbers [10]uint) string {
	var numberStrs []string
	for _, i := range numbers {
		numberStrs = append(numberStrs, strconv.Itoa(int(i)))
	}
	res := "(" + strings.Join(numberStrs[0:3][:], "") + ") " +
		strings.Join(numberStrs[3:6][:], "") + "-" +
		strings.Join(numberStrs[6:][:], "")
	return res
}

//

func CreatePhoneNumber312(numbers [10]uint) string {
	seperated := strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
	return fmt.Sprintf("(%s) %s-%s", seperated[0:3], seperated[3:6], seperated[6:10])
}

//

func CreatePhoneNumber313(numbers [10]uint) string {
	var str1 strings.Builder
	var str2 strings.Builder

	for i, v := range numbers {
		if i <= 2 {
			str1.WriteString(fmt.Sprintf("%d", v))
		} else if i == 6 {
			str2.WriteString(fmt.Sprintf("-%d", v))
		} else {
			str2.WriteString(fmt.Sprintf("%d", v))
		}
	}

	return fmt.Sprintf("(%s) %s", str1.String(), str2.String())
}

//

func CreatePhoneNumber314(nums [10]uint) string {
	var arrToStr = func(input []uint) (result string) {
		for _, val := range input {
			result += fmt.Sprintf("%v", val)
		}
		return
	}
	return fmt.Sprintf("(%v) %v-%v", arrToStr(nums[:3]), arrToStr(nums[3:][:3]), arrToStr(nums[6:]))
}

//

func CreatePhoneNumber315(numbers [10]uint) string {
	var numstr []string
	var str []string
	var run rune
	var phone string
	for i := 0; i <= len(numbers)-1; i++ {
		run = int32(numbers[i] + '0')
		numstr = append(numstr, string(run))

	}
	for j := 0; j <= len(numstr)-1; j++ {
		if j == 0 {
			str = append(str, ("(" + numstr[j]))
		} else if j == 2 {
			str = append(str, (numstr[j] + ")"))
		} else if j == 5 {
			str = append(str, (numstr[j] + "-"))
		} else if j == 3 {
			str = append(str, (" " + numstr[j]))
		} else {
			str = append(str, numstr[j])
		}
	}
	for _, char := range str {
		phone += char
	}
	return phone

}

//

func CreatePhoneNumber316(numbers [10]uint) string {
	b := []byte{}
	for i := 0; i < len(numbers); i++ {
		b = append(b, byte(numbers[i]+'0'))
	}
	return "(" + string(b[:3]) + ") " + string(b[3:6]) + "-" + string(b[6:])
}

//

func CreatePhoneNumber317(n [10]uint) string {
	pattern := []rune("(...) ...-....")
	u := 0
	for k, c := range pattern {
		if c == rune('.') {
			pattern[k] = rune(n[u] + '0')
			u++
		}
	}
	return string(pattern)
}

//

func CreatePhoneNumber318(numbers [10]uint) string {
	var numbersStr [10]string
	for i := 0; i < 10; i++ {
		numbersStr[i] = fmt.Sprintf("%d", numbers[i])
	}

	return fmt.Sprintf("(%s) %s-%s", strings.Join(numbersStr[:3], ""), strings.Join(numbersStr[3:6], ""), strings.Join(numbersStr[6:10], ""))
}

//

func CreatePhoneNumber319(numbers [10]uint) string {
	tmp := strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
	return "(" + tmp[:3] + ") " + tmp[3:6] + "-" + tmp[6:]
	//return fmt.Sprintf("(%s) %s-%s", tmp[:3], tmp[3:6], tmp[6:])
}

//

func CreatePhoneNumber320(numbers [10]uint) string {
	s := strings.ReplaceAll(fmt.Sprint(numbers), " ", "")
	return fmt.Sprintf("(%v) %v-%v", s[1:4], s[4:7], s[7:11])
}

//

func CreatePhoneNumber321(numbers [10]uint) string {
	var strSlice []string
	for _, num := range numbers {
		str := fmt.Sprintf("%v", num)
		strSlice = append(strSlice, str)
	}

	return "(" + strings.Join(strSlice[0:3], "") + ")" + " " + strings.Join(strSlice[3:6], "") + "-" + strings.Join(strSlice[6:], "")
}

//

func CreatePhoneNumber322(numbers [10]uint) string {
	var ft string
	for _, v := range numbers[0:3] {
		ft += fmt.Sprintf("%d", v)
	}
	var mt string
	for _, v := range numbers[3:6] {
		mt += fmt.Sprintf("%d", v)
	}
	var lt string
	for _, v := range numbers[6:] {
		lt += fmt.Sprintf("%d", v)
	}
	return fmt.Sprintf("(%s) %s-%s", ft, mt, lt)
}

func CreatePhoneNumber323(numbers [10]uint) string {
	return "(" + string([]byte{'0' + byte(numbers[0]), '0' + byte(numbers[1]), '0' + byte(numbers[2])}) + ") " + string([]byte{'0' + byte(numbers[3]), '0' + byte(numbers[4]), '0' + byte(numbers[5])}) + "-" + string([]byte{'0' + byte(numbers[6]), '0' + byte(numbers[7]), '0' + byte(numbers[8]), '0' + byte(numbers[9])})
}

//

func CreatePhoneNumber324(numbers [10]uint) string {
	r := regexp.MustCompile(`\s+`)
	s := r.ReplaceAllString(fmt.Sprintf("%v", numbers), "")

	return fmt.Sprintf("(%s) %s-%s", s[1:4], s[4:7], s[7:11])
}

//

func CreatePhoneNumber325(numbers [10]uint) string {
	var str string
	var arr []string
	for _, i := range numbers {
		arr = append(arr, fmt.Sprint(i))
	}
	a := strings.Join(arr[:3], "")
	b := strings.Join(arr[3:6], "")
	c := strings.Join(arr[6:], "")
	str = fmt.Sprintf("(%s) %s-%s", a, b, c)
	return str
}

//

func CreatePhoneNumber326(numbers [10]uint) string {
	tmpl := "(%s) %s-%s"
	var aarr []string
	var barr []string
	var carr []string
	for i, n := range numbers {
		v := strconv.Itoa(int(n))
		if i < 3 {
			aarr = append(aarr, v)
		} else if i < 6 {
			barr = append(barr, v)
		} else {
			carr = append(carr, v)
		}

	}
	a := strings.Join(aarr, "")
	b := strings.Join(barr, "")
	c := strings.Join(carr, "")
	s := fmt.Sprintf(tmpl, a, b, c)
	return s
}

//

func CreatePhoneNumber327(numbers [10]uint) string {
	phoneNumber := "("
	for i, value := range numbers {
		if i < 3 {
			if i == 2 {
				phoneNumber += strconv.Itoa(int(value))
				phoneNumber += ") "
			} else {
				phoneNumber += strconv.Itoa(int(value))
			}
		} else {
			if i == 5 {
				phoneNumber += strconv.Itoa(int(value))
				phoneNumber += "-"
			} else {
				phoneNumber += strconv.Itoa(int(value))
			}
		}
	}
	return phoneNumber
}

func CreatePhoneNumber328(numbers [10]uint) string {
	b := "("
	for i, v := range numbers {
		if i == 3 {
			b = b + ") "
		}
		if i == 6 {
			b = b + "-"
		}

		b = b + strconv.FormatUint(uint64(v), 10)
	}
	return b
}

//

func CreatePhoneNumber329(arr [10]uint) string {
	var str string
	for _, v := range arr {
		str += strconv.Itoa(int(v))
	}

	re, _ := regexp.Compile(`(?P<val1>\d{3})(?P<val2>\d{3})(?P<val3>\d{4})`)
	template := "(${val1}) ${val2}-${val3}"
	fmt.Println(re.ReplaceAllString(str, template))
	return re.ReplaceAllString(str, template)
}

//

func CreatePhoneNumber330(numbers [10]uint) string {
	phoneNumber := "("
	for i := range numbers {
		if i >= 0 && i <= 2 {
			phoneNumber += strconv.Itoa(int(numbers[i]))
			continue
		}

		if i == 3 {
			phoneNumber += ") "
			phoneNumber += strconv.Itoa(int(numbers[i]))
			continue
		}

		if i == 6 {
			phoneNumber += "-"
			phoneNumber += strconv.Itoa(int(numbers[i]))
			continue
		}

		phoneNumber += strconv.Itoa(int(numbers[i]))

	}
	return phoneNumber
}

//

func CreatePhoneNumber331(numbers [10]uint) string {
	mask := "(xxx) xxx-xxxx"

	for _, v := range numbers {
		val := int(v)
		num := strconv.Itoa(val)

		mask = strings.Replace(mask, "x", num, 1)
	}

	return mask
}

//

func CreatePhoneNumber332(numbers [10]uint) string {
	var ans = "("
	for _, v := range numbers[0:3] {
		ans += strconv.FormatInt(int64(v), 10)
	}
	ans += ") "
	for _, v := range numbers[3:6] {
		ans += strconv.FormatInt(int64(v), 10)
	}
	ans += "-"
	for _, v := range numbers[6:] {
		ans += strconv.FormatInt(int64(v), 10)
	}
	return ans
}

//

func CreatePhoneNumber333(numbers [10]uint) string {
	prefix := numbers[0]*100 + numbers[1]*10 + numbers[2]
	first := numbers[3]*100 + numbers[4]*10 + numbers[5]
	second := numbers[6]*1000 + numbers[7]*100 + numbers[8]*10 + numbers[9]

	return fmt.Sprintf("(%.3d) %.3d-%.4d", prefix, first, second)
}

//

func CreatePhoneNumber334(numbers [10]uint) string {
	var areaCode string
	areaCodeLength := 3
	for i := 0; i < areaCodeLength; i++ {
		areaCode += strconv.Itoa(int(numbers[i]))
	}

	var phoneNumberPrefix string
	prefixLength := 3
	prefixStart := areaCodeLength
	prefixEnd := areaCodeLength + prefixLength
	for i := prefixStart; i < prefixEnd; i++ {
		fmt.Print(numbers[i])
		phoneNumberPrefix += strconv.Itoa(int(numbers[i]))
	}

	var phoneNumberSuffix string
	suffixLength := 4
	suffixStart := areaCodeLength + prefixLength
	suffixEnd := areaCodeLength + prefixLength + suffixLength
	for i := suffixStart; i < suffixEnd; i++ {
		fmt.Print(numbers[i])
		phoneNumberSuffix += strconv.Itoa(int(numbers[i]))
	}

	phoneNumber := fmt.Sprintf("(%v) %v-%v", areaCode, phoneNumberPrefix, phoneNumberSuffix)
	return phoneNumber
}

//

func CreatePhoneNumber335(numbers [10]uint) string {
	areaCode := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers[:3])), ""), "[]")
	prefix := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers[3:6])), ""), "[]")
	lineNumber := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers[6:])), ""), "[]")
	return fmt.Sprintf("(%s) %s-%s", areaCode, prefix, lineNumber)
}

//

func CreatePhoneNumber336(numbers [10]uint) string {

	s := "("

	for j := 0; j < 3; j++ {
		jString := strconv.FormatUint(uint64(numbers[j]), 10)

		//s += strconv.Itoa(numbers[j])
		s += jString

	}
	s += ") "

	for k := 3; k < 6; k++ {

		jString := strconv.FormatUint(uint64(numbers[k]), 10)

		//s += strconv.Itoa(numbers[j])
		s += jString

	}

	s += "-"

	for k := 6; k < 10; k++ {

		jString := strconv.FormatUint(uint64(numbers[k]), 10)

		//s += strconv.Itoa(numbers[j])
		s += jString

	}

	return s
}

func CreatePhoneNumber337(numbers [10]uint) string {
	sb := "("

	for i, n := range numbers {
		// String to prepend if any
		p := ""
		switch i {
		case 3:
			p += ") "
		case 6:
			p += "-"
		}
		sb += p + string(rune('0'+n))
	}

	return sb
}

//

func arrayToString5(numbers []uint) string {
	newString := ""

	for _, number := range numbers {
		newString += strconv.FormatUint(uint64(number), 10)
	}

	return newString
}

func CreatePhoneNumber339(numbers [10]uint) string {
	phone_number := "("

	phone_number += arrayToString5(numbers[0:3])
	phone_number += ") "
	phone_number += arrayToString5(numbers[3:6])
	phone_number += "-"
	phone_number += arrayToString5(numbers[6:10])

	return phone_number
}

//

func CreatePhoneNumber340(numbers [10]uint) string {
	var _numbers []any
	for _, n := range numbers {
		_numbers = append(_numbers, n)
	}
	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", _numbers...)
}

//

func CreatePhoneNumber341(numbers [10]uint) string {
	phoneNumber := "("
	phoneNumber += fmt.Sprint(numbers[0])
	phoneNumber += fmt.Sprint(numbers[1])
	phoneNumber += fmt.Sprint(numbers[2])
	phoneNumber += ") "
	phoneNumber += fmt.Sprint(numbers[3])
	phoneNumber += fmt.Sprint(numbers[4])
	phoneNumber += fmt.Sprint(numbers[5])
	phoneNumber += "-"
	phoneNumber += fmt.Sprint(numbers[6])
	phoneNumber += fmt.Sprint(numbers[7])
	phoneNumber += fmt.Sprint(numbers[8])
	phoneNumber += fmt.Sprint(numbers[9])
	return phoneNumber
}

//

var regex = regexp.MustCompile(`[\[\]]`)

func CreatePhoneNumber342(numbers [10]uint) string {

	return regex.ReplaceAllString(
		fmt.Sprintf(
			"(%s) %s-%s",
			strings.Join(strings.Fields(fmt.Sprint(numbers[0:3])), ""),
			strings.Join(strings.Fields(fmt.Sprint(numbers[3:6])), ""),
			strings.Join(strings.Fields(fmt.Sprint(numbers[6:])), ""),
		),
		"",
	)
}

//

func CreatePhoneNumber343(numbers [10]uint) string {
	s := func(i int) string { return strconv.Itoa(int(numbers[i])) }

	return "(" + s(0) + s(1) + s(2) + ")" + " " + s(3) + s(4) + s(5) + "-" + s(6) + s(7) + s(8) + s(9)
}

//

func genNum(numbers [10]uint, l int, r int) (num uint) {
	for i := l; i < r; i++ {
		num = (num * 10) + numbers[i]
	}
	return
}

func CreatePhoneNumber344(numbers [10]uint) string {
	a, b, c := genNum(numbers, 0, 3), genNum(numbers, 3, 6), genNum(numbers, 6, 10)
	return fmt.Sprintf("(%03d) %03d-%04d", a, b, c)
}

//

func CreatePhoneNumber345(numbers [10]uint) string {
	var s [len(numbers)]string
	for i := 0; i < len(numbers); i++ {
		s[i] = strconv.FormatUint(uint64(numbers[i]), 10)
	}
	return "(" + s[0] + s[1] + s[2] + ")" + " " + s[3] + s[4] + s[5] + "-" + s[6] + s[7] + s[8] + s[9]
}

//

func CreatePhoneNumber346(numbers [10]uint) string {
	var s = fmt.Sprintf("(%v)*%v-%v", numbers[0:3], numbers[3:6], numbers[6:])
	replacer := strings.NewReplacer("[", "", "]", "", " ", "", "*", " ")
	return replacer.Replace(s)

}

//

func CreatePhoneNumber347(numbers [10]uint) string {
	area_code := NumArrayToString(numbers[0:3], 3)
	section_1 := NumArrayToString(numbers[3:6], 3)
	section_2 := NumArrayToString(numbers[6:10], 4)

	return fmt.Sprintf("(%s) %s-%s", area_code, section_1, section_2)
}

func NumArrayToString(numbers []uint, length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += strconv.FormatUint(uint64(numbers[i]), 10)
	}
	return result
}

//

func CreatePhoneNumber348(numbers [10]uint) string {
	var first string
	var mid string
	var last string
	for i := 0; i < len(numbers); i++ {
		if i < 3 {
			first += strconv.Itoa(int(numbers[i]))
			continue
		}
		if i >= 3 && i < 6 {
			mid += strconv.Itoa(int(numbers[i]))
			continue
		}
		last += strconv.Itoa(int(numbers[i]))
	}
	sprintf := fmt.Sprintf("(%s) %s-%s", first, mid, last)
	return sprintf
}

//

func CreatePhoneNumber349(numbers [10]uint) (r string) {
	for i, n := range numbers {
		switch i {
		case 0:
			r = "(" + strconv.FormatUint(uint64(n), 10)
		case 2:
			r += strconv.FormatUint(uint64(n), 10) + ")" + " "
		case 5:
			r += strconv.FormatUint(uint64(n), 10) + "-"
		default:
			r += strconv.FormatUint(uint64(n), 10)
		}
	}

	return
}

//

func CreatePhoneNumber350(numbers [10]uint) string {
	phoneNumber := ""

	for i := 0; i < 10; i++ {
		if i == 0 {
			phoneNumber = "(" + strconv.FormatUint(uint64(numbers[i]), 10)
		} else if i == 2 {
			phoneNumber += strconv.FormatUint(uint64(numbers[i]), 10) + ") "
		} else if i == 5 {
			phoneNumber += strconv.FormatUint(uint64(numbers[i]), 10) + "-"
		} else {
			phoneNumber += strconv.FormatUint(uint64(numbers[i]), 10)
		}
		fmt.Println(phoneNumber)
	}
	return phoneNumber

}

//

func CreatePhoneNumber351(numbers [10]uint) string {
	m := map[int]string{
		0: "(",
		3: ") ",
		6: "-",
	}
	var res string

	for i := 0; i < len(numbers); i++ {
		if v, ok := m[i]; ok {
			res += v
		}
		res += strconv.Itoa(int(numbers[i]))
	}

	return res
}

//

func GetNumbersString(numbers []uint) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(numbers), " "), ""), "[]")
}

func CreatePhoneNumber352(numbers [10]uint) string {
	localCode := GetNumbersString(numbers[:3])
	firstDigits := GetNumbersString(numbers[3:6])
	lastDigits := GetNumbersString(numbers[6:])
	return fmt.Sprintf("(%s) %s-%s", localCode, firstDigits, lastDigits)
}

//

func CreatePhoneNumber353(numbers [10]uint) string {
	var phone string
	for _, v := range numbers {
		phone += strconv.Itoa(int(v))
	}
	regex := regexp.MustCompile(`^(\d{3})(\d{3})(\d{4})$`)
	matches := regex.FindAllStringSubmatch(phone, -1)
	return `(` + matches[0][1] + `) ` + matches[0][2] + `-` + matches[0][3] + ``
}

//

func CreatePhoneNumber354(numbers [10]uint) string {

	//arrLength := len(numbers)
	phoneNumber := "("
	counter := 0

	for _, number := range numbers {
		if counter >= 0 && counter < 3 {
			phoneNumber += strconv.FormatUint(uint64(number), 10)

		} else if counter == 3 {
			phoneNumber += ") "
			phoneNumber += strconv.FormatUint(uint64(number), 10)

		} else if counter >= 4 && counter < 6 {
			//fmt.Println(number)
			phoneNumber += strconv.FormatUint(uint64(number), 10)

		} else if counter == 6 {
			phoneNumber += "-"
			phoneNumber += strconv.FormatUint(uint64(number), 10)

		} else {
			phoneNumber += strconv.FormatUint(uint64(number), 10)
		}

		//fmt.Println(phoneNumber)
		counter++
	}

	return phoneNumber
}

//

func CreatePhoneNumber355(num1 [10]uint) string {
	num := make([]int, len(num1))
	for i, v := range num1 {
		num[i] = int(v)
	}
	return "(" + strconv.Itoa(num[0]) + strconv.Itoa(num[1]) + strconv.Itoa(num[2]) + ") " + strconv.Itoa(num[3]) + strconv.Itoa(num[4]) + strconv.Itoa(num[5]) + "-" + strconv.Itoa(num[6]) + strconv.Itoa(num[7]) + strconv.Itoa(num[8]) + strconv.Itoa(num[9])
}

//

func CreatePhoneNumber356(numbers [10]uint) string {
	angkaAwal := convertSliceToString(numbers[:3])
	angkaTengah := convertSliceToString(numbers[3:6])
	angkaTerakhir := convertSliceToString(numbers[6:])

	phoneNumber := fmt.Sprintf("(%s) %s-%s", angkaAwal, angkaTengah, angkaTerakhir)

	return phoneNumber
}

func convertSliceToString(slice []uint) string {
	result := ""
	for _, num := range slice {
		result += strconv.Itoa(int(num))
	}
	return result
}

//

func CreatePhoneNumber357(numbers [10]uint) (r string) {
	for _, n := range numbers {
		r += strconv.Itoa(int(n))
	}
	return fmt.Sprintf("(%s) %s-%s", r[0:3], r[3:6], r[6:10])
}

//

func CreatePhoneNumber358(numbers [10]uint) (r string) {
	for _, n := range numbers {
		r += strconv.Itoa(int(n))
	}
	return "(" + r[:3] + ") " + r[3:6] + "-" + r[6:10]
}

//

func CreatePhoneNumber359(numbers [10]uint) string {
	strRet := "("
	str123 := "0123456789"

	for i, v := range numbers {
		if i == 3 {
			strRet += ") "
		}
		if i == 6 {
			strRet += "-"
		}
		strRet += string(str123[v]) //using the value as the index
	}
	return strRet
}

//

func CreatePhoneNumber360(numbers [10]uint) string {

	value := fmt.Sprintf("(%d) %d-%d", numbers[0:3], numbers[3:6], numbers[6:10])
	value = strings.ReplaceAll(value, "[", "")
	value = strings.ReplaceAll(value, "]", "")
	value = strings.ReplaceAll(value, " ", "")
	newstr := value[0:5] + " " + value[5:]

	return newstr
}

//

func intSliceToString(numbers []uint) string {
	stringText := make([]string, len(numbers))
	for i, number := range numbers {
		stringText[i] = fmt.Sprint(number)
	}

	return strings.Join(stringText, "")
}

func CreatePhoneNumber361(numbers [10]uint) string {
	phoneNumberParts := make([]string, 3)
	phoneNumberParts[0] = intSliceToString(numbers[0:3])
	phoneNumberParts[1] = intSliceToString(numbers[3:6])
	phoneNumberParts[2] = intSliceToString(numbers[6:])
	fmt.Println(phoneNumberParts)

	return fmt.Sprintf("(%s) %s-%s", phoneNumberParts[0], phoneNumberParts[1], phoneNumberParts[2])
}

//

func CreatePhoneNumber362(numbers [10]uint) string {
	strs := []rune{'('}

	for _, n := range numbers {
		strs = append(strs, rune(n+'0'))
	}

	newStrs := make([]rune, len(strs)+2)
	copy(newStrs, strs)
	newStrs = append(newStrs[:4], ')', ' ')
	newStrs = append(newStrs, strs[4:]...)

	newStrs2 := make([]rune, len(strs)+3)
	copy(newStrs2, newStrs)
	newStrs2 = append(newStrs2[:9], '-')
	newStrs2 = append(newStrs2, strs[7:]...)

	return string(newStrs2)
}

//

func CreatePhoneNumber363(numbers [10]uint) string {
	stringNumbers := make([]string, 0)

	for _, value := range numbers {
		stringNumbers = append(stringNumbers, fmt.Sprint(value)) // Append each value to the slice
	}
	format := "(%s) %s-%s"
	return fmt.Sprintf(format, strings.Join(stringNumbers[0:3], ""),
		strings.Join(stringNumbers[3:6], ""),
		strings.Join(stringNumbers[6:], ""))
}

//

func CreatePhoneNumber364(numbers [10]uint) string {
	s := make([]rune, 10)
	for i := range s {
		s[i] = rune(numbers[i]) + '0'
	}
	z := "(" + string(s[0]) + string(s[1]) + string(s[2]) + ") " + string(s[3]) + string(s[4]) + string(s[5]) + "-" + string(s[6]) + string(s[7]) + string(s[8]) + string(s[9])
	return z
}

//

func CreatePhoneNumber365(numbers [10]uint) string {
	str := "("
	for _, v := range numbers[:3] {
		char := rune(v + '0')
		str += string(char)
	}
	str += ") "
	for _, v := range numbers[3:6] {
		char := rune(v + '0')
		str += string(char)
	}
	str += "-"
	for _, v := range numbers[6:] {
		char := rune(v + '0')
		str += string(char)
	}
	return str
}

//

func CreatePhoneNumber366(numbers [10]uint) string {
	format := "(xxx) xxx-xxxx"
	for _, number := range numbers {
		num := int(number)
		format = strings.Replace(format, "x", strconv.Itoa(num), 1)
	}
	return format
}

//

func CreatePhoneNumber367(numbers [10]uint) string {
	var phone string

	var v0 string = strconv.FormatUint(uint64(numbers[0]), 10)
	var v1 string = strconv.FormatUint(uint64(numbers[1]), 10)
	var v2 string = strconv.FormatUint(uint64(numbers[2]), 10)
	var v3 string = strconv.FormatUint(uint64(numbers[3]), 10)
	var v4 string = strconv.FormatUint(uint64(numbers[4]), 10)
	var v5 string = strconv.FormatUint(uint64(numbers[5]), 10)
	var v6 string = strconv.FormatUint(uint64(numbers[6]), 10)
	var v7 string = strconv.FormatUint(uint64(numbers[7]), 10)
	var v8 string = strconv.FormatUint(uint64(numbers[8]), 10)
	var v9 string = strconv.FormatUint(uint64(numbers[9]), 10)

	phone = "(" + v0 + v1 + v2 + ") " + v3 + v4 + v5 + "-" + v6 + v7 + v8 + v9

	return phone
}

//

func CreatePhoneNumber368(numbers [10]uint) string {
	re := regexp.MustCompile("[^0-9]")
	s := re.ReplaceAllString(fmt.Sprintf("%v", numbers), "")
	return fmt.Sprintf("(%v) %v-%v", s[:3], s[3:6], s[6:])
}

//

func CreatePhoneNumber369(numbers [10]uint) string {
	var result string
	converter := func(number uint) string {
		return strconv.Itoa(int(number))
	}
	result += "("
	for n, v := range numbers {
		if n < 3 {
			result += converter(v)
		} else if n == 3 {
			result += ") "
			result += converter(v)
		} else if n < 6 {
			result += converter(v)
		} else if n == 6 {
			result += "-"
			result += converter(v)
		} else {
			result += converter(v)
		}

	}
	return result
}

//

func CreatePhoneNumber370(numbers [10]uint) string {
	s1, s2, s3 := "", "", ""
	result := ""

	for i := range numbers {
		if i <= 2 {
			s1 += strconv.Itoa(int(numbers[i]))
		} else if i > 2 && i <= 5 {
			s2 += strconv.Itoa(int(numbers[i]))
		} else if i > 5 {
			s3 += strconv.Itoa(int(numbers[i]))
		}

	}

	result = "(" + s1 + ")" + " " + s2 + "-" + s3

	return result
}

func UIntSliceToString(numbers []uint) string {
	var strNumbers string
	for _, num := range numbers {
		strNumbers += strconv.Itoa(int(num))
	}
	return strNumbers
}

func CreatePhoneNumber372(numbers [10]uint) string {
	firstPart := UIntSliceToString(numbers[0:3])
	secondPart := UIntSliceToString(numbers[3:6])
	thirdPart := UIntSliceToString(numbers[6:])
	phoneNumber := "(" + firstPart + ") " + secondPart + "-" + thirdPart
	return phoneNumber
}

//

func CreatePhoneNumber373(numbers [10]uint) string {
	var sonuc strings.Builder

	// Sayı dizisinin her bir elemanı üzerinde döngü.
	for j, i := range numbers {
		// İlk üç rakamı ekler.
		if j == 0 {
			sonuc.WriteString("(" + strconv.Itoa(int(i)))
		} else if j == 2 {
			// İkinci rakamın sonrasına parantez kapatma ve boşluk ekler.
			sonuc.WriteString(strconv.Itoa(int(i)) + ") ")
		} else if j == 5 {
			// Altıncı rakamın sonrasına tire ekler.
			sonuc.WriteString(strconv.Itoa(int(i)) + "-")
		} else if j == 9 {
			// Dokuzuncu rakamın sonrasına ekler ve döngüyü sonlandırır.
			sonuc.WriteString(strconv.Itoa(int(i)))
		} else {
			// Diğer durumlarda sadece rakam ekler.
			sonuc.WriteString(strconv.Itoa(int(i)))
		}
	}

	// Oluşturulan telefon numarasını string olarak döndürür.
	return sonuc.String()
}

//

func CreatePhoneNumber374(numbers [10]uint) string {
	initial, middle, end := numbers[:3], numbers[3:6], numbers[6:]
	return fmt.Sprintf("(%s) %s-%s", ToString(initial), ToString(middle), ToString(end))
}

func ToString(arr []uint) string {
	var res string
	for _, ar := range arr {
		res += strconv.Itoa(int(ar))
	}
	return res
}

//

func CreatePhoneNumber375(numbers [10]uint) string {
	return fmt.Sprintf(
		"(%v) %v-%v",
		numbersToString(numbers[0:3]),
		numbersToString(numbers[3:6]),
		numbersToString(numbers[6:]),
	)
}

func numbersToString(numbers []uint) string {
	return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

func CreatePhoneNumber376(numbers [10]uint) string {
	return "(" + convertToString1(numbers[:3]) + ") " + convertToString1(numbers[3:6]) + "-" + convertToString1(numbers[6:10])
}

func convertToString1(arr []uint) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ""), "[]")
}
