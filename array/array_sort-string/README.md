## Sorting Strings in Go

This Go program sorts a slice of strings in both **ascending** and **descending** order using the `sort` package.

### **Algorithm Explanation**

#### **Ascending Sort (`sort.Strings(xs)`)**
- Uses **Timsort**, a combination of **merge sort** and **insertion sort**.
- Modifies `xs` **in-place**, arranging elements in **lexicographical order**.

#### **Descending Sort (`sort.Sort(sort.Reverse(sort.StringSlice(xs)))`)**
- Converts `xs` into `sort.StringSlice`, which implements `sort.Interface`.
- Wraps it with `sort.Reverse`, which inverts comparisons for **reverse sorting**.
- Uses **Timsort** again to sort in descending order.

### **Final Output**
```
["Apple", "Banana", "Ben", "Kaaa", "No"]
["No", "Kaaa", "Ben", "Banana", "Apple"]
```