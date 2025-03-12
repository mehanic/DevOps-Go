## Sorting Algorithm in Go

This Go program sorts an integer slice in both ascending and descending order using the `sort` package.

### Algorithm Explanation

#### **Ascending Sort (`sort.Ints(xi)`)**
- Uses **Timsort**, a hybrid sorting algorithm combining **merge sort** and **insertion sort**.
- Modifies `xi` **in-place**, arranging elements in increasing order.

#### **Descending Sort (`sort.Sort(sort.Reverse(sort.IntSlice(xi)))`)**
- Converts `xi` to `sort.IntSlice`, which implements `sort.Interface`.
- Wraps it with `sort.Reverse`, which inverts the sorting order.
- Uses **Timsort** again to sort in descending order.

### **Final Output**
```
[0 1 3 3 5 7 8 9]
[9 8 7 5 3 3 1 0]
```