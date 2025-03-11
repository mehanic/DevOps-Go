### **Algorithm Steps Explanation**

1. **Function: `CreateTwoDimensionalArray(rows, columns, randomRange)`**:
   - This function creates a two-dimensional array (`arr`) of `rows` rows and `columns` columns, where each element is a random number between 0 and `randomRange - 1`.
   - It loops over each row (`i`) and within each row, generates random numbers (`k`) and appends them to the row (`a`).
   - After filling a row, it appends it to the 2D array (`arr`).

2. **Main Function**:
   - The `main` function calls `CreateTwoDimensionalArray` to generate a 2D array with 10 rows and 20 columns, and the random numbers are between 0 and 14.
   - The program then prints each row of the 2D array.

3. **Count Occurrences of `5`**:
   - The algorithm then checks each row of the 2D array. It iterates over each element of the row and counts the occurrences of the number `5`.
   - If the count of `5` in a row is 3 or more, it prints the row index (`i`) and the number of occurrences (`shetchik`).

### **Key Concepts**:
- **Random Number Generation**: The function uses `rand.Intn(randomRange)` to generate random integers within the specified range.
- **Nested Loops**: One loop handles the rows and another loop handles the columns within each row.
- **Condition Checking**: After counting the occurrences of `5` in each row, a condition checks if the count is greater than or equal to 3, triggering the output.

### **Example Output**:
The printed output shows all rows of the 2D array followed by any rows where the number `5` appears at least 3 times.

```
[3 5 1 2 5 ...]
[5 5 5 1 2 ...]
Repeat in row 1 with count 3
```