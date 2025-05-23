package main

import "fmt"

type celsius float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

// по сути это такая виртуальнгая функция как в си, которая потом переопределяется
type getRowFn func(row int) (string, string)

func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)

	for i := 0; i < rows; i++ {
		cell1, cell2 := getRow(i)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func ctof(row int) (string, string) {
	c := celsius(row)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func ftoc(row int) (string, string) {
	f := fahrenheit(row)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("°C", "°F", 15, ctof)
	fmt.Println()
	drawTable("°F", "°C", 15, ftoc)
}
