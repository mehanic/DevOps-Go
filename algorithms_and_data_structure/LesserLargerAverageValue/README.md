Let's break down the code and explain what it does step by step:

### **Code Explanation:**

```go
package main

import "fmt"

// Найти в массиве четные и нечетные элементы и
// поместить их в разные массивы и вывести эти массивы
func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}  // 1. Initialize an array arr1 with values {1, 2, 3, 4, 5}
	n := len(arr1)                   // 2. Get the length of the array arr1, which is 5
	sumi := 0                         // 3. Initialize sumi to store the sum of all elements in arr1
	var arr2 [100]int                 // 4. Declare arr2 to store elements less than the average
	var arr3 [100]int                 // 5. Declare arr3 to store elements greater than or equal to the average

	// 6. Calculate the sum of all elements in arr1
	for i := 0; i < n; i++ {
		sumi = sumi + arr1[i]
	}
	// 7. Calculate the average of the array arr1
	avg := sumi / n

	// 8. Initialize two indices for arr2 and arr3
	i2 := 0
	i3 := 0

	// 9. Iterate through arr1 to categorize elements into arr2 and arr3 based on their relation to avg
	for i := 0; i < n; i++ {
		if arr1[i] < avg {
			arr2[i2] = arr1[i]  // If element is less than avg, add it to arr2
			i2++                // Increment index for arr2
		}
		if arr1[i] >= avg {
			arr3[i3] = arr1[i]  // If element is greater than or equal to avg, add it to arr3
			i3++                // Increment index for arr3
		}
	}

	// 10. Print the original array, and the two categorized arrays
	fmt.Println(arr1)  // Print the original array arr1
	fmt.Println(arr2)  // Print the array arr2 (elements less than avg)
	fmt.Println(arr3)  // Print the array arr3 (elements greater than or equal to avg)
}
```

### **Step-by-step Explanation:**

1. **Initialize the original array `arr1`:**
   ```go
   arr1 := [5]int{1, 2, 3, 4, 5}
   ```
   - The array `arr1` is initialized with 5 integers: `[1, 2, 3, 4, 5]`.
   - `n := len(arr1)` gets the length of the array `arr1`, which is `5`.

2. **Calculate the sum of all elements in `arr1`:**
   ```go
   sumi := 0
   for i := 0; i < n; i++ {
       sumi = sumi + arr1[i]
   }
   ```
   - We iterate through each element of the array `arr1` and sum them up.
   - After the loop, `sumi` holds the sum of the elements: `1 + 2 + 3 + 4 + 5 = 15`.

3. **Calculate the average of the array `arr1`:**
   ```go
   avg := sumi / n
   ```
   - The average of the array is calculated by dividing the sum (`sumi = 15`) by the number of elements (`n = 5`).
   - So, `avg = 15 / 5 = 3`.

4. **Initialize two new arrays (`arr2` and `arr3`):**
   ```go
   var arr2 [100]int
   var arr3 [100]int
   ```
   - `arr2` is used to store values less than the average.
   - `arr3` is used to store values greater than or equal to the average.
   - Both arrays are initialized with 100 elements (defaulted to 0), but they will only be used up to the indices where values are inserted.

5. **Categorize elements into `arr2` and `arr3`:**
   ```go
   i2 := 0
   i3 := 0
   for i := 0; i < n; i++ {
       if arr1[i] < avg {
           arr2[i2] = arr1[i]
           i2++
       }
       if arr1[i] >= avg {
           arr3[i3] = arr1[i]
           i3++
       }
   }
   ```
   - We loop through each element in `arr1`.
   - If an element is less than the average (`avg = 3`), it is placed in `arr2`, and the index `i2` for `arr2` is incremented.
   - If an element is greater than or equal to the average, it is placed in `arr3`, and the index `i3` for `arr3` is incremented.

   Here's how the categorization works for the elements in `arr1`:
   - `arr1[0] = 1`: `1 < 3`, so it goes into `arr2`.
   - `arr1[1] = 2`: `2 < 3`, so it goes into `arr2`.
   - `arr1[2] = 3`: `3 >= 3`, so it goes into `arr3`.
   - `arr1[3] = 4`: `4 >= 3`, so it goes into `arr3`.
   - `arr1[4] = 5`: `5 >= 3`, so it goes into `arr3`.

6. **Print the arrays:**
   ```go
   fmt.Println(arr1)
   fmt.Println(arr2)
   fmt.Println(arr3)
   ```
   - Finally, the original array `arr1`, the array `arr2` (elements less than the average), and the array `arr3` (elements greater than or equal to the average) are printed.

### **Output:**
For the input array `arr1 = [1, 2, 3, 4, 5]`, the program will print:
```
[1 2 3 4 5]
[1 2 0 0 0]   // arr2: elements less than avg (3)
[3 4 5 0 0]   // arr3: elements greater than or equal to avg (3)
```

### **Summary:**
This program:
- Initializes an array `arr1` with some values.
- Computes the sum and average of the elements in `arr1`.
- Categorizes the elements into two arrays: one containing elements less than the average, and the other containing elements greater than or equal to the average.
- Finally, it prints all three arrays: the original array, the array with values less than the average, and the array with values greater than or equal to the average.