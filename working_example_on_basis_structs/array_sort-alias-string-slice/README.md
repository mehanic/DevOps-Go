# Sorting a Custom String Slice in Go

This Go program defines a custom slice type `people` and sorts it in both **ascending** and **descending** order using the `sort` package.

## **Algorithm Explanation**

### **Ascending Sort (`sort.Strings(xs)`)**
- `people` is just an alias for `[]string`, so it can be sorted using `sort.Strings(xs)`.
- Uses **Timsort**, a hybrid of **merge sort** and **insertion sort**.
- Sorts `xs` **in-place** in **lexicographical order**.

### **Descending Sort (`sort.Sort(sort.Reverse(sort.StringSlice(xs)))`)**
- Converts `xs` to `sort.StringSlice`, which implements `sort.Interface`.
- Wraps it with `sort.Reverse` to swap comparisons.
- Uses **Timsort** again to sort in descending order.

## **Final Output**
```
["Apple", "Banana", "Ben", "Kaaa", "No"]
["No", "Kaaa", "Ben", "Banana", "Apple"]
```