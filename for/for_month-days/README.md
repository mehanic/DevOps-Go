Let's break down the Go program and explain the logic behind each section:

### Code Breakdown

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		year := rand.Intn(2100) + 1
		mounth := rand.Intn(12) + 1
		leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
		daysInMonth := 31

		if leap && mounth == 2 {
			daysInMonth = 29
		} else if !leap && mounth == 2 {
			daysInMonth = 28
		}

		if !leap {
			switch mounth {
			case 2:
				daysInMonth = 28
			case 4, 6, 9, 11:
				daysInMonth = 30
			}
		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, mounth, day)
	}
}
```

### Step-by-Step Explanation:

1. **Importing Packages**:
   - `import "fmt"`: This is used to print formatted output to the console.
   - `import "math/rand"`: This is used for generating random numbers.
   - `import "time"`: This is used for getting the current time, which will seed the random number generator.

2. **Global Variable `era`**:
   - `var era = "AD"`: This initializes a global variable `era` with the string `"AD"`. This will be printed in front of each randomly generated year.

3. **Seeding the Random Number Generator**:
   - `rand.Seed(time.Now().UnixNano())`: This sets the seed for the random number generator to the current Unix timestamp in nanoseconds. This ensures that the random numbers generated are different each time the program is run.

4. **The Loop**:
   - `for i := 0; i < 10; i++ { ... }`: This loop will run 10 times. For each iteration, the program will randomly generate a year, month, and day and print the result.

5. **Random Year**:
   - `year := rand.Intn(2100) + 1`: This generates a random integer between `1` and `2100` (inclusive). The `rand.Intn(2100)` generates a number between `0` and `2099`, and adding `1` ensures the year is between `1` and `2100`.

6. **Random Month**:
   - `mounth := rand.Intn(12) + 1`: This generates a random month between `1` and `12` (inclusive). `rand.Intn(12)` generates a number between `0` and `11`, and adding `1` ensures the month is between `1` and `12`.

7. **Leap Year Check**:
   - `leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)`: This checks whether the given `year` is a leap year. The conditions for a leap year are:
     - If the year is divisible by `400`, it is a leap year.
     - If the year is divisible by `4` but not divisible by `100`, it is a leap year.
     - Otherwise, it is not a leap year.

8. **Determining Days in Month**:
   - `daysInMonth := 31`: The program starts by assuming that the month has 31 days (as most months do).
   
   - **Leap Year February**:
     - `if leap && mounth == 2 { ... }`: If the month is February (`mounth == 2`) and the year is a leap year, it sets `daysInMonth` to 29 (February has 29 days in a leap year).
     - `else if !leap && mounth == 2 { ... }`: If the year is not a leap year and the month is February, it sets `daysInMonth` to 28 (February has 28 days in a common year).
   
   - **Non-Leap Year Adjustments**:
     - `if !leap { ... }`: If the year is not a leap year, the program adjusts the days in months based on the month:
       - For February (`mounth == 2`), the days are set to 28.
       - For April, June, September, and November (which have 30 days), the program sets `daysInMonth` to 30.

9. **Random Day Generation**:
   - `day := rand.Intn(daysInMonth) + 1`: This generates a random day for the chosen month. It uses `rand.Intn(daysInMonth)` to get a random number between `0` and `daysInMonth - 1`, and then adds `1` to ensure the day is between `1` and `daysInMonth` (inclusive).

10. **Printing the Result**:
    - `fmt.Println(era, year, mounth, day)`: The program prints the era (`"AD"`), the randomly generated year, month, and day.

### Example Output:

This program will generate 10 random dates in the format:
```
AD 2022 12 15
AD 1989 3 9
AD 1201 7 22
AD 1903 9 18
AD 2050 2 10
AD 1700 4 17
AD 2024 5 3
AD 1600 6 25
AD 2015 8 11
AD 2100 1 1
```

### Summary:
- The program generates a random year, month, and day.
- It checks if the year is a leap year and adjusts the number of days in February accordingly.
- It ensures that the number of days for each month is correct, considering leap years and months with fewer than 31 days.
- It then prints the results in the format `"AD year month day"`.
- The program repeats this process 10 times, generating and printing 10 random dates.