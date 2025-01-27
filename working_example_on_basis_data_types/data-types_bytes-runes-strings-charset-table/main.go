package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n",
		"literal", "dec", "hex", "encoded",
		strings.Repeat("-", 45))

	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
}

// literal    dec        hex        encoded
// ---------------------------------------------
// A          65         41         41
// B          66         42         42
// C          67         43         43
// D          68         44         44
// E          69         45         45
// F          70         46         46
// G          71         47         47
// H          72         48         48
// I          73         49         49
// J          74         4a         4a
// K          75         4b         4b
// L          76         4c         4c
// M          77         4d         4d
// N          78         4e         4e
// O          79         4f         4f
// P          80         50         50
// Q          81         51         51
// R          82         52         52
// S          83         53         53
// T          84         54         54
// U          85         55         55
// V          86         56         56
// W          87         57         57
// X          88         58         58
// Y          89         59         59
// Z          90         5a         5a

/*
EXAMPLE UNICODE BLOCKS

1 byte
------------------------------------------------------------
asciiStart     = '\u0001'      ->  32
asciiStop      = '\u007f'      ->  127

upperCaseStart = '\u0041'      ->  65
upperCaseStop  = '\u005a'      ->  90

lowerCaseStart = '\u0061'      ->  97
lowerCaseStop  = '\u007a'      ->  122


2 bytes
------------------------------------------------------------
latin1Start    = '\u0080'      ->  161
latin1Stop     = '\u00ff'      ->  255


3 bytes
------------------------------------------------------------
dingbatStart   = '\u2700'      ->  9984
dingbatStop    = '\u27bf'      ->  10175


4 bytes
------------------------------------------------------------
emojiStart     = '\U0001f600'  ->  128512
emojiStop      = '\U0001f64f'  ->  128591
*/
