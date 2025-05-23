### Explanation of the Code:

This program counts the **occurrences of each unique element** in an array and prints the results. The code is written in Go and performs the following steps:

#### **Steps in the Program:**

1. **Initialization:**
   - `arr := []int{1, 11, 2, 1, 3, 4, 2, 2}`: This is the input array that contains numbers.
   - `numbers := []int{}`: This slice will store the unique elements from `arr` (elements that have been counted).
   - `repeats := []int{}`: This slice will store the count of each element's occurrence.

2. **Outer Loop (Iterate Through `arr`):**
   - The outer `for` loop goes through each element in `arr`. It will use `temp` to store the current element to be checked.
   - `temp := arr[i]`: The variable `temp` holds the current number from `arr` that needs to be checked for occurrences.
   
3. **Inner Loop (Count Occurrences of Each Element):**
   - The inner `for` loop compares the current element (`temp`) with all other elements in the array `arr` to count how many times `temp` occurs.
   - `if temp == arr[j]`: If `temp` matches `arr[j]`, `counter` is incremented to keep track of the number of occurrences of `temp`.

4. **Checking if the Element Has Already Been Counted:**
   - The third `for` loop checks if `temp` has already been added to the `numbers` slice (which stores already counted numbers). If `temp` is already in `numbers`, we set `isNotExist = false`, indicating that we don't need to count it again.
   - `for k := 0; k < len(numbers); k++`: This loop iterates through the `numbers` slice to check if `temp` already exists.

5. **Add New Element and Its Count:**
   - If `isNotExist` is `true`, this means that `temp` hasn't been counted yet. Therefore, it is added to the `numbers` slice, and the count (`counter`) is added to the `repeats` slice.
   - `numbers = append(numbers, temp)`: Adds the unique number to the `numbers` slice.
   - `repeats = append(repeats, counter)`: Adds the count of `temp` to the `repeats` slice.

6. **Output the Results:**
   - Finally, a loop is used to print each element in `numbers` and its corresponding count from `repeats`.
   - `fmt.Println(numbers[i], repeats[i])`: This prints each unique number and how many times it occurred.

#### **What Happens During Execution:**

Given the array `arr := []int{1, 11, 2, 1, 3, 4, 2, 2}`, the program will count the occurrences of each element.

- The element `1` occurs **2** times.
- The element `11` occurs **1** time.
- The element `2` occurs **3** times.
- The element `3` occurs **1** time.
- The element `4` occurs **1** time.

So, the output will be:
```
1 2
11 1
2 3
3 1
4 1
```

#### **Key Points to Understand:**

1. **Outer Loop:** 
   - Iterates over each element of the `arr`.
   
2. **Inner Loop:**
   - Compares the current element (`temp`) with every element in `arr` to count how many times it occurs.

3. **Avoid Double Counting:**
   - The third `for` loop ensures that each unique element is counted only once by checking if the element is already in the `numbers` slice.

4. **Final Output:**
   - For each unique element in `arr`, its value and its count are printed in the format `element count`.

#### **Complexity:**
- The program has a time complexity of O(n^2) because it contains two nested loops. The outer loop iterates over each element in the array, and for each element, it compares it with every other element using the inner loop.

#### **Output Example:**
```
1 2
11 1
2 3
3 1
4 1
```

In this case, `1` appears twice, `11` appears once, `2` appears three times, `3` appears once, and `4` appears once. The program correctly prints these counts in the specified format.