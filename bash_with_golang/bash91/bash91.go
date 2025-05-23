package main

import (
	"bufio"
	"fmt"
//	"os"
	"strings"
)

func main() {
	// First part: simulate `echo -e "line1\nline2" | awk 'BEGIN{print "START"} {print} END{print "END"}'`
	fmt.Println("START")
	fmt.Println("line1")
	fmt.Println("line2")
	fmt.Println("END")

	// Second part: simulate `echo | awk '{var1="v1"; var2="v2"; var3="v3"; print var1, var2, var3;}'`
	var1 := "v1"
	var2 := "v2"
	var3 := "v3"
	fmt.Println(var1, var2, var3)

	// Third part: simulate `echo -e "line1 f2 f3\nline2 f4 f5\nline3 f6 f7" | awk '{print "Line no: "NR", No of fields: "NF, "$0="$0,"$1="$1, "$2="$2, "$3="$3}'`
	input := "line1 f2 f3\nline2 f4 f5\nline3 f6 f7"
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		fields := strings.Fields(line)
		nFields := len(fields)
		
		// Prepare output according to the awk command
		fmt.Printf("Line no: %d, No of fields: %d, $0=%s, $1=%s, $2=%s, $3=%s\n",
			lineNumber, nFields,
			line,
			getField(fields, 0),
			getField(fields, 1),
			getField(fields, 2),
		)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

// Helper function to safely get a field by index
func getField(fields []string, index int) string {
	if index < len(fields) {
		return fields[index]
	}
	return ""
}

